package main

import (
	"context"
	"log"
	"os"
	"io"
	"time"
	"fmt"
	"strings"
)


func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags)
	ctx, cancel := context.WithTimeout(context.Background(), 5100*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beadInterval = 500 * time.Millisecond
	heartbeat, out := task(ctx, beadInterval)
loop:
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				break loop
			}
			fmt.Println("beat pulse 💡")
		case r, ok := <-out:
			if !ok {
				break loop
			}
			t := strings.Split(r.String(), "m=")
			fmt.Printf("value: %v [s]\n", t[1])
		case <-time.After(wdtTimeout):
			errorLogger.Println("watchdog timer expired")
			break loop
		}
	}

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "Hello"
	ch2 <- "World"
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
	// msgはランダムに出力される
}

func task(ctx context.Context, beatInterval time.Duration) (
	<-chan struct{}, <-chan time.Time,
) {
	heartbeat := make(chan struct{})
	out := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(2 * beatInterval)
		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
			// select文を利用する意図
			// 非ブロッキング送信: 受信側不在でも処理を止めずに進めるため
			// 遅延回避: バッファ満杯時に待たずスキップしてループ周期を維持
			// 最新性重視: 古いハートビートを溜めず、最新の信号だけを優先
			// デッドロック/スタック防止: 受信側停止や遅延時に送信側が詰まるのを避ける
			// キャンセル/タイムアウト連携: 他ケース（ctx.Doneなど）と同時に監視・分岐できるため
		}
		sendValue := func(t time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- t:
					return
				}
			}
			// for文を利用する意図
			// 継続送信の実現: 値を一定間隔・継続的に送るための基本構造
			// バックプレッシャー対応: 受信状況に応じて送信を自然にブロック/再開できる
			// 終了条件の集約: selectでctx.Done()やタイムアウトを随時監視し安全に抜けられる
			// 簡潔な状態管理: ループ内でリトライ・スキップ・遅延などの分岐を一か所で扱える
			// ハートビート連携: 送信処理と並行して心拍送信や監視を織り込める構成にしやすい
		}
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			case <-pulse.C:
				if i == 3 {
					time.Sleep(1 * time.Second)
				}
				sendPulse()
				i++
			case t := <-task.C:
				sendValue(t)
			}
		}
		// for文を利用する意図
		// 負荷分離と応答性確保: 心拍（pulse）と実処理（task）を同じfor-selectに集約しつつ、送信処理はsendPulse/sendValueに分離して、ループ自体を軽く保ち遅延を防ぐ。
		// 非ブロッキング化: ループはイベント多重化に専念し、各処理のブロック可能箇所は関数側で制御（例: sendPulseはdefault付きselectで即戻す）。
		// キャンセル一元化: ctx.Done()監視をループ一箇所に集約し、全処理の停止制御を明確化。
		// 可読性・保守性: イベント駆動の「待つ場所」と、実際の「送る/処理する場所」を分け、責務を明確化してテストもしやすくする。
		// 優先度/ポリシー実装の余地: 将来的に心拍の落とし許容、タスクの再試行/バックオフなど、関数側でポリシーを独立に調整可能。
	}()

	return heartbeat, out	
}
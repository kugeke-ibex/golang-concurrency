# 主な Context の用途

main goroutine から sub goroutine をキャンセルさせる

- main goroutine parent context
  - fun1 cancel
  - fun2 cancel + timeout

# func WithCancel(paranet Context)

親 Context からキャンセル関数を生成し、明示的にキャンセルを伝播させる。I/O や複数 goroutine の協調停止に使う。

- 戻り値: (Context, cancel func)
- キャンセル条件: `cancel()`呼び出し or 親のキャンセル
- 代表用途: 明示キャンセル、リソース解放の確実化

```go
ctx, cancel := context.WithCancel(parent)
defer cancel() // 早期return時も確実に通知
go func() {
    // 何かの処理
    select {
    case <-ctx.Done():
        // 後片付けして終了
    }
}()
// 必要になったら
cancel()
```

# func WithDeadline(parent Context, d time.Time)

指定した絶対時刻 d でキャンセルされる Context を返す。外部 API や DB アクセスで「この時刻までに完了」を保証したいときに使う。

- 戻り値: (Context, cancel func)
- キャンセル条件: d に到達、親のキャンセル、`cancel()`呼び出し
- 代表用途: 期限付きリクエストの全体制御

```go
deadline := time.Now().Add(1500 * time.Millisecond)
ctx, cancel := context.WithDeadline(parent, deadline)
defer cancel()
// ctxを渡してI/Oや処理を実行。期限超過で ctx.Err() == context.DeadlineExceeded
```

# func WithTimeout(parent Context, timeout time.Duration)

指定した相対時間 timeout 経過でキャンセルされる Context を返す。`WithDeadline`の相対時間版で、より一般的に使われる。

- 戻り値: (Context, cancel func)
- キャンセル条件: timeout 経過、親のキャンセル、`cancel()`呼び出し
- 代表用途: 各処理にタイムボックスを設ける、リーク防止のための確実なキャンセル

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
// 処理が2秒以内に終わらなければ context.DeadlineExceeded
```

# context.Background()

ルート（最上位）の空の Context を返す。キャンセルされず、値も期限も持たない不変の Context。

- 用途: サーバ起動時の親 Context、テストや例示のルート、main 関数の起点など
- 特性: 子 Context（WithCancel/WithTimeout/WithDeadline など）の親として使う
- 対比: 一時的で未決の場面は `context.TODO()` を用いる（意味上の違い）

```go
// 典型: サービスのルート Context
func main() {
    ctx := context.Background()
    // 必要に応じて子 Context を派生
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()
    // ctx を下位の処理へ渡す
}
```

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	flags := log.Lshortfile // 時刻情報をログに含める(ファイル名と行番号)
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags) // ログをファイルと標準エラーに出力
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	warnLogger.Println("warning A")

	errorLogger.Fatalln("critical error")
}
package main

import (
	"log"
	"os"
)

// > Go 言語の log パッケージを使ってみる - 倭マン's BLOG
// > https://waman.hatenablog.com/entry/2017/09/29/011614
//const (
//  Ldate         = 1 << iota  // 日付
//  Ltime                      // 時刻
//  Lmicroseconds              // 時刻のマイクロ秒
//  Llongfile                  // ソースファイル（ディレクトリパスを含む）
//  Lshortfile                 // ソースファイル（ファイル名のみ）
//  LUTC                       // タイムゾーンに依らない UTC 時刻
//  LstdFlags     = Ldate | Ltime  // 日付 (Ldata) と時刻 (Ltime) ：デフォルト
//)

func init() {
	// 時刻と時刻のマイクロ秒、ディレクトリパスを含めたファイル名を出力
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	log.Printf("test")

	// Create new logger ( Output to stdout )
	logger := log.New(os.Stdout, "[MyLogger] ", log.Llongfile|log.LstdFlags)
	logger.Println("test2")
}

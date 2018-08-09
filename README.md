## go-sample-box

### Development environment

 - [Downloads - The Go Programming Language](https://golang.org/dl/)
 - [Go with Visual Studio Code](https://code.visualstudio.com/docs/languages/go)
 - [Go と Travis CI を連携したり Golint を実行してみたり - kakakakakku blog](https://kakakakakku.hatenablog.com/entry/2015/12/25/233540)

### VSCodeでメソッド補完等が効かない

  - [Golangで自分自身で定義したパッケージをインポートする方法あれこれ - Qiita](https://qiita.com/shopetan/items/eddcacec21cc7ea274f9)

  この記事にならって、`~/go/src/[username]/[repository]`配下にcloneしてきて、
  パッケージの参照を`"github.com/xshoji/go-sample-box/go-first-sample/structs"`とかにしたらうまくいった

### クロスコンパイル

 - [Go のクロスコンパイル環境構築 - Qiita](https://qiita.com/Jxck_/items/02185f51162e92759ebe)

```
GOOS=linux GOARCH=amd64 go build -o /tmp/webapp main.go
```

## References

 - [gostor/awesome-go-storage： A curated list of awesome Go storage projects and libraries](https://github.com/gostor/awesome-go-storage)
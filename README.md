## go-sample-box

## 開発の準備

まず、gvmを導入してgoを複数バージョン管理できるようにします。

> .00_Golangメモ.md  
> https://gist.github.com/xshoji/ea284689bda36fbdaa67e83ba713517f#file-01_memo-md

次に、このリポジトリを自分の好きなディレクトリにcloneしてます。

```
cd Develop/go
git clone https://github.com/xshoji/go-sample-box
```

gvmのsrcディレクトリにシンボリックリンクを貼ります。

```
ln -s ~/Develop/go/go-sample-box ~/.gvm/pkgsets/go1.9.7/global/src/github.com/xshoji/go-sample-box
```

Intellijで開発します。

> .00_Intellij_IDEAメモ.md  
> https://gist.github.com/xshoji/7c16af201ab283d1c2dcfee65a5aea7d#file-02_golang-md

実行は、

```
cd ~/.gvm/pkgsets/go1.9.7/global/src/github.com/xshoji/go-sample-box/argumentGoFlags
```

等の各ディレクトリで`go run main.go`で実行します。

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
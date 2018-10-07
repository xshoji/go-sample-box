package main

import (
	"github.com/xshoji/go-sample-box/jsonDynamic/caseSameKeyDifferentStructure"
	"github.com/xshoji/go-sample-box/jsonDynamic/caseDifferentKeyDifferentStructure"
)

// > 動的な要素を持つJSONをいい感じにUnmarshalする - Qiita
// > https://qiita.com/kanga/items/6929c6900ccbfa9dc933
//
// > Structをomitemptyする方法 - Qiita
// > https://qiita.com/taroshin/items/59e62d8ec852c5f84bc1
//
// > Dynamic JSON umarshalling in Go – Nathan Smith – Medium
// > https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0
//
// > flexjson/model.go at master · neocortical/flexjson
// > https://github.com/neocortical/flexjson/blob/master/flexjson/model.go
//
// > go - Unmarshaling dynamic JSON - Code Review Stack Exchange
// > https://codereview.stackexchange.com/questions/68915/unmarshaling-dynamic-json
//
// > ［Go］ JSONの中身に応じて違うstructにデコードする - Qiita
// > https://qiita.com/hnakamur/items/f54e867345c52624d5d7
//
// > mattyw： Using go to unmarshal json lists with multiple types
// > http://mattyjwilliams.blogspot.com/2013/01/using-go-to-unmarshal-json-lists-with.html
//
// > interface要素を持つstructへのJSON Unmarshal - すぎゃーんメモ
// > https://memo.sugyan.com/entry/2018/06/23/232559
func main() {
	// ケース1: 同じキー名だけど中身の構造が違うjsonのunmarshalの例
	caseSameKeyDifferentStructure.Run()

	// ケース2: 違うキー名で中身の構造が違うjsonのunmarshalの例
	caseDifferentKeyDifferentStructure.Run()
}

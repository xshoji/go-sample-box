package utils

// JoinString 文字列を連結する
// 外部に公開する場合は先頭を大文字に
func JoinString(a, b string) (result string) {
	result = a + b
	return
}

// GetMultiReturns 複数の返り値を返す
func GetMultiReturns() (a, b string) {
	a = "Aiueo"
	b = "Kakikukeko"
	return
}

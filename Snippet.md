# Snippet

```go
// primitive types
var valueInt8 int8
valueInt8 =
valueInt := 100
valueInt64 := time.Now().UnixNano()
valueFloat64 := 100.12
valueRune := '\n'
valueString := "test"
valueBytes := []byte("test")
fmt.Println("<< primitive >>")
fmt.Printf("valueInt, valueInt64, valueFloat, valueRune, valueString\n%v, %v, %v, %v, %v\n", valueInt, valueInt64, valueFloat64, valueRune, valueString)

// print type
fmt.Printf("valueInt %v(%T)\n", valueInt, valueInt)

//
//
//
//
//
// cast
// > Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記
// > http://matope.hatenablog.com/entry/2014/04/22/101127
// FormatIntの第2引数は基数。2なら2進数、16なら16進数になる
// int -> string
strconv.Itoa(valueInt)
// int64 -> string
strconv.FormatInt(int64(valueInt8), 10)
// floag -> string
strconv.FormatFloat(valueFloat64, 'f', 4, 32)
// string -> []byte
[]byte(valueString)
// string -> int
valueInt, _ = strconv.Atoi(valueString)
// string -> int64
valueInt64, _ = strconv.ParseInt(valueString, 10, 64)
// []byte -> string
string(valueBytes)


//
//
//
//
//
//string cut
a := "aiueo"
fmt.Println(a[:3]) // aiu


//
//
//
//
//
//map, for
mapValues := map[string]string{
    "aaa": "aaa_value",
    "bbb": "bbb_value",
    "ccc": "ccc_value",
}
v, ok := mapValues["aaa"] // v -> "aaa_value", ok -> true
v, ok := mapValues["xxx"] // v -> "",          ok -> false

// for
fmt.Println("<< for >>")
fmt.Printf("valueStringSlice size:%d\n", len(valueStringSlice))
for _, v := range valueStringSlice {
    fmt.Println(v)
}
for k, v := range mapValues {
    fmt.Println(k + ":" + v)
}
fmt.Println()

//Nested maps
//> Goで多次元マップ（複数のキーからなるマップ）を実現したいときにはどうするか - Qiita
//> https://qiita.com/ruiu/items/476f65e7cec07fd3d4d7
nestedMap := make(map[string]map[string]string)
if _, ok := nestedMap["key1"]; !ok {
    nestedMap["key1"] = make(map[string]string)
}
nestedMap["key1"]["key2"] = "value"
fmt.Println(nestedMap)
fmt.Println()

//Nested maps (key object)
//> Go maps in action - The Go Blog
//> https://blog.golang.org/go-maps-in-action#TOC_5.
type Key struct {
    key1, key2 string
}
nestedMap2 := make(map[Key]string)
nestedMap2[Key{"key1", "key2"}] = "value"
fmt.Println(nestedMap2)
fmt.Println()


//
//
//
//
//
// anonymous struct
anonymousStruct := struct {
	Name     string
	Age      int
	Language string
}{
	"taro",
	100,
	"Japanese",
}

// anonymous struct array
anonymousStruct := []struct {
	Name     string
	Age      int
	Language string
}{
	{
		Name:     "taro",
		Age:      12,
		Language: "taro",
	},
}


//
//
//
//
//
// Create random integer
fmt.Println("<< random integer >>")
createRandomNumber := func() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000-1) + 1
}
fmt.Println(createRandomNumber())

// Create random string
fmt.Println("<< random string >>")
createRandomString := func() string {
	seed := strconv.FormatInt(time.Now().UnixNano(), 10)
	shaBytes := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(shaBytes[:])
}
fmt.Println(createRandomString())

//
//
//
//
//
// DateTime format
// > Golangでの日付のフォーマット指定の方法について - Qiita
// > https://qiita.com/unbabel/items/c8782420391c108e3cac
fmt.Println("<< now >>")
now := time.Now()
fmt.Println("plane:")
fmt.Println(now)
fmt.Println("formatted:")
fmt.Println(now.Format("2006 / 01 [January(Jan)] / 02 [Monday(Mon)] 15:04:05 [MST]"))

//
//
//
//
//
// regexp.MatchString
fmt.Println("<< regexp.MatchString >>")
fmt.Println(regexp.MatchString(`^aaa`, mozi))
fmt.Println(regexp.MatchString(`ccc$`, mozi))
fmt.Println(regexp.MatchString(`ddd`, mozi))

//
//
//
//
//
// ファイルを読み込む
// - [Go でファイルを1行ずつ読み込む（csv ファイルも） - Qiita](https://qiita.com/ikawaha/items/28186d965780fab5533d)
// １行ずつ読む
func readEachLines(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
// 一気に読む
func readAllLines(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		log.Panic(err)
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", string(contents))
}

//
//
//
//
//
// 多段パッケージ風staticクラスの実装例
type common struct{}
var Common = common{}
func (c *common) GetHttp(url string) string {}

//
//
//
//
//
// httpリクエスト

// GET
resp, err := http.Get(urlFull)

// POST
//form := url.Values(postData)
//requestBody := form.Ecncode()
//contentType := "text/plain"
requestBody := `{"aaa":"bbb"}`
contentType := "application/json"
resp, err := http.Post(urlFull, contentType, strings.NewReader(requestBody))

// Response handling
body, err := ioutil.ReadAll(resp.Body)
var result interface{}
json.Unmarshal(body, &result)
defer func() {
	err := resp.Body.Close()
	if err != nil {
		log.Panic("resp.Body.Close() failed.")
	}
}()


//
//
//
//
//
// リクエストBodyをinterface{}型のjsonオブジェクトに変換する
func ToJsonObject(body []byte) interface{} {
	var jsonObject interface{}
	json.Unmarshal(body, &jsonObject)
	return jsonObject
}

// interface{}型のjsonオブジェクトからキー指定で値を取り出す（object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb"）
func Get(object interface{}, keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	var result interface{}
	var exists bool
	for _, key := range keys {
		exists = false
		if _, ok := object.(map[string]interface{}); ok {
			exists = true
			object = object.(map[string]interface{})[key]
			result = object
			continue
		}
		if values, ok := object.([]interface{}); ok {
			for i, v := range values {
				if strconv.FormatInt(int64(i), 10) == key {
					exists = true
					object = v
					result = object
					continue
				}
			}
		}
	}
	if exists {
		return result
	}
	return nil
}

// 値をjson形式の文字列に変換する
func ToJsonString(v interface{}) string {
	result, _ := json.Marshal(v)
	return string(result)
}

//
//
//
// * generage code and run
// go generate ./... && go run main.go
//
// Enum
//
//go:generate stringer -type=DeviceType
type DeviceType int

const (
	DeviceTypeNull DeviceType = iota
	DeviceTypeIos
	DeviceTypeAndroid
	DeviceTypeWindows
	DeviceTypeLinux
)


//
//
//
//
// Test
func TestGetSuccess(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result1, result2 := ReplaceString("test[0...10]")
		if result1 != "test" {
			t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
		}
		if result2 != "[0...10]" {
			t.Fatalf("failed ReplaceString, %v, %v\n", result1, result2)
		}
	})
}



//
//
//
// Safety defer pattern
//-------------------------------
// Define defer function
deferFunc := func() {
    log.Println("Call defer function")
}
defer deferFunc() // set defer

// Make kill signal channel
signals := make(chan os.Signal, 1)
signal.Notify(signals, os.Kill, os.Interrupt)

go func() {
    <-signals
    fmt.Println("")
    log.Println("Catch signals")
    deferFunc()
    log.Println("Execute os.Exit()")
    os.Exit(0)
}()

//-------------------------------
// Define main logic
for i := 0; i < 5; i++ {
    log.Printf("Loop count is %v\n", i)
    duration := time.Duration(1000 * time.Millisecond)
    time.Sleep(duration)
}

log.Println("Finish loop")

```

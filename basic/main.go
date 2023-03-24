package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/xshoji/go-sample-box/basic/structs"
	"github.com/xshoji/go-sample-box/basic/utils"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//
	//
	// primitive types
	var valueInt8 int8
	valueInt8 = 1
	valueInt := 100
	valueInt64 := time.Now().UnixNano()
	valueFloat64 := 100.12
	valueRune := '\n'
	valueString := "test"
	valueBytes := []byte("test")
	fmt.Println("<< primitive >>")
	fmt.Printf("valueInt, valueInt64, valueFloat, valueRune, valueString\n%v, %v, %v, %v, %v\n", valueInt, valueInt64, valueFloat64, valueRune, valueString)
	fmt.Println()

	//
	//
	// 文字シーケンス
	valueString2 := `{
  "aaa", "bbb",
  "ccc": 111
}`
	fmt.Println("<< foundString sequence >>")
	fmt.Println(valueString2)
	fmt.Println()

	//
	//
	// cast
	// > Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記
	// > http://matope.hatenablog.com/entry/2014/04/22/101127
	fmt.Println("<< cast >>")
	// FormatIntの第2引数は基数。2なら2進数、16なら16進数になる
	fmt.Printf("valueInt: %d    -> valueString: %s\n", valueInt, strconv.Itoa(valueInt))
	fmt.Printf("valueInt8: %d   -> valueString: %s\n", valueInt, strconv.FormatInt(int64(valueInt8), 10))
	fmt.Printf("valueInt64: %d  -> valueString: %s\n", valueInt64, strconv.FormatInt(valueInt64, 10))
	fmt.Printf("valueFloat: %v  -> valueString: %s\n", valueFloat64, strconv.FormatFloat(valueFloat64, 'f', 4, 32))
	fmt.Printf("valueFloat: %v  -> valueInt: %v\n", valueFloat64, int(valueFloat64))
	fmt.Printf("valueInt: %v    -> valueFloat: %v\n", valueInt, float64(valueInt))
	fmt.Printf("valueString: %v -> value[]byte: %v\n", valueString, []byte(valueString))
	valueInt, _ = strconv.Atoi(valueString)
	fmt.Printf("valueString: %v -> valueInt: %v\n", valueString, valueInt)
	valueInt64, _ = strconv.ParseInt(valueString, 10, 64)
	fmt.Printf("valueString: %v -> valueInt64: %v\n", valueString, valueInt64)
	fmt.Printf("value[]byte: %v -> valueString: %s\n", valueBytes, string(valueBytes))
	fmt.Println()

	//
	//
	// pointer
	_valueInt := &valueInt
	fmt.Println("<< pointer >>")
	fmt.Printf("valueInt pointer: %v\n", _valueInt)
	fmt.Printf("valueInt pointers value: %v\n", *_valueInt)
	fmt.Println()

	//
	//
	// array
	valueIntArray := [...]int{1, 2, 3}
	fmt.Printf("valueIntArray\n%v\n", valueIntArray)
	fmt.Println()

	//
	//
	// slice
	var valueStringSlice []string
	valueStringSlice = []string{"aaa", "bbb", "ccc"}
	fmt.Printf("valueStringSlice\n%v\n", valueStringSlice)
	valueStringSlice = append(valueStringSlice, "ddd")
	fmt.Printf("%v\n", valueStringSlice)
	valueStringSlice = valueStringSlice[0:2]
	fmt.Printf("%v\n", valueStringSlice)
	valueStringSlice = make([]string, 5)
	fmt.Println(valueStringSlice)
	fmt.Println()

	//
	//
	//map
	mapValues := map[string]string{
		"aaa": "aaa",
		"bbb": "aaa",
		"ccc": "aaa",
	}

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

	//
	//
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

	//
	//
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
	// public method
	// Package prefix is required when calling method
	fmt.Println("<< function >>")
	fmt.Printf("%v\n", utils.JoinString("aaa", "bbb"))
	fmt.Printf("%v\n", utils.JoinString(utils.GetMultiReturns()))
	fmt.Println()

	//
	//
	// initialize struct
	fmt.Println("<< structs, method >>")
	taro := structs.NewUser("taro", 16)
	jiro := structs.NewUser("jiro", 99)
	// メソッドの呼び出し
	taro.Talk()
	jiro.Talk()
	//
	// 本体を引数にする引数名付きの生成パターン
	hanako := structs.NewUserDefault(structs.User{
		Name: "hanako",
	})
	yasuyo := structs.NewUserDefault(structs.User{
		Name: "yasuyo",
		Age:  100,
	})
	hanako.Talk()
	yasuyo.Talk()
	fmt.Println()

	//
	//
	// anonymous function
	anonymousFunction := func() {
		fmt.Println("This is anonymousFunction.")
	}
	fmt.Println("<< anonymous function >>")
	anonymousFunction()
	fmt.Println()

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
	fmt.Println("<< anonymous struct >>")
	fmt.Print(anonymousStruct)
	fmt.Println()

	toJson := func(v interface{}) string {
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	anonymousStruct2Name := "taro"
	anonymousStruct2Age := 100
	anonymousStruct2Lang := "Japanese"
	anonymousStruct2 := struct {
		Name     *string
		Age      *int
		Language *string
	}{}
	fmt.Println("<< anonymous struct with pointer>>")
	fmt.Println(anonymousStruct2)
	fmt.Println(toJson(anonymousStruct2))
	anonymousStruct2.Name = &anonymousStruct2Name
	fmt.Println(anonymousStruct2)
	fmt.Println(toJson(anonymousStruct2))
	anonymousStruct2.Age = &anonymousStruct2Age
	anonymousStruct2.Language = &anonymousStruct2Lang
	fmt.Println(anonymousStruct2)
	fmt.Println(toJson(anonymousStruct2))
	fmt.Println()

	//
	//
	// Create random integer
	fmt.Println("<< random integer >>")
	createRandomNumber := func() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(1000000-1) + 1
	}
	fmt.Println(createRandomNumber())
	fmt.Println()

	//
	//
	// Create random foundString
	fmt.Println("<< random foundString >>")
	createRandomString := func() string {
		seed := strconv.FormatInt(time.Now().UnixNano(), 10)
		shaBytes := sha256.Sum256([]byte(seed))
		return hex.EncodeToString(shaBytes[:])
	}
	fmt.Println(createRandomString())
	fmt.Println()

	fmt.Println("<< size >>")
	sizeString := "aaa"
	sizeArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sizeSlice := make([]string, 15)
	fmt.Printf("Size of foundString : %v\n", len(sizeString))
	fmt.Printf("Size of array : %v\n", len(sizeArray))
	fmt.Printf("Size of slice : %v\n", len(sizeSlice))
	fmt.Println()

	//
	//
	// Nullable variables
	// > go - How to make a nullable field in Golang struct? - Stack Overflow
	// > https://stackoverflow.com/questions/51998165/how-to-make-a-nullable-field-in-golang-struct
	fmt.Println("<< Nullable values handling >>")
	fmt.Println("[ pointer ]")
	user := struct {
		Name *string
		Age  *int
	}{}
	user.Name = &valueString
	user.Age = &valueInt
	fmt.Println("Refer values")
	fmt.Println(*user.Name)
	fmt.Println(*user.Age)
	fmt.Println(user.Name == nil)
	fmt.Println(user.Age == nil)

	fmt.Println()
	fmt.Println("Refer nil values")
	user.Name = nil
	user.Age = nil
	fmt.Println(user.Name == nil)
	fmt.Println(user.Age == nil)
	fmt.Println()

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
	fmt.Println()

	//
	//
	// Replace
	fmt.Println("<< strings.Replace >>")
	mozi := "aaabbbaaaccc"
	fmt.Println(strings.Replace(mozi, "aaa", "!!!", -1))
	fmt.Println()

	//
	//
	// regexp.MatchString
	fmt.Println("<< regexp.MatchString >>")
	fmt.Println(regexp.MatchString(`^aaa`, mozi))
	fmt.Println(regexp.MatchString(`ccc$`, mozi))
	fmt.Println(regexp.MatchString(`ddd`, mozi))
	fmt.Println()

	//
	//
	// ReplaceAllString
	fmt.Println("<< ReplaceAllString >>")
	mozi = "[111-222-333] [skdflskdjflsd] xxxx yyyy pppp 500 200 100"
	re := regexp.MustCompile(`(\[.*\])`)
	fmt.Println(re.ReplaceAllString(mozi, "test"))
	//> 正規表現：文字列を「含まない」否定の表現まとめ ｜ WWWクリエイターズ
	//> http://www-creators.com/archives/1827
	re = regexp.MustCompile(`(\[[^\]]*\])`)
	fmt.Println(re.ReplaceAllString(mozi, "test"))
	fmt.Println()

	//
	//
	//
	// FindAllString
	fmt.Println("<< FindAllString >>")
	mozi = "[111-222-333] [skdflskdjflsd] xxxx yyyy pppp 500 200 xxxx 100"
	findRegexp := regexp.MustCompile("xxxx")
	stringSlice := findRegexp.FindAllString(mozi, -1)
	for _, foundString := range stringSlice {
		fmt.Println(foundString)
	}
	fmt.Println()

	//
	//
	// > go - How to run a shell command in a specific folder - Stack Overflow
	// > https://stackoverflow.com/questions/43135919/how-to-run-a-shell-command-in-a-specific-folder
	//
	// Exec os command
	fmt.Println("<< Exec os command >>")
	cmd := exec.Command("ls", "-al")
	cmd.Dir = "/tmp" // Working directory
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	fmt.Println()

	//
	//
	//
	//
	//
	//
	// hash algorithm
	// - [Golang - How to hash a string using MD5.](https://gist.github.com/sergiotapia/8263278)
	fmt.Println("md5.Sum: ", fmt.Sprintf("%x", md5.Sum([]byte("test"))))
	fmt.Println("sha1.Sum: ", fmt.Sprintf("%x", sha1.Sum([]byte("test"))))
	fmt.Println("sha256.Sum256: ", fmt.Sprintf("%x", sha256.Sum256([]byte("test"))))
	fmt.Println("sha512.Sum512: ", fmt.Sprintf("%x", sha512.Sum512([]byte("test"))))

}

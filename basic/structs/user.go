package structs

import (
	"fmt"
)

// User user
type User struct {
	Name string
	Age  int
}

// NewUser NewUser
//  - [［Go］ 構造体の初期化方法まとめ - Qiita](https://qiita.com/cotrpepe/items/b8e7f70f27813a846431#newanimal-%E3%81%A8-animal-%E3%81%AE%E9%81%95%E3%81%84)
func NewUser(name string, age int) *User {
	user := new(User)
	user.Name = name
	user.Age = age
	return user
}

// NewUser NewUser
//  - [［Go］ 構造体の初期化方法まとめ - Qiita](https://qiita.com/cotrpepe/items/b8e7f70f27813a846431#newanimal-%E3%81%A8-animal-%E3%81%AE%E9%81%95%E3%81%84)
func NewUserDefault(arg User) *User {
	user := new(User)
	user.Name = arg.Name
	// default
	user.Age = 20
	if arg.Age != 0 {
		user.Age = arg.Age
	}
	return user
}

// Talk talk
func (u User) Talk() {
	fmt.Printf("My name is %s, %d years old\n", u.Name, u.Age)
}

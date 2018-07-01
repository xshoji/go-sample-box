package structs

import (
	"fmt"
)

// User user
type User struct {
	Name string
	Age  int
}

// Talk talk
func (u User) Talk() {
	fmt.Printf("My name is %s, %d years old\n", u.Name, u.Age)
}

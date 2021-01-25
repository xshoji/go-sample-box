package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/packagenamecollision/oauth/logic/authenticator"
	authenticator2 "github.com/xshoji/go-sample-box/packagenamecollision/user/authenticator"
)

func main() {
	fmt.Println(authenticator.IsAuthenticated())
	fmt.Println(authenticator2.IsAuthenticated())
}

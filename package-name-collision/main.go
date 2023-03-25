package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/package-name-collision/oauth/logic/authenticator"
	authenticator2 "github.com/xshoji/go-sample-box/package-name-collision/user/authenticator"
)

func main() {
	fmt.Println(authenticator.IsAuthenticated())
	fmt.Println(authenticator2.IsAuthenticated())
}

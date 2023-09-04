package main

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	message := "Hello, OTUS!"
	fmt.Print(stringUtils.Reverse(message))
}

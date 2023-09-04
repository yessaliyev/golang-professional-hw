package main

import (
	"fmt"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	text := "Hello, OTUS!"
	fmt.Print(stringUtils.Reverse(text))
}

package main

import (
	"fmt"

	"github.com/cnych/go-github-actions/hello"
)

func main() {
	fmt.Println(hello.Greet())
	fmt.Println(hello.Greet())
}

package main

import (
	"fmt"

	"github.com/sharmameghansh/go-demo-1.git/hello"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("%s\n", hello.HelloGolang())
	fmt.Println(quote.Hello())
}

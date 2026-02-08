package main

import (
	"fmt"
	"runtime"
)

func main() {

	_, line, _, _ := runtime.Caller(0)
	fmt.Println("\t" + line)

}
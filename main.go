package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {

	_, line, _, _ := runtime.Caller(0)
	fmt.Println("R Caller:", line)

	var modulePath string = func() string {
		info, _ := debug.ReadBuildInfo()
		return info.Path
	}()

	fmt.Println("Mod Path:", modulePath)

}
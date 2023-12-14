package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("no build info")
	}
	fmt.Printf("Build Info %s", info)
}

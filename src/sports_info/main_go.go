package main

import (
	"fmt"
	"runtime"
	"sports_info/game"
)

func main() {
	fmt.Println("main act, golang version:", runtime.Version())
	game.Start()
}

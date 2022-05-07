package main

import (
	"csock/config"
	"fmt"
)

func main() {
	fmt.Println("Initializing...")
	conf := config.LoadConfig()
	ctx := Context{Config: &conf}

	uiStart(&ctx)
}

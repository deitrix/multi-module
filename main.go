package main

import (
	"fmt"

	"github.com/deitrix/multi-module/hello"
	"github.com/deitrix/multi-module/world"
)

func main() {
	fmt.Println(world.Hello())
	fmt.Println(hello.World())
}

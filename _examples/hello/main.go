package main

import (
	"fmt"

	"github.com/hopesea/godub/v2/audioop"
)

func main() {
	e := audioop.NewError("Hello, world: %d", 100)
	fmt.Println(e.Error())
}

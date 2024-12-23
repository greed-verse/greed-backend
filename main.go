package main

import (
	"fmt"

	"github.com/greed-verse/greed/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Sprintf("Failed Setup: %s", err))
	}
}

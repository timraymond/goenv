package main

import (
	"fmt"
	"os"

	"github.com/timraymond/goenv/cmd/goenv/internal"
)

func main() {
	if err := internal.Run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

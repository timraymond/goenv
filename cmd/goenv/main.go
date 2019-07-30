package main

import (
	"fmt"
	"os"

	"github.com/timraymond/goenv/cmd/goenv/internal"
)

func main() {
	cmd := &internal.Command{
		Builder: &OSBuilder{
			DirMode:  os.ModePerm,
			FileMode: 0644,
		},
	}

	if err := internal.Run(os.Args[1:], cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

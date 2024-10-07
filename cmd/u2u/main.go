package main

import (
	"fmt"
	"os"

	"github.com/sesanetwork/go-sesa/cmd/sesa/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

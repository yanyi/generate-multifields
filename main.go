package main

import (
	"github.com/yanyi/generate-multifields/cmd"
	"github.com/yanyi/generate-multifields/internal/errwrapper"
)

func main() {
	if err := cmd.Execute(); err != nil {
		errwrapper.Fatal(err)
	}
}

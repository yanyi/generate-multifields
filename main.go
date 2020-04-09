package main

import (
	"github.com/yanyi/generate-multifields-gql/cmd"
	"github.com/yanyi/generate-multifields-gql/internal/errwrapper"
)

func main() {
	if err := cmd.Execute(); err != nil {
		errwrapper.Fatal(err)
	}
}

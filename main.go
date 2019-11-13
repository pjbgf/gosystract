package main

import (
	"io"
	"os"

	cli "github.com/pjbgf/gosystract/cmd/cli"
	"github.com/pjbgf/gosystract/cmd/systract"
)

func main() {
	run(os.Stdout, os.Args)
}

func run(output io.Writer, args []string) {
	err := cli.Run(output, args, systract.Extract)
	if err != nil {
		panic(err)
	}
}

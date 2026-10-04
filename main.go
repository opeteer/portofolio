package main

import (
	"os"

	"github.com/opeteer/opeteer-cli/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}

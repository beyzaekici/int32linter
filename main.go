package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"mylinter/linter"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}

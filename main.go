package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"int32linter/linter"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}

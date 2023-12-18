package main

import (
	"github.com/beyzaekici/int32linter/linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}

package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"github.com/beyzaekici/int32linter/linter"
)

func main() {
	singlechecker.Main(linter.Analyzer)
}

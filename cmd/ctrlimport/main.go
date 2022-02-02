package main

import (
	"github.com/sakedog/ctrlimport"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctrlimport.Analyzer) }

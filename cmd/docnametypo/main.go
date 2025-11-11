package main

import (
	"github.com/cce/docnametypo/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(analyzer.Analyzer) }

package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/cloverrose/connectnew"
)

func main() { unitchecker.Main(connectnew.Analyzer) }

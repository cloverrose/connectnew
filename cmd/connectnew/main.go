package main

import (
	"github.com/cloverrose/connectnew"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(connectnew.Analyzer) }

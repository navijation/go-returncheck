package main

import (
	returncheck "github.com/navijation/go-returncheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(returncheck.Analyzer)
}

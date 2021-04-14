package main

import (
	"github.com/algolia/go/playground/lint/recvcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(recvcheck.Analyzer)
}
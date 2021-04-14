package main

import (
	"github.com/benjbaron/recvcheck-lint/recvcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(recvcheck.Analyzer)
}

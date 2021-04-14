package main

import (
	"github.com/benjbaron/recvcheck-lint/recvcheck"
	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin analyzerPlugin

type analyzerPlugin struct{}

func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		recvcheck.Analyzer,
	}
}

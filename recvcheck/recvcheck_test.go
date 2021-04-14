package recvcheck_test

import (
	"testing"

	"github.com/algolia/go/playground/lint/recvcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRecvCheck(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), recvcheck.Analyzer, "recvcheck")
}

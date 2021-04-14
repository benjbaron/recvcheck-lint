package recvcheck_test

import (
	"testing"

	"github.com/benjbaron/recvcheck-lint/recvcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRecvCheck(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), recvcheck.Analyzer, "recvcheck")
}

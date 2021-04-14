package recvcheck

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Analyzer declares the recvcheck analyzer.
var Analyzer = &analysis.Analyzer{
	Name: "recvcheck",
	Doc:  "reports invalid receiver names",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fd, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			if fd.Recv == nil {
				return true
			}

			if len(fd.Recv.List) == 0 {
				return true
			}

			if len(fd.Recv.List[0].Names) == 0 {
				return true
			}

			recvType := types.ExprString(fd.Recv.List[0].Type)
			recvName := fd.Recv.List[0].Names[0].Name
			methodName := fd.Name.Name

			recvTypeLetter := recvType[0]
			if recvTypeLetter == '*' {
				recvTypeLetter = recvType[1]
			}

			l := strings.ToLower(string(recvTypeLetter))
			if l == recvName {
				return true
			}

			pass.Reportf(fd.Pos(), "invalid receiver name %q of type %q for method %q, should be %q",
				recvName, recvType, methodName, l)
			return true
		})
	}

	return nil, nil
}

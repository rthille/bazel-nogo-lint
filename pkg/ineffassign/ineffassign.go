package ineffassign

import (
	"go/token"

	ineffassignAPI "github.com/golangci/ineffassign"

	"github.com/anuvu/bazel-nogo-lint/pkg/util"
	"golang.org/x/tools/go/analysis"
)

const Name = "ineffassign"

var Analyzer = &analysis.Analyzer{
	Name: Name,
	Doc:  "Detects when assignments to existing variables are not used",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	issues := ineffassignAPI.Run(util.GetAllFileNames(pass.Files))
	if len(issues) == 0 {
		return nil, nil
	}
	for _, i := range issues {
		pass.Reportf(token.Pos(i.Pos.Offset), "[%s] ineffectual assignment to %s", Name, util.FormatCode(i.IdentName))
	}
	return nil, nil
}

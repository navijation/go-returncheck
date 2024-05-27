package returncheck

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:  "returncheck",
	Doc:   "Go analyzer that requires function return values be used or explicitly ignored with `_`",
	URL:   "https://github.com/navijation/go-returncheck",
	Flags: flag.FlagSet{},
	Run: func(*analysis.Pass) (interface{}, error) {
		return nil, nil
	},
	// TODO: it should actually be possible to run this analyzer in spite of type errors, although
	// certain features like comment directives on function definitions will not work correctly
	RunDespiteErrors: false,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	ResultType:       nil,
	// TODO: facts are necessary to support comment directives on function and struct definitions;
	// but we will support those later
	FactTypes: []analysis.Fact{},
}

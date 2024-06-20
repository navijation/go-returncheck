package analyzer_test

import (
	"path/filepath"
	"testing"

	"github.com/navijation/go-returncheck/analyzer"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

var testdataPath, _ = filepath.Abs("./testdata/") //nolint:gochecknoglobals

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	a := analyzer.NewAnalyzer()

	require.NotNil(t, a)

	analysistest.Run(t, testdataPath, a, "v1")
}

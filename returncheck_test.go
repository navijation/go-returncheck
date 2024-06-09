package returncheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestReturncheck(t *testing.T) {
	t.Run("v1 test cases, default configuration", func(t *testing.T) {
		t.Skip("enabled once v1 is ready for release")
		analysistest.Run(t, analysistest.TestData(), Analyzer, "v1/...")
	})
}

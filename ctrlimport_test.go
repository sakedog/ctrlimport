package ctrlimport_test

import (
	"testing"

	"github.com/sakedog/ctrlimport"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	setupFlags(t, "ng",
		"github.com/google/go-cmp/...",
		"net/http/httptest",
	)

	setupFlags(t, "ok",
		"github.com/google/go-cmp/cmp",
	)

	analysistest.RunWithSuggestedFixes(t, testdata, ctrlimport.Analyzer, "a")
}

func setupFlags(t *testing.T, flag string, paths ...string) {
	t.Helper()

	for _, path := range paths {
		if err := ctrlimport.Analyzer.Flags.Set(flag, path); err != nil {
			t.Fatal(err)
		}
	}
}

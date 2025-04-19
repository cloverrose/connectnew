package connectnew

import (
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gostaticanalysis/testutil"
)

func Test(t *testing.T) {
	t.Parallel()
	testdata := analysistest.TestData()
	testdata = testutil.WithModules(t, testdata, nil)
	pkgs := "a"
	analysistest.Run(t, testdata, Analyzer, strings.Split(pkgs, ",")...)
}

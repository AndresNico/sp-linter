package annotation_test

import (
	"testing"

	"github.com/AndresNico/sp-linter/analyzers/annotation"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnnotationCheck(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), annotation.Analyzer)
}

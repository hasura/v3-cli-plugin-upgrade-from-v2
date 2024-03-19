package analysis_test

// This test package performs analysis on test_data/*.metadata.json
// files, performs analysis on them, then verifies the results against
// matching test_data/*.analysis.json files.
//
// A diff of the expected and performed analysis is performed by the
// github.com/wI2L/jsondiff package and reported to the user as a
// test error if there are differences found.

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wI2L/jsondiff"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/analysis"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
)

func compareAnalysis(t *testing.T, filePath, expectedPath string) {
	metadata := util.ReadJSON(filePath)
	expectedFeatures := util.ReadJSON(expectedPath)

	reportdata := report.ReportData{Metadata: metadata}
	analysis.Analysis(false, &reportdata) // Set debugging to true if desired

	patch, err := jsondiff.Compare(expectedFeatures, reportdata.CheckList)
	if err != nil {
		panic(fmt.Sprintf("Error comparing analysis: %s", err))
	}

	if patch != nil {
		patchBytes, err := json.MarshalIndent(patch, "", "    ")
		if err != nil {
			panic(fmt.Sprintf("Error outputting patch: %s", err))
		}
		t.Errorf("Analysis of metadata (%s) did not match expecte value: %s", filePath, string(patchBytes))
	}
}

func TestAnalysis(t *testing.T) {

	files, err := filepath.Glob("../test_data/*.metadata.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, filePath := range files {
		expectedPath := strings.Replace(filePath, "metadata.json", "analysis.json", 1)

		testname := fmt.Sprintf("Analysis test: %s", filePath)
		t.Run(testname, func(t *testing.T) {
			compareAnalysis(t, filePath, expectedPath)
		})
	}
}

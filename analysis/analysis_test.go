package analysis_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wI2L/jsondiff"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/analysis"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
)

func readJSON(filePath string) map[string]interface{} {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %s: %s", filePath, err))
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling JSON %s: %s", filePath, err))
	}

	return jsonData
}

func TestAnalysis(t *testing.T) {

	files, err := filepath.Glob("../test_data/*.metadata.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, filePath := range files {
		expectedPath := strings.Replace(filePath, "metadata.json", "analysis.json", 1)

		fmt.Printf("%s -> %s\n", filePath, expectedPath)

		metadata := readJSON(filePath)
		expectedFeatures := readJSON(expectedPath)

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
}

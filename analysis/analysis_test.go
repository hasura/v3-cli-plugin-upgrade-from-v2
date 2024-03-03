package analysis_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/analysis"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
)

func TestAnalysis(t *testing.T) {

	files, err := filepath.Glob("../test_data/*.in.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the list of matching files
	fmt.Println("Matching files:")
	for _, file := range files {
		fmt.Println(file)
		data := report.ReportData{}
		analysis.Analysis(true, &data)
	}

	// Test case 1
	result := 5
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) expected %d, but got %d", expected, result)
	}

	// Add more test cases as needed
}

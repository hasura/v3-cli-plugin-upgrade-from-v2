package analysis

import (
	"fmt"
	"upgrade-from-v2/features"
	"upgrade-from-v2/report"
)

// Mutates the CheckList struct to set features that are used
func Analysis(data *report.ReportData) {
	fmt.Println("Debugging Analysis...")
	Actions(&data.CheckList, data.Metadata)
}

func Actions(features *features.Checklist, metadata map[string]interface{}) {
	fmt.Println("Debugging Actions...")
	md, exists := metadata["metadata"]
	if !exists {
		return
	}
	fmt.Println("Debugging Metadata...")
	switch v := md.(type) {
	case map[string]interface{}:
		_, exists := v["actions"]
		if !exists {
			return
		}
		fmt.Println("Found actions")
		features.Actions.UsesActions = true
	default:
		return
	}
}

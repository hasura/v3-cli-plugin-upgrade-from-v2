package analysis

import (
	"upgrade-from-v2/features"
	"upgrade-from-v2/report"
)

// Mutates the CheckList struct to set features that are used
func Analysis(data *report.ReportData) {
	md, exists := data.Metadata["metadata"]
	if !exists {
		panic("Invalid Metadata Format")
	}

	switch v := md.(type) {
	case map[string]interface{}:
		Actions(&data.CheckList, v)
	default:
		panic("Invalid Metadata Format")
	}
}

func Actions(features *features.Checklist, metadata map[string]interface{}) {
	_, exists := metadata["actions"]
	if !exists {
		return
	}
	features.Actions.UsesActions = true
}

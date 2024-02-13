package analysis

import (
	"upgrade-from-v2/features"
	"upgrade-from-v2/report"
)

// Mutates the CheckList struct to set features that are used
func Analysis(data *report.ReportData) {
	md, exists := data.Metadata["metadata"]
	if !exists {
		panic("Invalid Metadata Format: Missing `metadata` key.")
	}

	switch v := md.(type) {
	case map[string]interface{}:
		Actions(&data.CheckList, v)
	default:
		panic("Invalid Metadata Format: `metadata` field is not an object.")
	}
}

func Actions(features *features.Checklist, metadata map[string]interface{}) {
	actions, exists := metadata["actions"]
	if !exists {
		return
	}
	features.Actions.UsesActions = true

	switch v := actions.(type) {
	case []interface{}:
		for _, a := range v {
			switch w := a.(type) {
			case map[string]interface{}:
				Kriti(features, w)
			default:
				panic("Invalid Metadata Format: `action` is not an object.")
			}
		}
	default:
		panic("Invalid Metadata Format: `actions` field is not an array.")
	}
}

func Kriti(features *features.Checklist, action map[string]interface{}) {
	definition, exists := action["definition"]
	if !exists {
		panic("Action missing definition.")
	}
	switch d := definition.(type) {
	case map[string]interface{}:
		_, response_transform_exists := d["response_transform"]
		if response_transform_exists {
			features.Actions.UsesKriti = true
		}
		_, request_transform_exists := d["request_transform"]
		if request_transform_exists {
			features.Actions.UsesKriti = true
		}
	default:
		panic("Action definition is not an object.")
	}
}

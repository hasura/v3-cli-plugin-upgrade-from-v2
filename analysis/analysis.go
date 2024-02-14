package analysis

import (
	"fmt"
	"os"
	"text/template"
	"upgrade-from-v2/report"
)

// Mutates the CheckList struct to set features that are used
// Template based analysis - Sets flags in checklist of features
func Analysis(debugging bool, data *report.ReportData) {

	// Functions
	usesActions := func(name string) string {
		data.CheckList.Actions.UsesActions = true
		return fmt.Sprintf("Action: %s", name)
	}
	usesKriti := func(name string) string {
		data.CheckList.Actions.UsesKriti = true
		return fmt.Sprintf("Kriti: %s", name)
	}
	usesSources := func(name string) string {
		data.CheckList.Sources.UsesSources = true
		return fmt.Sprintf("Source: %s", name)
	}
	usesFromEnv := func(name string) string {
		data.CheckList.Sources.UsesFromEnv = true
		return fmt.Sprintf("FromENV: %s", name)
	}
	// TODO: Create a generic "usesFeature" somehow
	// usesFeature := func(tag string, path ...string) string { ...
	var funcs = template.FuncMap{
		"usesActions": usesActions,
		"usesKriti":   usesKriti,
		"usesFromEnv": usesFromEnv,
		"usesSources": usesSources,
	}

	// Parse template
	const templatePath = "./analysis/analysis.md"
	t1 := template.New("analysis.md").Funcs(funcs)
	t2, e2 := t1.ParseFiles(templatePath)
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse template file %s: %s", templatePath, e2))
	}

	// Execute template
	var e3 error
	if debugging {
		e3 = t2.Execute(os.Stdout, data) // If debugging
	} else {
		nullFile, err := os.Create(os.DevNull)
		if err != nil {
			panic(err)
		}
		e3 = t2.Execute(nullFile, data)
	}
	if e3 != nil {
		panic(fmt.Sprintf("Error executing template file %s: %s", templatePath, e3))
	}
}

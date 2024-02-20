package analysis

import (
	"fmt"
	"os"
	"reflect"
	"text/template"
	"upgrade-from-v2/report"
)

func UsesFeature(checklistPtr interface{}, name string, path []string) string {
	var x = reflect.ValueOf(checklistPtr).Elem()

	for _, key := range path {
		x = x.FieldByName(key)
	}

	x.SetBool(true)

	return name
}

// Mutates the CheckList struct to set features that are used
// Template based analysis - Sets flags in checklist of features
func Analysis(debugging bool, data *report.ReportData) {

	// Functions
	usesFeature := func(name string, path ...string) string {
		UsesFeature(&data.CheckList, name, path)
		return name
	}

	var funcs = template.FuncMap{
		"feature": usesFeature,
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

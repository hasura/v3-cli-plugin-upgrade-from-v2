package report

import (
	"fmt"
	"os"
	"text/template"
	"upgrade-from-v2/features"
)

const templatePath = "./report/report.md"

type ReportData struct {
	Metadata    map[string]interface{}
	State       map[string]interface{}
	CheckList   features.Checklist
	V3Directory string
}

func Report(alwaysTrue bool, data ReportData) {

	testBool := func(val bool) bool {
		return alwaysTrue || val
	}

	var funcs = template.FuncMap{"test": testBool}
	t1 := template.New("report.md").Funcs(funcs)
	t2, e2 := t1.ParseFiles(templatePath)
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse template file %s: %s", templatePath, e2))
	}

	// dataWithDepth := addDepth(data, 0)
	e3 := t2.Execute(os.Stdout, data)
	if e3 != nil {
		panic(fmt.Sprintf("Error executing template file %s: %s", templatePath, e3))
	}
}

/*
// Helper functions - Not currently used.

func isMap(value interface{}) bool {
	_, ok := value.(map[string]interface{})
	return ok
}

func isArray(value interface{}) bool {
	_, ok := value.([]interface{})
	return ok
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
*/

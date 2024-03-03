package report

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/features"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/writers"
)

//go:embed *.md
var reportTemplateFS embed.FS

const templatePath = "report.md"

type ReportData struct {
	Metadata    map[string]interface{}
	State       map[string]interface{}
	Summary     map[string]int
	CheckList   features.Checklist
	V3Directory string
}

func Report(alwaysTrue bool, data ReportData) {

	testBool := func(val bool) bool {
		return alwaysTrue || val
	}

	var funcs = template.FuncMap{"test": testBool}
	t1 := template.New(templatePath).Funcs(funcs)
	t2, e2 := t1.ParseFS(reportTemplateFS, templatePath)
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse report template: %s", e2))
	}

	w := writers.NewRemoveDuplicateBlankLinesWriter(os.Stdout)

	e3 := t2.Execute(w, data)
	if e3 != nil {
		panic(fmt.Sprintf("Error executing report template: %s", e3))
	}
}

package analysis

import (
	"embed"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/writers"
)

func UsesFeature(checklistPtr interface{}, path []string) {
	var x = reflect.ValueOf(checklistPtr).Elem()

	for _, key := range path {
		x = x.FieldByName(key)
	}

	x.SetBool(true)
}

//go:embed *.md
var analysisTemplateFS embed.FS

const templatePath = "analysis.md"

// Mutates the CheckList struct to set features that are used
// Template based analysis - Sets flags in checklist of features
func Analysis(debugging bool, data *report.ReportData) {

	// Functions
	usesFeature := func(path ...string) string {
		UsesFeature(&data.CheckList, path)
		return strings.Join(path[:], ".")
	}

	regexTest := func(pattern string, value string) bool {
		re := regexp.MustCompile(pattern)
		return re.MatchString(value)
	}

	var funcs = template.FuncMap{
		"feature": usesFeature,
		"regex":   regexTest,
	}

	// Parse template
	t1 := template.New(templatePath).Funcs(funcs)
	t2, e2 := t1.ParseFS(analysisTemplateFS, templatePath)
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse analysis templates: %s", e2))
	}

	// Execute template
	var e3 error
	if debugging {
		w := writers.NewRemoveDuplicateBlankLinesWriter(os.Stderr)
		e3 = t2.Execute(w, data) // If debugging
	} else {
		nullFile, err := os.Create(os.DevNull)
		if err != nil {
			panic(err)
		}
		e3 = t2.Execute(nullFile, data)
	}
	if e3 != nil {
		panic(fmt.Sprintf("Error executing analysis template: %s", e3))
	}
}

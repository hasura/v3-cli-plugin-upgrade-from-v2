package report

import (
	"fmt"
	"math/rand"
	"os"
	"text/template"
	"upgrade-from-v2/features"
)

const templatePath = "./report/report.md"

type ReportData struct {
	Metadata map[string]interface{}
	State    map[string]interface{}
	Features []features.Feature
}

func Report(data ReportData) {
	t1 := template.New("report.md").Funcs(funcs)
	t2, e2 := t1.ParseFiles(templatePath)
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse template file %s: %s", templatePath, e2))
	}

	// dataWithDepth := addDepth(data, 0)
	t2.Execute(os.Stdout, data)
}

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

func addDepth(data interface{}, depth int) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			result[key] = map[string]interface{}{"Depth": depth, "Value": value, "Children": addDepth(value, depth+1)}
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, elem := range v {
			result[i] = addDepth(elem, depth)
		}
		return result
	default:
		return data
	}
}

var funcs = template.FuncMap{"isMap": isMap, "isArray": isArray, "randomBool": randomBool}

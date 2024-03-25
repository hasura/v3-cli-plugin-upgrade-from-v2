package generate

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed templates/*.go.*
var templateFS embed.FS

func Generate(location_template string, content_template_path string, data map[string]interface{}) error {
	t1 := template.New(content_template_path)

	t2, e2 := t1.ParseFS(templateFS, fmt.Sprintf("templates/%s", content_template_path))
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse report template: %s", e2))
	}

	e3 := t2.Execute(os.Stdout, data)

	fmt.Println(e3)
	// TODO: Return something

	return nil
}

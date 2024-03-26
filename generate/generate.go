package generate

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.go.*
var templateFS embed.FS

func Generate(location_template_string string, content_template_path string, data map[string]interface{}) error {

	// LOCATION

	location_template := template.New("location-template")

	tL2, eL2 := location_template.Parse(location_template_string)
	if eL2 != nil {
		panic(fmt.Sprintf("Couldn't parse location template: %s", eL2))
	}
	var locationBuffer bytes.Buffer

	if eL3 := tL2.Execute(&locationBuffer, data); eL3 != nil {
		panic(eL3)
	}

	location := locationBuffer.String()

	// CONTENT

	// Create parent directories
	if err := os.MkdirAll(filepath.Dir(location), 0755); err != nil {
		return err
	}

	// File handle
	fo, err := os.Create(location)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	t1 := template.New(content_template_path)

	t2, e2 := t1.ParseFS(templateFS, fmt.Sprintf("templates/%s", content_template_path))
	if e2 != nil {
		panic(fmt.Sprintf("Couldn't parse report template: %s", e2))
	}

	e3 := t2.Execute(fo, data)

	if e3 != nil {
		panic(e3)
	}

	return nil
}

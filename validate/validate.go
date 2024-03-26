package validate

import (
	"embed"
	"fmt"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schemas/*.openapi.json
var schemaFS embed.FS

func RunValidator(metadata_path string) {
	schemaBytes, err := schemaFS.ReadFile("schemas/metadata.openapi.json")
	if err != nil {
		panic(err)
	}

	// Register callbacks before compilation for validated collection
	jsonschema.Callbacks["generate"] = func(t string, s string, o map[string]interface{}) {
		fmt.Println(t)
		fmt.Println(s)
		// fmt.Println(util.FormatJSON(o))
	}

	sch, err := jsonschema.CompileString("https://github.com/hasura/graphql-engine/blob/main/metadata.openapi.json", string(schemaBytes))
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}

	md := util.ReadJSON(metadata_path)
	errV := sch.Validate(md["metadata"])

	if errV != nil {
		panic(fmt.Sprintf("%#v", errV))
	}
}

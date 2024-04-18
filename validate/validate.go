package validate

import (
	"embed"
	"fmt"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schemas/*.openapi.json
var schemaFS embed.FS

func RunValidatorPath(metadata_path string) {
	md := util.ReadJSON(metadata_path)
	RunValidator(md)
}

func RunValidator(md map[string]interface{}) {
	schemaBytes, err := schemaFS.ReadFile("schemas/metadata.openapi.json")
	if err != nil {
		panic(err)
	}

	sch, err := jsonschema.CompileString("https://github.com/hasura/graphql-engine/blob/main/metadata.openapi.json", string(schemaBytes))
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}

	errV := sch.Validate(md["metadata"])

	if errV != nil {
		panic(fmt.Sprintf("%#v", errV))
	}
}

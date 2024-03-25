/*
TODO: This module is only required for testing during development
*/
package cmd

import (
	"embed"
	"fmt"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi-validator/schema_validation"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		createValidator()
	},
}

//go:embed schemas/*.openapi.json
var schemaFS embed.FS

var metadata string

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVar(&metadata, "metadata", "", "Hasura V2 Metadata Document")
	validateCmd.MarkFlagRequired("metadata")
}

func createValidator() {
	schema, err := schemaFS.ReadFile("schemas/metadata.openapi.json")
	md := util.ReadJSON(metadata)

	if err != nil {
		panic(err)
	}

	// 2. Create a new OpenAPI document using libopenapi
	document, docErrs := libopenapi.NewDocument(schema)

	if docErrs != nil {
		panic(docErrs)
	}

	m, modelErrs := document.BuildV3Model()

	if modelErrs != nil {
		panic(modelErrs)
	}

	sch := m.Model.Components.Schemas.GetOrZero(`Metadata`).Schema()

	v := schema_validation.NewSchemaValidator()

	valid, errors := v.ValidateSchemaObject(sch, md)

	for e := range errors {
		fmt.Println(e)
	}
	if !valid {
		panic(`Invalid metadata`)
	}

	panic(`This is a stub command.`)
}

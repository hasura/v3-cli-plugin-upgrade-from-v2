/*
TODO: This module is only required for testing during development
*/
package cmd

import (
	"embed"
	"fmt"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		runValidator()
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

func runValidator() {
	schemaBytes, err := schemaFS.ReadFile("schemas/metadata.openapi.json")
	if err != nil {
		panic(err)
	}

	sch, err := jsonschema.CompileString("https://github.com/hasura/graphql-engine/blob/main/metadata.openapi.json", string(schemaBytes))
	if err != nil {
		panic(fmt.Sprintf("%#v", err))
	}

	jsonschema.Callback = func(t string, s string, o map[string]interface{}) {
		fmt.Println(t)
		fmt.Println(s)
		fmt.Println(util.FormatJSON(o))
	}

	md := util.ReadJSON(metadata)
	errV := sch.Validate(md["metadata"])

	if errV != nil {
		panic(fmt.Sprintf("%#v", errV))
	}
}

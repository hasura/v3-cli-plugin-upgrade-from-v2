/*
TODO: This module is only required for testing during development
*/
package cmd

import (
	"fmt"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/validate"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {

		// Register callbacks before compilation for validated collection
		jsonschema.Callbacks["generate"] = test_validate
		jsonschema.Callbacks["report"] = test_validate

		validate.RunValidatorPath(metadata)
	},
}

var metadata string

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVar(&metadata, "metadata", "", "Hasura V2 Metadata Document")
	validateCmd.MarkFlagRequired("metadata")
}

func test_validate(t string, s string, o map[string]interface{}) {
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(len(o))
}

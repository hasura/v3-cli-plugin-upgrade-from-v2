/*
TODO: This module is only required for testing during development
*/
package cmd

import (
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/validate"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		validate.RunValidator(metadata)
	},
}

var metadata string

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVar(&metadata, "metadata", "", "Hasura V2 Metadata Document")
	validateCmd.MarkFlagRequired("metadata")
}

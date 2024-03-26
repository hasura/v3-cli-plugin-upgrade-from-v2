package cmd

import (
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Hasura V3 subgraph.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func run() {
	generate.Generate("foo/bar.yaml", "test.go.yaml", map[string]interface{}{})
	panic(`generate command is currently a work in progress`)
}

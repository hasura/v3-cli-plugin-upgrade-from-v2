package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/v2api"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/validate"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Hasura V3 subgraph.",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// E.g. cd ~/hasura/v3-cli-iteration/app && npx ts-node index.ts internal upgrade-from-v2 \
// --projectDir "$1" --sm "$2" -s "$3" \
// --metadata "$4" --tag "$5" --path "$6"
var generatorCommand string
var subgraphLocation string
var supergraphManifestName string
var projectDir string

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().StringVar(&generatorCommand, "v3-generator", os.Getenv("HASURA_V3_GENERATOR"), "Hasura V3 Generator Command")
	generateCmd.PersistentFlags().StringVar(&subgraphLocation, "v3-subgraph-location", os.Getenv("HASURA_V3_SUBGRAPH_LOCATION"), "DDN Subgraph Location")
	generateCmd.PersistentFlags().StringVar(&supergraphManifestName, "v3-supergraph-manifest-name", os.Getenv("HASURA_V3_SUPERGRAPH_MANIFEST_NAME"), "Name of the supergraph manifest (default: base)")
	generateCmd.PersistentFlags().StringVar(&projectDir, "v3-project-directory", os.Getenv("HASURA_V3_PROJECT_DIRECTORY"), "Hasura V3 Project Directory (default: $HASURA_V3_SUBGRAPH_LOCATION/..)")

	generateCmd.MarkPersistentFlagRequired("v3-generator")         // TODO: Set a default for this using hasura CLI?
	generateCmd.MarkPersistentFlagRequired("v3-subgraph-location") // TODO: Set a default for this using hasura CLI?
}

func run() {

	metadata := v2api.FetchV2Metadata(V2URL, V2AdminSecret)

	jsonschema.Callbacks["report"] = do_nothing
	jsonschema.Callbacks["generate"] = invoke_generator

	validate.RunValidator(metadata)

	// Using the CLI Iterator project instead of custom templating
	// generate.Generate("foo/bar.yaml", "test.go.yaml", map[string]interface{}{})
	panic(`generate command is currently a work in progress`)
}

func do_nothing(t string, s string, o map[string]interface{}) {
}

func invoke_generator(t string, s string, o map[string]interface{}) {

	if supergraphManifestName == "" {
		supergraphManifestName = "base"
	}

	if projectDir == "" {
		// Assume that subgraphs can't be nested deeper than this.
		projectDir = fmt.Sprintf("%s/..", subgraphLocation)
	}

	cmd := exec.Command("/bin/sh", "-c", generatorCommand, "_", projectDir, supergraphManifestName, subgraphLocation, util.FormatJSON(o), t, s)
	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	if err != nil {
		panic(err.Error())
	}
}

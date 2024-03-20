// Initially generated with `cobra init.`

package cmd

import (
	"os"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "v3-cli-plugin-upgrade-from-v2",
	Short: "Upgrade Hasura V2 projects to V3",
	Long:  `Report on the compatibility of Hasura V2 projects, then generate a compatible Hasura V3 project.`,
}

var (
	Debugging     bool
	V2URL         string
	V2AdminSecret string
)

func init() {
	// "Global" flags
	rootCmd.PersistentFlags().BoolVar(&Debugging, "debug", util.GetenvBool("DEBUG"), "Enable debugging output (DEBUG)")
	rootCmd.PersistentFlags().StringVar(&V2URL, "v2-url", os.Getenv("HASURA_V2_URL"), "Project URL for the V2 project (HASURA_V2_URL)")
	rootCmd.PersistentFlags().StringVar(&V2AdminSecret, "v2-admin-secret", os.Getenv("HASURA_V2_ADMIN_SECRET"), "Hasura V2 project admin-secret (HASURA_V2_ADMIN_SECRET)")

	// Since there is no ENV integration for cobra, we have to manually check that the env flag isn't being used before setting required flags
	// NOTE: We could wrap this up in helper functions if the need becomes common
	if os.Getenv("HASURA_V2_URL") == "" {
		rootCmd.MarkPersistentFlagRequired("v2-url")
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

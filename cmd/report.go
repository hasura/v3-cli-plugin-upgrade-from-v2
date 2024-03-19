package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/analysis"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/features"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/util"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/v2api"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a V3 compatiblity report for a V2 project",
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

var (
	format           string
	detectEverything bool
)

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().StringVar(&format, "format", util.GetenvStringWithDefault("HASURA_REPORT_FORMAT", "md"), "Report format {md,json} (HASURA_REPORT_FORMAT)")
	reportCmd.Flags().BoolVar(&detectEverything, "debug-detect-everything", util.GetenvBool("DEBUG_DETECT_EVERYTHING"), "Debug: Detect Everything (DEBUG_DETECT_EVERYTHING)")
}

func main() {
	var reportData report.ReportData // To be set depending on protocol

	if strings.HasPrefix(V2URL, "http") {
		metadata := v2api.FetchV2Metadata(V2URL, V2AdminSecret)
		state := v2api.FetchV2InternalState(V2URL, V2AdminSecret)
		reportData = report.ReportData{
			Metadata:  metadata,
			State:     state,
			CheckList: features.List,
		}
	} else {
		metadata := util.ReadJSON(V2URL)
		reportData = report.ReportData{
			Metadata:  metadata,
			CheckList: features.List,
		}
	}

	analysis.Analysis(Debugging, &reportData)
	reportData.Summary = features.SumCategories(reportData.CheckList)

	fmt.Println("")
	switch format {
	case "md":
		report.Report(detectEverything, reportData)

	case "json":
		jsonData, err := json.Marshal(reportData.CheckList)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		fmt.Println(string(jsonData))

	default:
		panic(`Only md and json formats are supported`)
	}
	fmt.Println("")
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/analysis"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/features"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/report"
	"github.com/hasura/v3-cli-plugin-upgrade-from-v2/v2api"
)

type Option struct {
	ENV string
	CLI string
}

var OptionDebugging = Option{"DEBUG", "debug"}
var OptionDetectEverything = Option{"DEBUG_DETECT_EVERYTHING", "debug-detect-everything"}
var OptionV2URL = Option{"HASURA_V2_URL", "v2-url"}
var OptionV2AdminSecret = Option{"HASURA_V2_ADMIN_SECRET", "v2-admin-secret"}
var OptionV2MetadataSchema = Option{"HASURA_V2_METADATA_SCHEMA", "v2-metadata-schema"}
var OptionV3Directory = Option{"HASURA_V3_DIRECTORY", "v3-directory"}

var (
	debugging        bool
	detectEverything bool
	v2URL            string
	v2AdminSecret    string
	v2MetadataSchema string
	v3Directory      string
)

func init() {
	// Command-line flags
	flag.StringVar(&v2URL, OptionV2URL.CLI, os.Getenv(OptionV2URL.ENV), "V2 Project URL")
	flag.StringVar(&v2AdminSecret, OptionV2AdminSecret.CLI, os.Getenv(OptionV2AdminSecret.ENV), "V2 Admin Secret")
	flag.StringVar(&v2MetadataSchema, OptionV2MetadataSchema.CLI, os.Getenv(OptionV2MetadataSchema.ENV), "V2 Metadata Schema")
	flag.StringVar(&v3Directory, OptionV3Directory.CLI, os.Getenv(OptionV3Directory.ENV), "V3 Directory")
	flag.BoolVar(&debugging, OptionDebugging.CLI, os.Getenv(OptionDebugging.ENV) != "", "Debug Mode")
	flag.BoolVar(&detectEverything, OptionDetectEverything.CLI, os.Getenv(OptionDetectEverything.ENV) != "", "Debug Detect Everything")

	// Parse command-line flags
	flag.Parse()

	if v2URL == "" || v2AdminSecret == "" || v3Directory == "" {
		errorMsg, _ := fmt.Printf("Please ensure you have set the following options [ENV] %s [%s], %s [%s], %s [%s] \n", OptionV2URL.CLI, OptionV2URL.ENV, OptionV2AdminSecret.CLI, OptionV2AdminSecret.ENV, OptionV3Directory.CLI, OptionV3Directory.ENV)
		panic(errorMsg)
	}
}

func main() {
	metadata := v2api.FetchV2Metadata(v2URL, v2AdminSecret)
	state := v2api.FetchV2InternalState(v2URL, v2AdminSecret)
	reportData := report.ReportData{
		Metadata:    metadata,
		State:       state,
		CheckList:   features.List,
		V3Directory: v3Directory,
	}

	analysis.Analysis(debugging, &reportData)
	reportData.Summary = features.SumCategories(reportData.CheckList)

	fmt.Println("")
	report.Report(detectEverything, reportData)
	fmt.Println("")

	// // Disabled due to this library seemingly having issues with OpenAPIV3
	// if v2MetadataSchema != "" {
	// 	foo(v2MetadataSchema, metadata)
	// }
}

// func foo(schemaLocation string, doc map[string]interface{}) {

// 	var schemaURI = schemaLocation
// 	if !strings.Contains(schemaLocation, "://") {
// 		schemaURI = fmt.Sprintf("file://%s", schemaLocation)
// 	}

// 	fmt.Println("")
// 	fmt.Println("")
// 	fmt.Println("DEBUGGING: ")
// 	fmt.Printf("Validating Metadata against OpenAPI Schema: %s \n", schemaURI)
// 	fmt.Println("")
// 	fmt.Println("")

// 	schemaLoader := gojsonschema.NewReferenceLoader(schemaURI)
// 	// := gojsonschema.NewReferenceLoader("file:///home/me/document.json")
// 	documentLoader := gojsonschema.NewRawLoader(doc["metadata"])

// 	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	if result.Valid() {
// 		fmt.Printf("The document is valid\n")
// 	} else {
// 		fmt.Printf("The document is not valid. see errors :\n")
// 		for _, desc := range result.Errors() {
// 			fmt.Printf("- %s\n", desc)
// 		}
// 	}

// 	fmt.Println(doc)
// }

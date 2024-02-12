package main

import (
	"flag"
	"fmt"
	"os"
	"upgrade-from-v2/v2api"
)

type Option struct {
	ENV string
	CLI string
}

var Optionv2URL = Option{"HASURA_V2_URL", "v2-url"}
var Optionv2AdminSecret = Option{"HASURA_V2_ADMIN_SECRET", "v2-admin-secret"}
var Optionv3Directory = Option{"HASURA_V3_DIRECTORY", "v3-directory"}

var (
	v2URL         string
	v2AdminSecret string
	v3Directory   string
)

func init() {
	// Command-line flags
	flag.StringVar(&v2URL, Optionv2URL.CLI, os.Getenv(Optionv2URL.ENV), "V2 Project URL")
	flag.StringVar(&v2AdminSecret, Optionv2AdminSecret.CLI, os.Getenv(Optionv2AdminSecret.ENV), "V2 Admin Secret")
	flag.StringVar(&v3Directory, Optionv3Directory.CLI, os.Getenv(Optionv3Directory.ENV), "V3 Directory")

	// Parse command-line flags
	flag.Parse()

	if v2URL == "" || v2AdminSecret == "" || v3Directory == "" {
		errorMsg, _ := fmt.Printf("Please ensure you have set the following options [ENV] %s [%s], %s [%s], %s [%s] \n", Optionv2URL.CLI, Optionv2URL.ENV, Optionv2AdminSecret.CLI, Optionv2AdminSecret.ENV, Optionv3Directory.CLI, Optionv3Directory.ENV)
		panic(errorMsg)
	}
}

func main() {
	fmt.Printf("V2 URL: %s\n", v2URL)
	fmt.Printf("V3 Directory: %s\n", v3Directory)
	fmt.Println(v2api.FetchV2Info(v2URL, v2AdminSecret))
}

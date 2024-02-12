package main

import (
	"fmt"
	"os"
	"upgrade-from-v2/v2api"
)

// TODO: Support CLI flags, etc.

const ENVv2Project = "HASURA_V2_PROJECT"
const ENVapiKey = "HASURA_V2_ADMIN_SECRET"
const ENVv3Project = "HASURA_V3_PROJECT"

func main() {

	v2Project := os.Getenv(ENVv2Project)
	apiKey := os.Getenv(ENVapiKey)
	v3Project := os.Getenv(ENVv3Project)

	if v2Project == "" || apiKey == "" || v3Project == "" {
		errorMsg, _ := fmt.Printf("Please ensure you have set ENV variables %s, %s, %s \n", ENVv2Project, ENVapiKey, ENVv3Project)
		panic(errorMsg)
	}

	fmt.Println(v2api.FetchV2Info())
}

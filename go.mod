module github.com/hasura/v3-cli-plugin-upgrade-from-v2

go 1.21.4

// replace github.com/santhosh-tekuri/jsonschema/v5 => /Users/lyndon/code/jsonschema-go
replace github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 => github.com/sordina/jsonschema-go/v5 v5.3.3

require (
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1
	github.com/spf13/cobra v1.8.0
	github.com/wI2L/jsondiff v0.5.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
)

# V3 CLI Plugin: Upgrade V2 to V3 [WIP]

This is a Hasura V3 CLI plugin. It attempts to create a compatible V3 project based off a V2 project.

Usage:

```
hasura3 upgrade-from-v2 report --v2-project https://myproject.cloud.hasura.com --admin-secret "$HASURA_ADMIN_SECRET" --v3-project
```

Design:

This plugin takes a running Hasura V2 project and interrogates it for the following information:

* Metadata
* Schema Introspection
* Telemetry and Metrics

Then it outputs:

* A compatibility report
* A V3 project that attempts to be as compatible as possible


## Distribution of the plugin

Github Actions CI builds and distributes releases of this plugin:

* CI invokes the `Makefile` to build the artifacts and manifest
* Uploads the results to GCloud bucket with credentials specificed in GH actions
* Creates a PR to the plugins index repo


## TODO

* [x] Fetch All V2 Project information in v2api package
  * [x] Metadata
  * [x] Schema Introspection Information (via "dump_internal_state")
* [x] Encode Feature matrix
* [x] Perform Analysis
* [x] Output report
* [x] Analysis golden tests
* [ ] Output V3 Template
* [ ] Update to new V3 format
* [ ] Schema diff
* [ ] Create TS implementations for actions
* [ ] Create TS lib for Kriti

## Creating an HTML Report

Create markdown:

HASURA_V2_URL=https://foobar.hasura.app HASURA_V2_ADMIN_SECRET=XXX hasura3 upgrade-from-v2 report > report.md


## Developing and Publishing

You can develop and use this application as a standalone go CLI app until you are ready to distribute and publish your changes.

Once you are happy with your changes, you should update the CLI plugins index here: https://github.com/hasura/cli-plugins-index

Development loop:

```sh
HASURA_V2_PROJECT=https://XYZ.hasura.app HASURA_V2_ADMIN_SECRET=XXX make dev
```


## Testing

Run:

```
make test
```

If you wish to update test output you can simply delete the previous `test_data/*.analysis.json` file.


## Links

* CLI Plugins Index: https://github.com/hasura/cli-plugins-index


## Appendix

Golang app created with generator-go and yeoman.

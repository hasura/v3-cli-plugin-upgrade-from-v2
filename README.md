# V3 CLI Plugin: Upgrade V2 to V3 [WIP]

This is a Hasura V3 CLI plugin. It attempts to create a compatible V3 project based off a V2 project.

Usage:

```
hasura3 upgrade-from-v2 --v2-project https://myproject.cloud.hasura.com --admin-secret "$HASURA_ADMIN_SECRET" --v3-project .
```

Design:

This plugin takes a running Hasura V2 project and interrogates it for the following information:

* Metadata
* Schema Introspection
* Telemetry and Metrics

Then it outputs:

* A compatibility report
* A V3 project that attempts to be as compatible as possible

It will read configuration from the v3 project directory: `.upgrade-from-v2.yml` or create this file with defaults if not present.


## TODO

* [ ] Fetch All V2 Project information in v2api package
* [ ] Encode Feature matrix
* [ ] Output report
* [ ] Output V3 Template
* [ ] Update to new V3 format
* [ ] Create TS implementations for actions
* [ ] Create TS lib for Kriti

## Developing and Publishing

You can develop and use this application as a standalone go CLI app until you are ready to distribute and publish your changes.

Once you are happy with your changes, you should update the CLI plugins index here: https://github.com/hasura/cli-plugins-index

Development loop:

```sh
HASURA_V2_PROJECT=https://XYZ.hasura.app HASURA_V2_ADMIN_SECRET=XXX HASURA_V3_PROJECT=. make dev
```


## Links

* CLI Plugins Index: https://github.com/hasura/cli-plugins-index


## Appendix

Golang app created with generator-go and yeoman.

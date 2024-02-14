
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

## Features

These are the features used by your project:

{{if .CheckList.Actions.UsesActions }}
* Actions [Supported] {{end -}}
{{if .CheckList.Actions.UsesKriti}}
  * Kriti [Unsupported]
    Kriti is currently unsupported, but you can write your own request/response transformations with the TypeScript connector...
{{end}}

## Executive Summary

Your project is partially supported. In order to use V3 you will need to write some custom TS.

The following features that you are using are not currently supported:

* TODO

{{if .V3Directory }}
A copy of this report has been saved to {{.V3Directory}}/v2_upgrade_report.md
{{else}}
To save a copy of this report specify --v3-directory [V3_DIRECTORY]
{{end}}
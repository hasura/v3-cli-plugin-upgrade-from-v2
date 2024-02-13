
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

{{/*

{{define "recur"}}
  {{if isMap .}}
    Map:
    {{range $key, $value := . }}
      * K: {{ $key }}
      * V: {{template "recur" $value}}
    {{end}}
  {{else if isArray .}}
    Array:
    {{range $item := .}}
      {{template "recur" $item}}
    {{end}}
  {{else if .}}
    Value:
      {{.}}
  {{end}}
{{end}}


## Metadata

{{range $key, $value := .Metadata.metadata -}}
  * {{$key}}
{{end}}


## Internal State

{{range $key, $value := .State -}}
  * {{$key}}
{{end}}

### Actions

{{range $item := .Metadata.metadata.actions -}}
  * {{$item}}
{{end}}

*/}}

## Features

These are the features used by your project:

{{if .CheckList.Actions.UsesActions }}
* Actions [Supported] {{end -}}
{{if .CheckList.Actions.UsesKriti}}
  * Kriti [Unsupported] {{end -}}

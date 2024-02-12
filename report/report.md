
# V2 to V3 Upgrade Report

This report outlines the compatibility of your v2 project with Hasura V3.

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

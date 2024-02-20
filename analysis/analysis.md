
# DEBUGGING: Analysis of Features

This template is not used for user-facing output but will be displayed if DEBUG=true.

The go templating features are used for dynamic traversal of map[string]interface{} data structures.

{{range $action := .Metadata.metadata.actions -}}
  * {{ feature $action.name "Actions" "Used" }}
  {{if $action.definition.request_transform -}}
    * {{ feature "request_transform" "Kriti" "Used" }}
  {{- end }}
  {{if $action.definition.response_transform -}}
    * {{ feature "response_transform" "Kriti" "Used" }}
  {{- end }}
{{end}}

{{range $source := .Metadata.metadata.sources -}}
  * {{feature $source.name "Sources" "Used"}}
  {{if $source.configuration.connection_info.database_url.from_env -}}
    * {{feature $source.configuration.connection_info.database_url.from_env "Sources" "FromEnv" }}
  {{end}}
  {{if eq $source.kind "postgres" -}}
    * {{feature $source.name "Sources" "PG" }}
  {{end}}
{{end}}

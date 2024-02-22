
# DEBUGGING: Analysis of Features

This template is not used for user-facing output but will be displayed if DEBUG=true.

The go templating features are used for dynamic traversal of map[string]interface{} data structures.

{{range $source := .Metadata.metadata.sources -}}
  * {{$source.name}}: {{feature "Sources" "Used"}}
  {{if $source.configuration.connection_info.database_url.from_env -}}
    * {{feature "Sources" "FromEnv" }}
  {{end}}
  {{if eq $source.kind "postgres" -}}
    * {{feature "Sources" "PG" }}
  {{end}}
{{end}}

{{range $action := .Metadata.metadata.actions -}}
  * Action {{$action.name}}: {{ feature "Actions" "Used" }}
  {{if $action.definition.request_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * request_transform: {{ feature "Actions" "Transforms" "RequestTransforms" "Used" }}
  {{- end }}
  {{if $action.definition.response_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * response_transform: {{ feature "Actions" "Transforms" "ResponseTransforms" "Used" }}
  {{- end }}
{{end}}
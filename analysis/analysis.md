
# DEBUGGING: Analysis of Features

This template is not used for user-facing output but will be displayed if DEBUG=true.

The go templating features are used for dynamic traversal of map[string]interface{} data structures.

{{range $action := .Metadata.metadata.actions -}}
  * {{ usesActions $action.name }}
  {{if $action.definition.request_transform -}}
    * {{ usesKriti "request_transform" }}
  {{- end }}
  {{if $action.definition.response_transform -}}
    * {{ usesKriti "response_transform" }}
  {{- end }}
{{end}}

{{.Metadata.metadata.sources}}

{{range $source := .Metadata.metadata.sources -}}
  {{usesSources $source.name}}
  {{if $source.configuration.connection_info.database_url.from_env -}}
    * {{usesFromEnv $source.configuration.connection_info.database_url.from_env }}
  {{end}}
  {{if eq $source.kind "postgres" -}}
    * {{usesPG $source.name }}
  {{end}}
{{end}}

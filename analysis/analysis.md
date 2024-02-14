
# DEBUGGING: Analysis of Features

This template is not used for user-facing output

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

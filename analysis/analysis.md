
# DEBUGGING: Analysis of Features

This template is not used for user-facing output but will be displayed if DEBUG=true.

The go templating features are used for dynamic traversal of map[string]interface{} data structures.

## Sources

{{range $source := .Metadata.metadata.sources -}}
  * {{$source.name}}: {{feature "Sources" "Used"}}
  {{if $source.configuration.connection_info.database_url.from_env -}}
    * {{feature "Sources" "FromEnv" }}
  {{end}}

  {{if eq $source.kind "postgres" -}}
    * {{feature "Sources" "PG" }}
  {{end}}
  {{if eq $source.kind "pg" -}}
    * {{feature "Sources" "PG" }}
  {{end}}
  {{if eq $source.kind "citus" -}}
    * {{feature "Sources" "Citus" }}
  {{end}}
  {{if eq $source.kind "cockroach" -}}
    * {{feature "Sources" "Cockroach" }}
  {{end}}
  {{if eq $source.kind "mssql" -}}
    * {{feature "Sources" "SQLServer" }}
  {{end}}
  {{if eq $source.kind "bigquery" -}}
    * {{feature "Sources" "BigQuery" }}
  {{end}}

  ### Tables

  Since only PG is currently supported, we just check according to `PostgresTableMetadata`.
  {{range $table := $source.tables}}
    * {{$table.table}}
    {{if $table.configuration.comment}}
      * {{feature "Tables" "CommentsOnModels" }}
    {{end}}
    {{if $table.is_enum}}
      * {{feature "Tables" "EnumTables" }}
    {{end}}
    {{if $table.apollo_federation_config}}
      * {{feature "Tables" "ApolloFederation" }}
    {{end}}

    ### Relationships

    {{range $rel := $table.array_relationships}}
      {{feature "Relationships" "Used"}}
      {{if $rel.using.foreign_key_constraint_on}}
        {{feature "Relationships" "LocalRelationshipWithForeignKey"}}
      {{else}}
        {{feature "Relationships" "LocalRelationshipWithoutForeignKey"}}
      {{end}}
    {{end}}

    {{range $rel := $table.object_relationships}}
      {{feature "Relationships" "Used"}}
      {{if $rel.using.foreign_key_constraint_on}}
        {{feature "Relationships" "LocalRelationshipWithForeignKey"}}
      {{else}}
        {{feature "Relationships" "LocalRelationshipWithoutForeignKey"}} (TODO: Check if this is right)
      {{end}}
    {{end}}

  {{end}}

  TODO: Dataconnector sources, such as SQLite, Mongo, etc.
{{end}}

## Remote Schemas

{{range $schema := .Metadata.metadata.remote_schemas }}
  {{feature "RemoteSchemas" "Used"}}

  Is it really worth stepping through all the RS features, when none of them are available?

  TODO: Step through config, permissions, etc.
{{end}}



## Actions

{{range $action := .Metadata.metadata.actions -}}
  * Action {{$action.name}}: {{ feature "Actions" "Used" }}

	{{if eq $action.definition.type "query"}}
		* Query {{$action.name}}: {{ feature "Actions" "Queries" }}
	{{end}}
	{{if eq $action.definition.type "mutation"}}
		* Mutation {{$action.name}}: {{ feature "Actions" "Mutations" }}
	{{end}}

	### Types

  TODO: Check if there is more nuance around this than I think
	* All actions use types: {{ feature "Actions" "Types" "Used" }}
		* {{ feature "Actions" "Types" "CustomTypes" }}
		* {{ feature "Actions" "Types" "CustomScalarTypes" }}

  {{if $action.definition.request_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * request_transform: {{ feature "Actions" "Transforms" "RequestTransforms" "Used" }}
  {{- end }}
  {{if $action.definition.response_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * response_transform: {{ feature "Actions" "Transforms" "ResponseTransforms" "Used" }}
  {{- end }}

	### HTTP Configuration

  {{if $action.definition.request_transform -}}
    * {{ feature "Actions" "HttpConfiguration" "Used" }}
    * {{ feature "Actions" "HttpConfiguration" "ConfiguredEndpoints" }}

		{{if $action.definition.request_transform.url -}}
			* {{ feature "Actions" "HttpConfiguration" "URLTemplating" }}
		{{end}}
		{{if ne $action.definition.request_transform.method "POST" -}}
			* {{ feature "Actions" "HttpConfiguration" "RequestMethod" }}
		{{end}}
  {{- end }}

  {{range $header := $action.definition.headers -}}
		{{ if $header.value}}
			{{ feature "Actions" "HttpConfiguration" "StaticHeaders"}}
		{{end}}
		{{ if $header.value_from_env }}
			{{ feature "Actions" "HttpConfiguration" "DynamicHeaders"}}
		{{end}}
	{{end}}

	{{if $action.definition.forward_client_headers}}
		{{ feature "Actions" "HttpConfiguration" "ForwardClientHeaders"}}
	{{end}}

	### Async Actions

	{{if eq $action.definition.kind "asynchronous" }}
		{{ feature "Actions" "AsyncActions" }}
	{{end}}

{{end}}
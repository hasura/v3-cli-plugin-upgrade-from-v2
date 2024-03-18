
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
		{{feature "Tables" "Used" }}
    * Note: CommentsOnModels data isn't in metadata

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
				Not sure if this gets triggered.
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

    {{range $rel := $table.remote_relationships}}
			{{feature "Relationships" "RemoteRelationships"}}
    {{end}}
  {{end}}

  TODO: Dataconnector sources, such as SQLite, Mongo, etc.
{{end}}

## Remote Schemas

{{range $schema := .Metadata.metadata.remote_schemas }}
  {{feature "RemoteSchemas" "Used"}}

  Is it really worth stepping through all the RS features, when none of them are available?

	{{if $schema.definition.url_from_env}}
		{{feature "RemoteSchemas" "Configuration" "Used"}}
		{{feature "RemoteSchemas" "Configuration" "FromEnv"}}
	{{end}}

	{{if ne $schema.definition.timeout_seconds 60.0}}
		Default Timeout is 60 seconds
		{{feature "RemoteSchemas" "Configuration" "Timeout"}}
	{{end}}

	{{range $header := $schema.definition.headers}}
		{{feature "RemoteSchemas" "Configuration" "Headers"}}
		{{if $header.value_from_env}}
			{{feature "RemoteSchemas" "Configuration" "DynamicHeaders"}}
		{{end}}
	{{end}}
{{end}}


## Queries

These can't be detected by metadata inspection.

Will need telemetry information.


## Mutations

Likewise, will need query information to mark these features.


## Logical Models and Native Queries

{{range $source := .Metadata.metadata.sources -}}
	{{range $model := $source.logical_models}}
		{{feature "LogicalModels" "Used"}}
		{{if $model.select_permissions}}
			This should be moved to logical models feature rather than NQs.
			{{feature "NativeQueries" "Queries" "Permissions"}}
		{{end}}
	{{end}}

	{{range $model := $source.native_queries}}
		{{feature "NativeQueries" "Used"}}
	{{end}}
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

	### Permissions

	{{if $action.permissions }}
		{{ feature "Actions" "BasicPermissions" }}
	{{end}}

	### Relationships

	Note: This feature is present on custom types, not explicitly on actions.
	      Nevertheless, we put this in the actions loop for grouping.

	{{ range $relationship := .Metadata.metadata.custom_types.objects }}
		{{ feature "Actions" "Relationships" "Used" }}

		{{ if $relationship.remote_table }}
			{{ feature "Actions" "Relationships" "ActionToDB" }}
		{{end}}
	{{end}}

	* A2RS unsupported in V2
	* A2A unsupported in V2

	### Transforms

  {{if $action.definition.request_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * {{ feature "Actions" "Transforms" "RequestTransforms" "Used" }}
		{{if ne $action.definition.request_transform.method "POST" -}}
			* {{ feature "Actions" "Transforms"  "RequestTransforms" "RequestMethod" }}
		{{end}}
		{{if $action.definition.request_transform.url -}}
			* {{ feature "Actions" "Transforms" "RequestTransforms" "RequestURL" }}
		{{end}}
		{{if $action.definition.request_transform.body -}}
			* {{ feature "Actions" "Transforms" "RequestTransforms" "RequestBody" }}
		{{end}}
	{{end}}
  {{if $action.definition.response_transform -}}
    * {{ feature "Actions" "Transforms" "Used" }}
    * {{ feature "Actions" "Transforms" "ResponseTransforms" "Used" }}
		{{if $action.definition.response_transform.body -}}
			* {{ feature "Actions" "Transforms" "PayloadTransforms" "ResponseBodyTransforms" }}
		{{end}}
	{{end}}
{{end}}


## Event Triggers

Reiterating the sources, to find event triggers.

{{range $source := .Metadata.metadata.sources -}}
	{{range $table := $source.tables}}
		{{range $trigger := $table.event_triggers}}

			### Used

			* {{ feature "EventTriggers" "Used" }}

			### Types

			{{if $trigger.definition.insert}}
				* {{ feature "EventTriggers" "Types" "Insert" }}
			{{end}}
			{{if $trigger.definition.update}}
				* {{ feature "EventTriggers" "Types" "Update" }}
				{{if $trigger.definition.update.columns}}
					* {{ feature "EventTriggers" "Types" "ColumnSpecificUpdates" }}
				{{end}}
			{{end}}
			{{if $trigger.definition.delete}}
				* {{ feature "EventTriggers" "Types" "Delete" }}
			{{end}}
			{{if $trigger.definition.enable_manual}}
				* {{ feature "EventTriggers" "Types" "Console" }}
			{{end}}

			### Trigger Protocols?

			### RequestTransforms

			{{ if $trigger.request_transform}}
				* {{ feature "EventTriggers" "RequestTransforms" "Used" }}
				{{ if ne $trigger.request_transform.method "POST"}}
					* {{ feature "EventTriggers" "RequestTransforms" "RequestMethod" }}
				{{end}}
				{{ if $trigger.request_transform.url }}
					* {{ feature "EventTriggers" "RequestTransforms" "RequestURL" }}
				{{end}}
				{{ if $trigger.request_transform.body }}
					* {{ feature "EventTriggers" "RequestTransforms" "RequestBody" }}
					* {{ feature "EventTriggers" "PayloadTransform" }}
				{{end}}
			{{end}}

			{{ if $trigger.headers}}
				* {{ feature "EventTriggers" "RequestTransforms" "Headers" }}
			{{end}}

			### Response Transforms

			{{ if $trigger.response_transform}}
				* {{ feature "EventTriggers" "ResponseTransforms" }}
			{{end}}

			### Payload Transforms

			Same as request transform for body?

			### Retry Logic

			{{if gt $trigger.retry_conf.num_retries 0.0}}
				* {{ feature "EventTriggers" "RetryLogicConfiguration" }}
			{{end }}
		{{end}}
	{{end}}
{{end}}


## Authentication

Can't be detected.


## Authorization

Table based permissions:

{{range $source := .Metadata.metadata.sources -}}
	{{range $table := $source.tables}}
		{{range $permission := $table.insert_permissions}}
			{{feature "Authorization" "Used"}}
		{{end}}
	{{end}}
{{end}}

TODO: Other permission sources.


## RESTified Endpoints

{{ range $endpoint := .Metadata.metadata.rest_endpoints}}
	* {{ feature "RESTifiedEndpoints" "Used" }}
	* {{ feature "RESTifiedEndpoints" "BasicGraphql" }}
	{{if gt (len $endpoint.methods) 1}}
		* {{ feature "RESTifiedEndpoints" "DifferentHTTPMethods" }}
	{{end}}
{{end}}

## Allow Lists

{{ range $collection := .Metadata.metadata.query_collections}}
	{{if eq $collection.name "allowed-queries"}}
		* {{ feature "AllowLists" "Used" }}
		{{ range $query := $collection.definition.queries}}
			{{if regex "@cached" $query.query }}
				* {{ feature "AllowLists" "CachingMetrics" }}
			{{end}}
		{{end}}
	{{end}}
{{end}}


## API Limits

{{if .Metadata.metadata.api_limits}}
	{{feature "APILimits" "Used"}}
{{end}}
{{if .Metadata.metadata.rate_limit}}
	{{feature "APILimits" "RateLimits"}}
{{end}}
{{if .Metadata.metadata.depth_limit}}
	{{feature "APILimits" "DepthLimits"}}
{{end}}
{{if .Metadata.metadata.node_limit}}
	{{feature "APILimits" "NodeLimits"}}
{{end}}
{{if .Metadata.metadata.time_limit}}
	{{feature "APILimits" "TimeLimits"}}
{{end}}
{{if .Metadata.metadata.batch_limit}}
	{{feature "APILimits" "BatchLimits"}}
{{end}}


## DynamicEnvironmentVariables

This doesn't seem to be able to be detected.


## Subscriptions

This doesn't seem to be able to be detected.


## Observability

{{if eq .Metadata.metadata.opentelemetry.status "enabled"}}
	{{feature "Observability" "Used"}}
	{{feature "Observability" "OpenTelemetry"}}
{{end}}


# Analysis of Internal State

## Tables - Comments on models check:

{{range $source := .State.sources}}
	{{range $table := (index $source 1).tables }}
		{{ if $table.core_info.description }}
			FOOOO: {{feature "Tables" "CommentsOnModels" }}
		{{end}}
	{{end}}
{{end}}
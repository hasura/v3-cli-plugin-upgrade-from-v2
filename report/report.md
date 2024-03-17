
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

These are the features used by your project:


{{if test .CheckList.Sources.Used}}
## Sources
{{end}}

{{if test .CheckList.Sources.FromEnv}}
### FromENV [Supported]
{{end}}

{{if test .CheckList.Sources.PG}}
### Postgres [Supported]
{{end}}

{{if test .CheckList.Sources.Citus}}
### Citus [Unsupported]
{{end}}

{{if test .CheckList.Sources.Cockroach}}
### Cockroach [Unsupported]
{{end}}

{{if test .CheckList.Sources.SQLServer}}
### SQL Server [Unsupported]
{{end}}

{{if test .CheckList.Sources.BigQuery}}
### BigQuery [Unsupported]
{{end}}


{{if test .CheckList.Tables.Used}}
## Tables
{{end}}

{{if test .CheckList.Tables.CommentsOnModels}}
### Comments On Models [Supported]
{{end}}

{{if test .CheckList.Tables.EnumTables}}
### Enum Tables [Unsupported]
{{end}}

{{if test .CheckList.Tables.ApolloFederation}}
### Apollo Federation [Unsupported]
{{end}}


{{if test .CheckList.Relationships.Used}}
## Relationships
{{end}}

{{if test .CheckList.Relationships.LocalRelationshipWithForeignKey}}
### Local Relationships with Foreign Keys [Supported]
{{end}}

{{if test .CheckList.Relationships.LocalRelationshipWithoutForeignKey}}
### Local Relationships without Foreign Keys [Supported]
{{end}}

{{if test .CheckList.Relationships.RemoteRelationships}}
### Remote Relationships [Supported]
{{end}}


{{if test .CheckList.RemoteSchemas.Used}}
## Remote Schemas
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Used}}
### Configuration [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.FromEnv}}
#### FromEnv [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Timeout}}
#### Timeout [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Headers}}
#### Headers [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.DynamicHeaders}}
#### Dynamic Headers [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.Used}}
### Relationships [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.ToDatabase}}
#### To Databases [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.ToRemoteSchema}}
#### To Remote Schemas [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.ArgumentPresets}}
#### Argument Presets [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.Permissions}}
### Permissions [Unsupported]
{{end}}

{{if test .CheckList.RemoteSchemas.BypassAuth}}
### Bypass Auth [Unsupported]
{{end}}


{{if test .CheckList.Queries.Used}}
## Queries
{{end}}

{{if test .CheckList.Queries.QueryList}}
### Query List [Supported]
{{end}}

{{if test .CheckList.Queries.QueryByPrimaryKey}}
### Query by Primary Key [Supported]
{{end}}

{{if test .CheckList.Queries.DistinctOn}}
### Uses "DISTINCT ON" [Supported]
{{end}}

{{if test .CheckList.Queries.Where.Used}}
### Uses "WHERE" [Supported]
{{end}}

{{if test .CheckList.Queries.Where.BasicUsage}}
#### Basic Usage [Supported]
{{end}}

{{if test .CheckList.Queries.Where.WithBoolExpression}}
#### With Boolean Expressions [Supported]
{{end}}

{{if test .CheckList.Queries.Where.Operators}}
#### Operators [Partially Supported]
{{end}}

{{if test .CheckList.Queries.Where.AggregateExpressions}}
#### Aggregate Expressions [Unsupported]
{{end}}

{{if test .CheckList.Queries.Pagination}}
### Pagination [Supported]
{{end}}

{{if test .CheckList.Queries.Limit}}
### Limits [Supported]
{{end}}

{{if test .CheckList.Queries.Offset}}
### Offsets [Supported]
{{end}}

{{if test .CheckList.Queries.OrderBy}}
### Uses "ORDER BY" [Supported]
{{end}}

{{if test .CheckList.Queries.Aggregates}}
### Aggregates [Unsupported]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.Used}}
### Simple Object Queries [Supported]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.ScalarIntegerAndText}}
#### Scalar Integers and Text [Supported]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.ScalarJSON}}
#### Scalar JSON [Unsupported]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.NestedObjects}}
#### Nested Objects [Unsupported]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.AggregateNestedObjects}}
#### Aggregate Nested Objects [Unsupported]
{{end}}

{{if test .CheckList.Queries.BoolExpressions}}
### Boolean Expressions [Partially Supported]
{{end}}

{{if test .CheckList.Queries.Introspection}}
### Introspection [Supported]
{{end}}

{{if test .CheckList.Queries.Permissions.Used}}
### Permissions [Supported]
{{end}}

{{if test .CheckList.Queries.Permissions.ColumnPermisisons}}
#### Column Permisisons [Supported]
{{end}}

{{if test .CheckList.Queries.Permissions.PermissionRules}}
#### Permission Rules [?]
{{end}}


{{if test .CheckList.Mutations.Used}}
## Mutations [Unsupported]
{{end}}

{{if test .CheckList.Mutations.Simple}}
### Simple Mutations [Unsupported]
{{end}}

{{if test .CheckList.Mutations.Complex}}
### Complex Mutations [Unsupported]
{{end}}


{{if test .CheckList.LogicalModels.Used}}
## Logical Models [Supported]
{{end}}


{{if test .CheckList.NativeQueries.Used}}
## Native Queries
{{end}}

{{if test .CheckList.NativeQueries.Queries.Used}}
### Queries [Supported]
{{end}}

{{if test .CheckList.NativeQueries.Queries.Permissions}}
#### Permissions [Supported]
{{end}}

{{if test .CheckList.NativeQueries.Mutations.Used}}
### Mutations [Partially Supported]
{{end}}

{{if test .CheckList.NativeQueries.Mutations.Permissions}}
#### Permissions [Unsupported]
{{end}}


{{if test .CheckList.Actions.Used}}
## Actions
{{end}}

{{if test .CheckList.Actions.Queries}}
### Query Actions [?]
{{end}}

{{if test .CheckList.Actions.Mutations}}
### Mutation Actions [?]
{{end}}

{{if test .CheckList.Actions.Types.Used}}
### Types [?]
{{end}}

*Note: All V2 actions use custom types.*

{{if test .CheckList.Actions.Types.CustomTypes}}
#### Custom Types [?]
{{end}}

{{if test .CheckList.Actions.Types.CustomScalarTypes}}
#### Custom Scalar Types [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.Used}}
### Http Configuration [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.ConfiguredEndpoints}}
#### Endpoints [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.URLTemplating}}
#### URL Templating [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.RequestMethod}}
#### Request Method [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.StaticHeaders}}
#### Static Headers [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.DynamicHeaders}}
#### Dynamic Headers [?]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.ForwardClientHeaders}}
#### Forward Client Headers [?]
{{end}}

{{if test .CheckList.Actions.AsyncActions}}
### Async Actions [?]
{{end}}

{{if test .CheckList.Actions.BasicPermissions}}
### Basic Permissions [?]
{{end}}

{{if test .CheckList.Actions.Relationships.Used}}
### Relationships [?]
{{end}}

{{if test .CheckList.Actions.Relationships.ActionToDB}}
#### To Database [?]
{{end}}

{{if test .CheckList.Actions.Relationships.ActionToRS}}
#### To Remote Schema [?]
{{end}}

{{if test .CheckList.Actions.Transforms.Used}}
### Transforms [?]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.Used}}
#### Request Transforms [?]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.ContextObjects}}
##### Context Objects [?]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestMethod}}
##### Request Method [?]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestURL}}
##### Request URL [?]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestBody}}
##### Request Body [?]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.Used}}
#### Response Transforms [?]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.ContextObjects}}
##### Context Objects [?]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.DebuggingMode}}
##### Debugging Mode [?]
{{end}}

{{if test .CheckList.Actions.Transforms.PayloadTransforms.Used}}
#### Payload Transforms [?]
{{end}}

{{if test .CheckList.Actions.Transforms.PayloadTransforms.ResponseBodyTransforms}}
##### Response Body Transforms [?]
{{end}}

{{if test .CheckList.EventTriggers.Used}}
## Event Triggers
{{end}}

{{if test .CheckList.Authentication.Used}}
## Authentication
{{end}}

{{if test .CheckList.Authorization.Used}}
## Authorization
{{end}}

{{if test .CheckList.Relay.Used}}
## Relay
{{end}}

{{if test .CheckList.RESTifiedEndpoints.Used}}
## RESTified Endpoints
{{end}}

{{if test .CheckList.RESTifiedEndpoints.DifferentHTTPMethods}}
### Multiple HTTP Methods
{{end}}

{{if test .CheckList.AllowLists.Used}}
## Allow Lists
{{end}}

{{if test .CheckList.APILimits.Used}}
## API Limits
{{end}}

{{if test .CheckList.DynamicEnvironmentVariables.Used}}
## Dynamic Environment Variables
{{end}}

{{if test .CheckList.Subscriptions.Used}}
## Subscriptions
{{end}}

{{if test .CheckList.Observability.Used}}
## Obervability
{{end}}


# Executive Summary

The following categories of feature support were detected in your Hasura V2 project:

| Support | Count |
| --- | --- |
{{ range $k, $v := .Summary -}}
| {{$k}} | {{$v}} |
{{end}}

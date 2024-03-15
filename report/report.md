
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

These are the features used by your project:


{{if test .CheckList.Sources.Used}}
## Sources [Supported]
{{end}}

{{if test .CheckList.Sources.FromEnv}}
### FromENV [?]
{{end}}

{{if test .CheckList.Sources.PG}}
### Postgres [Supported]
{{end}}

{{if test .CheckList.Sources.SQLServer}}
### SQL Server [?]
{{end}}

{{if test .CheckList.Sources.MySQL}}
### MySQL [?]
{{end}}

{{if test .CheckList.Sources.Mongo}}
### MongoDB [?]
{{end}}


{{if test .CheckList.Tables.Used}}
## Tables
{{end}}

{{if test .CheckList.Tables.CommentsOnModels}}
### Comments On Models [?]
{{end}}

{{if test .CheckList.Tables.EnumTables}}
### Enum Tables [?]
{{end}}

{{if test .CheckList.Tables.ApolloFederation}}
### Apollo Federation [?]
{{end}}


{{if test .CheckList.Relationships.Used}}
## Relationships
{{end}}

{{if test .CheckList.Relationships.LocalRelationshipWithForeignKey}}
### Local Relationships with Foreign Keys [?]
{{end}}

{{if test .CheckList.Relationships.LocalRelationshipWithoutForeignKey}}
### Local Relationships without Foreign Keys [?]
{{end}}

{{if test .CheckList.Relationships.RemoteRelationships}}
### Remote Relationships [?]
{{end}}


{{if test .CheckList.Queries.Used}}
## Queries
{{end}}

{{if test .CheckList.Queries.QueryByPrimaryKey}}
### Query by Primary Key [?]
{{end}}

{{if test .CheckList.Queries.DistinctOn}}
### Uses "DISTINCT ON" [?]
{{end}}

{{if test .CheckList.Queries.Where.Used}}
### Uses "WHERE" [?]
{{end}}

{{if test .CheckList.Queries.Where.BasicUsage}}
#### Basic Usage [?]
{{end}}

{{if test .CheckList.Queries.Where.WithBoolExpression}}
#### With Boolean Expressions [?]
{{end}}

{{if test .CheckList.Queries.Where.Operators}}
#### Operators [?]
{{end}}

{{if test .CheckList.Queries.Where.AggregateExpressions}}
#### Aggregate Expressions [?]
{{end}}

{{if test .CheckList.Queries.Pagination}}
### Pagination [?]
{{end}}

{{if test .CheckList.Queries.Limit}}
### Limits [?]
{{end}}

{{if test .CheckList.Queries.Offset}}
### Offsets [?]
{{end}}

{{if test .CheckList.Queries.OrderBy}}
### Uses "ORDER BY" [?]
{{end}}

{{if test .CheckList.Queries.Aggregate}}
### Aggregates [?]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.Used}}
### Simple Object Queries [?]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.ScalarIntegerAndText}}
#### Scalar Integers and Text [?]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.ScalarJSON}}
#### Scalar JSON [?]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.NestedObjects}}
#### Nested Objects [?]
{{end}}

{{if test .CheckList.Queries.SimpleObjectQueries.AggregateNestedObjects}}
#### Aggregate Nested Objects [?]
{{end}}

{{if test .CheckList.Queries.BoolExpressions}}
### Boolean Expressions [?]
{{end}}

{{if test .CheckList.Queries.Introspection}}
### Introspection [?]
{{end}}

{{if test .CheckList.Queries.Permissions.Used}}
### Permissions [?]
{{end}}

{{if test .CheckList.Queries.Permissions.ColumnPermisisons}}
#### Column Permisisons [?]
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
## Native Queries [Partially Supported]
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

{{if test .CheckList.RemoteSchemas.Used}}
## Remote Schemas
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Used}}
### Configuration [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.FromEnv}}
#### FromEnv [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Timeout}}
#### Timeout [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.Headers}}
#### Headers [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Configuration.DynamicHeaders}}
#### Dynamic Headers [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.Used}}
### Relationships [?]
{{end}}

{{if test .CheckList.RemoteSchemas.Relationships.Used}}
#### Relationships [?]
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

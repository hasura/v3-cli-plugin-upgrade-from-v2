
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.


## Features

These are the features used by your project:


{{if test .CheckList.Sources.Used -}}
### Sources [Supported]

{{if test .CheckList.Sources.FromEnv -}}
#### FromENV [?]
{{- end}}

{{if test .CheckList.Sources.PG -}}
#### Postgres [Supported]
{{- end}}

{{if test .CheckList.Sources.SQLServer -}}
#### SQL Server [?]
{{- end}}

{{if test .CheckList.Sources.MySQL -}}
#### MySQL [?]
{{- end}}

{{if test .CheckList.Sources.Mongo -}}
#### MongoDB [?]
{{- end}}
{{end}}

{{if test .CheckList.Tables.Used -}}
### Tables

{{if test .CheckList.Tables.CommentsOnModels -}}
#### Comments On Models [?]
{{- end}}

{{if test .CheckList.Tables.EnumTables -}}
#### Enum Tables [?]
{{- end}}

{{if test .CheckList.Tables.ApolloFederation -}}
#### Apollo Federation [?]
{{- end}}
{{end}}

{{if test .CheckList.Relationships.Used -}}
### Relationships

{{if test .CheckList.Relationships.LocalRelationshipWithForeignKey -}}
#### Local Relationships with Foreign Keys [?]
{{- end}}

{{if test .CheckList.Relationships.LocalRelationshipWithoutForeignKey -}}
#### Local Relationships without Foreign Keys [?]
{{- end}}

{{if test .CheckList.Relationships.RemoteRelationships -}}
#### Remote Relationships [?]
{{- end}}
{{end}}

{{if test .CheckList.Queries.Used -}}
### Queries

{{if test .CheckList.QueriesAndMutations.QueryByPrimaryKey -}}
#### Query by Primary Key [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.DistinctOn -}}
#### Uses "DISTINCT ON" [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.UsesWhere -}}
#### Uses "WHERE" [?]

{{if test .CheckList.QueriesAndMutations.Where.BasicUsage -}}
##### Basic Usage [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Where.WithBoolExpression -}}
##### With Boolean Expressions [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Where.Operators -}}
##### Operators [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Where.AggregateExpressions -}}
##### Aggregate Expressions [?]
{{- end}}
{{- end}}

{{if test .CheckList.QueriesAndMutations.Pagination -}}
#### Pagination [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Limit -}}
#### Limits [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Offset -}}
#### Offsets [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.OrderBy -}}
#### Uses "ORDER BY" [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Aggregate -}}
#### Aggregates [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.SimpleObjectQueries.Used -}}
#### Simple Object Queries [?]

{{if test .CheckList.QueriesAndMutations.SimpleObjectQueries.ScalarIntegerAndText -}}
##### Scalar Integers and Text [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.SimpleObjectQueries.ScalarJSON -}}
##### Scalar JSON [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.SimpleObjectQueries.NestedObjects -}}
##### Nested Objects [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.SimpleObjectQueries.AggregateNestedObjects -}}
##### Aggregate Nested Objects [?]
{{- end}}
{{- end}}

{{if test .CheckList.QueriesAndMutations.BoolExpressions -}}
#### Boolean Expressions [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Introspection -}}
#### Introspection [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.RunSQL -}}
#### RunSQL [?]

It may not be possible to detect usage of this feature...
{{- end}}

{{if test .CheckList.QueriesAndMutations.Permissions.Used -}}
#### Permissions [?]

{{if test .CheckList.QueriesAndMutations.Permissions.ColumnPermisisons -}}
##### Column Permisisons [?]
{{- end}}

{{if test .CheckList.QueriesAndMutations.Permissions.PermissionRules -}}
##### Permission Rules [?]
{{- end}}
{{- end}}
{{end}}

{{if test .CheckList.Mutations.Used -}}
### Mutations [Unsupported]

{{if test .CheckList.Mutations.Simple -}}
#### Simple Mutations [Unsupported]
{{- end}}

{{if test .CheckList.Mutations.Complex -}}
#### Complex Mutations [Unsupported]
{{- end}}
{{- end}}

{{if test .CheckList.LogicalModels.Used -}}
### Logical Models [Supported]
{{end}}

{{if test .CheckList.NativeQueries.Used -}}
### Native Queries [Partially Supported]

{{if test .CheckList.NativeQueries.Queries.Used -}}
#### Queries [Supported]

{{if test .CheckList.NativeQueries.Queries.Permissions -}}
##### Permissions [Supported]
{{- end}}
{{- end}}

{{if test .CheckList.NativeQueries.Mutations.Used -}}
#### Mutations [Partially Supported]

{{if test .CheckList.NativeQueries.Mutations.Permissions -}}
##### Permissions [Unsupported]
{{- end}}
{{- end}}
{{end}}

{{if test .CheckList.Actions.Used -}}
### Actions

{{if test .CheckList.Actions.Queries -}}
#### Query Actions [?]
{{- end}}

{{if test .CheckList.Actions.Mutations -}}
#### Mutation Actions [?]
{{- end}}

{{if test .CheckList.Actions.Types.Used -}}
#### Types [?]

*Note: All V2 actions use custom types.*

{{if test .CheckList.Actions.Types.CustomTypes -}}
##### Custom Types [?]
{{- end}}

{{if test .CheckList.Actions.Types.CustomScalarTypes -}}
##### Custom Scalar Types [?]
{{- end}}
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.Used -}}
#### Http Configuration [?]

{{if test .CheckList.Actions.HttpConfiguration.ConfiguredEndpoints -}}
##### Endpoints [?]
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.URLTemplating -}}
##### URL Templating [?]
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.RequestMethod -}}
##### Request Method [?]
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.StaticHeaders -}}
##### Static Headers [?]
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.DynamicHeaders -}}
##### Dynamic Headers [?]
{{- end}}

{{if test .CheckList.Actions.HttpConfiguration.ForwardClientHeaders -}}
##### Forward Client Headers [?]
{{- end}}
{{- end}}

{{if test .CheckList.Actions.AsyncActions -}}
#### Async Actions [?]
{{- end}}

{{if test .CheckList.Actions.BasicPermissions -}}
#### Basic Permissions [?]
{{- end}}

{{if test .CheckList.Actions.Relationships.Used -}}
#### Relationships [?]

{{if test .CheckList.Actions.Relationships.ActionToDB -}}
##### To Database [?]
{{- end}}

{{if test .CheckList.Actions.Relationships.ActionToRS -}}
##### To Remote Schema [?]
{{- end}}
{{- end}}

{{if test .CheckList.Actions.Transforms.Used -}}
#### Transforms [?]

{{if test .CheckList.Actions.Transforms.RequestTransforms.Used -}}
##### Request Transforms [?]

{{if test .CheckList.Actions.Transforms.RequestTransforms.ContextObjects -}}
###### Context Objects [?]
{{- end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestMethod -}}
###### Request Method [?]
{{- end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestURL -}}
###### Request URL [?]
{{- end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestBody -}}
###### Request Body [?]
{{- end}}
{{- end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.Used -}}
##### Response Transforms [?]

{{if test .CheckList.Actions.Transforms.ResponseTransforms.ContextObjects -}}
###### Context Objects [?]
{{- end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.DebuggingMode -}}
###### Debugging Mode [?]
{{- end}}
{{- end}}

{{if test .CheckList.Actions.Transforms.PayloadTransforms.Used -}}
##### Payload Transforms [?]

{{if test .CheckList.Actions.Transforms.PayloadTransforms.ResponseBodyTransforms -}}
###### Response Body Transforms [?]
{{- end}}
{{- end}}
{{- end}}
{{end}}

{{if test .CheckList.RemoteSchemas.Used -}}
### Remote Schemas

{{if test .CheckList.RemoteSchemas.Configuration.Used -}}
#### Configuration [?]

{{if test .CheckList.RemoteSchemas.Configuration.FromEnv -}}
##### FromEnv [?]
{{- end}}

{{if test .CheckList.RemoteSchemas.Configuration.Timeout -}}
##### Timeout [?]
{{- end}}

{{if test .CheckList.RemoteSchemas.Configuration.Headers -}}
##### Headers [?]
{{- end}}

{{if test .CheckList.RemoteSchemas.Configuration.DynamicHeaders -}}
##### Dynamic Headers [?]
{{- end}}
{{- end}}

{{if test .CheckList.RemoteSchemas.Relationships.Used -}}
#### Relationships [?]

{{if test .CheckList.RemoteSchemas.Relationships.Used -}}
##### Relationships [?]
{{- end}}
{{- end}}
{{end}}

{{if test .CheckList.EventTriggers.Used -}}
### Event Triggers
{{end}}

{{if test .CheckList.Authentication.Used -}}
### Authentication
{{end}}

{{if test .CheckList.Authorization.Used -}}
### Authorization
{{end}}

{{if test .CheckList.Relay.Used -}}
### Relay
{{end}}

{{if test .CheckList.RESTifiedEndpoints.Used -}}
### RESTified Endpoints

{{if test .CheckList.RESTifiedEndpoints.DifferentHTTPMethods -}}
#### Multiple HTTP Methods
{{end}}
{{end}}

{{if test .CheckList.AllowLists.Used -}}
### Allow Lists
{{end}}

{{if test .CheckList.APILimits.Used -}}
### API Limits
{{end}}

{{if test .CheckList.DynamicEnvironmentVariables.Used -}}
### Dynamic Environment Variables
{{end}}

{{if test .CheckList.Subscriptions.Used -}}
### Subscriptions
{{end}}

{{if test .CheckList.Observability.Used -}}
### Obervability
{{end}}


## Executive Summary

The following categories of feature support were detected in your Hasura V2 project:

| Support | Count |
| --- | --- |
{{ range $k, $v := .Summary -}}
| {{$k}} | {{$v}} |
{{end}}
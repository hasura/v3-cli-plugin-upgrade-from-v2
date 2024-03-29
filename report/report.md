
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
### Query Actions [Supported]
{{end}}

{{if test .CheckList.Actions.Mutations}}
### Mutation Actions [Supported]
{{end}}

{{if test .CheckList.Actions.Types.Used}}
### Types [Supported]
{{end}}

{{if test .CheckList.Actions.Types.CustomTypes}}
#### Custom Types [Supported]

*Note: All V2 actions use custom types.*
{{end}}

{{if test .CheckList.Actions.Types.CustomScalarTypes}}
#### Custom Scalar Types [Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.Used}}
### Http Configuration [Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.ConfiguredEndpoints}}
#### Endpoints [Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.URLTemplating}}
#### URL Templating [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.RequestMethod}}
#### Request Method [Partially Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.StaticHeaders}}
#### Static Headers [Partially Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.DynamicHeaders}}
#### Dynamic Headers [Partially Supported]
{{end}}

{{if test .CheckList.Actions.HttpConfiguration.ForwardClientHeaders}}
#### Forward Client Headers [Unsupported]
{{end}}

{{if test .CheckList.Actions.AsyncActions}}
### Async Actions [Unsupported]
{{end}}

{{if test .CheckList.Actions.BasicPermissions}}
### Basic Permissions [Supported]
{{end}}

{{if test .CheckList.Actions.Relationships.Used}}
### Relationships [Supported]
{{end}}

{{if test .CheckList.Actions.Relationships.ActionToDB}}
#### To Database [Supported]
{{end}}

{{if test .CheckList.Actions.Relationships.ActionToRS}}
#### To Remote Schema [Unsupported]
{{end}}

{{if test .CheckList.Actions.Relationships.ActionToAction}}
#### To Action [Supported]
{{end}}

{{if test .CheckList.Actions.Transforms.Used}}
### Transforms [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.Used}}
#### Request Transforms [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.ContextObjects}}
##### Context Objects [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestMethod}}
##### Request Method [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestURL}}
##### Request URL [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.RequestTransforms.RequestBody}}
##### Request Body [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.Used}}
#### Response Transforms [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.ContextObjects}}
##### Context Objects [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.ResponseTransforms.DebuggingMode}}
##### Debugging Mode [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.PayloadTransforms.Used}}
#### Payload Transforms [Manual Configuration Required]
{{end}}

{{if test .CheckList.Actions.Transforms.PayloadTransforms.ResponseBodyTransforms}}
##### Response Body Transforms [Manual Configuration Required]
{{end}}


{{if test .CheckList.EventTriggers.Used}}
## Event Triggers [Unsupported]

Further breakdown of the feature is not included as there is currently no V3 support for event triggers.
{{end}}


{{if test .CheckList.Authentication.Used}}
## Authentication
{{end}}

{{if test .CheckList.Authentication.JWTs.Used}}
### JWTs [Supported]
{{end}}

{{if test .CheckList.Authentication.JWTs.Secret}}
#### Secret [Supported]
{{end}}

{{if test .CheckList.Authentication.JWTs.JWK}}
#### JWK [Supported]
{{end}}

{{if test .CheckList.Authentication.Webhook}}
### Webhook [Supported]
{{end}}

{{if test .CheckList.Authentication.AdminRole}}
### Admin Role [Supported]
{{end}}

{{if test .CheckList.Authentication.UnauthorisedAccess}}
### Unauthorised Access [?]
{{end}}

{{if test .CheckList.Authentication.MultipleAdminSecrets}}
### Multiple Admin Secrets [Unsupported]
{{end}}

{{if test .CheckList.Authentication.MultipleJWTs}}
### Multiple JWTs [Unsupported]
{{end}}


{{if test .CheckList.Authorization.Used}}
## Authorization
{{end}}

{{if test .CheckList.Authorization.RoleFromSessionVariable}}
### Role From Session Variable [Supported]
{{end}}

{{if test .CheckList.Authorization.RoleBasedSchemeGeneration}}
### Role Based Scheme Generation [?]
{{end}}

{{if test .CheckList.Authorization.ConfigureRowPermissions}}
### Configure Row Permissions [Supported]
{{end}}

{{if test .CheckList.Authorization.ColumnPermissions}}
### Column Permissions [Supported]
{{end}}

{{if test .CheckList.Authorization.AggregationPermissions}}
### Aggregation Permissions [Unsupported]
{{end}}

{{if test .CheckList.Authorization.RowFetchLimit}}
### Row Fetch Limit [?]
{{end}}

{{if test .CheckList.Authorization.RootFieldVisibility}}
### Root Field Visibility [Supported]
{{end}}

{{if test .CheckList.Authorization.ColumnPresets}}
### Column Presets [?]
{{end}}

{{if test .CheckList.Authorization.BackendOnlyMutations}}
### Backend Only Mutations [Unsupported]
{{end}}

{{if test .CheckList.Authorization.PermissionsOperators}}
### Permissions Operators [Partially Supported]
{{end}}


{{if test .CheckList.Relay.Used}}
## Relay
{{end}}

{{if test .CheckList.Relay.Queries.Used}}
### Queries [Supported]
{{end}}

{{if test .CheckList.Relay.Queries.GlobalID}}
#### GlobalID [Supported]
{{end}}

{{if test .CheckList.Relay.Queries.ConnectionObjectsNodesEdges}}
#### Connection Objects Nodes Edges [?]
{{end}}

{{if test .CheckList.Relay.Queries.ConnectionSpec}}
#### ConnectionSpec [Unsupported]
{{end}}

{{if test .CheckList.Relay.Subscriptions}}
### Subscriptions [Unsupported]
{{end}}

{{if test .CheckList.Relay.Mutations}}
### Mutations [Unsupported]
{{end}}


{{if test .CheckList.RESTifiedEndpoints.Used}}
## RESTified Endpoints [Unsupported]
{{end}}

{{if test .CheckList.RESTifiedEndpoints.BasicGraphql}}
### Basic Graphql [Unsupported]
{{end}}

{{if test .CheckList.RESTifiedEndpoints.ArgumentsFromPayloadBody}}
### Arguments From Payload Body [Unsupported]
{{end}}

{{if test .CheckList.RESTifiedEndpoints.ArgumentsFromURLParams}}
### Arguments From URL Params [Unsupported]
{{end}}

{{if test .CheckList.RESTifiedEndpoints.DifferentHTTPMethods}}
### Multiple HTTP Methods [Unsupported]
{{end}}


{{if test .CheckList.AllowLists.Used}}
## Allow Lists [Unsupported]
{{end}}

{{if test .CheckList.AllowLists.ConfigureQueryCollectionsWithRoleAccess}}
## Configure Query Collections With Role Access [Unsupported]
{{end}}

{{if test .CheckList.AllowLists.MultipleQueryColectionsAndRoles}}
## Multiple Query Colections And Roles [Unsupported]
{{end}}

{{if test .CheckList.AllowLists.CachingMetrics}}
## Caching Metrics [Unsupported]
{{end}}

{{if test .CheckList.AllowLists.EnterpriseSuppport.Used}}
## Enterprise Suppport [Unsupported]
{{end}}

{{if test .CheckList.AllowLists.EnterpriseSuppport.CustomRedisConfigurations}}
## Custom Redis Configurations [Unsupported]
{{end}}


{{if test .CheckList.APILimits.Used}}
## API Limits [Unsupported]
{{end}}

{{if test .CheckList.APILimits.RateLimits}}
## Rate Limits [Unsupported]
{{end}}

{{if test .CheckList.APILimits.DepthLimits}}
## Depth Limits [Unsupported]
{{end}}

{{if test .CheckList.APILimits.NodeLimits}}
## Node Limits [Unsupported]
{{end}}

{{if test .CheckList.APILimits.TimeLimits}}
## Time Limits [Unsupported]
{{end}}

{{if test .CheckList.APILimits.BatchLimits}}
## Batch Limits [Unsupported]
{{end}}


{{if test .CheckList.DynamicEnvironmentVariables.Used}}
## Dynamic Environment Variables [Unsupported]
{{end}}

{{if test .CheckList.DynamicEnvironmentVariables.ZeroDowntimeSecretResolution}}
### Zero Downtime Secret Resolution [Unsupported]
{{end}}


{{if test .CheckList.Subscriptions.Used}}
## Subscriptions [Unsupported]
{{end}}

{{if test .CheckList.Observability.Used}}
## Obervability [?]

This is not currently able to be detected.
{{end}}


# Executive Summary

The following categories of feature support were detected in your Hasura V2 project:

| Support | Count |
| --- | --- |
{{ range $k, $v := .Summary -}}
| {{$k}} | {{$v}} |
{{end}}

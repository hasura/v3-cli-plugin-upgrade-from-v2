
# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.


## Features

These are the features used by your project:


{{if test .CheckList.Sources.Used -}}
### Sources
{{end}}

{{if test .CheckList.Tables.Used -}}
### Tables
{{end}}

{{if test .CheckList.Relationships.Used -}}
### Relationships
{{end}}

{{if test .CheckList.QueriesAndMutations.Used -}}
### Queries and Mutations
{{end}}

{{if test .CheckList.Actions.Used -}}
### Actions
{{end}}

{{if test .CheckList.RemoteSchemas.Used -}}
### Remote Schemas
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


# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

## Features

These are the features used by your project:

{{if .CheckList.Actions.UsesActions }}
* Actions [Supported]
  {{if .CheckList.Actions.Transforms.UsesTransforms -}}
  * Kriti [Unsupported]
    Kriti is currently unsupported, but you can write your own request/response transformations with the TypeScript connector...
  {{end -}}
{{end -}}

{{if .CheckList.Sources.UsesSources }}
* Sources [Supported]
  {{- if .CheckList.Sources.UsesFromEnv}}
  * FromENV [Supported]
  * TODO: SSL
  * TODO: Max Connections
  * TODO: Timeout
  * TODO: GQL Customisations (DataSource level)
  * TODO: Read Replicas
  * TODO: Dynamic Routing
  * TODO: Latency Test
  * TODO: Reload Datasource
    * TODO: Introspect datasource
    * TODO: Generate the DB schema
  {{- end -}}
{{end}}

### Tables/Models

* Track or Untrack Models(& Bulk)
* Track or Untrack Foreign Key Relationships(& Bulk)
* Track or Untrack Functions (& Bulk)
* Comments on Models (tables)
* Enum Tables
* Apollo federation

### Relationships

* Local Relationship with Foreign key
* Local Relationship without FK
* Remote Relationship 

### Queries & Mutations

* Auto-generated Query Features
* Query List
* Query by PK
* distinct_on
* where
  * Basic Usage
  * With Bool Expression
  * All Operators ( Generic, Text, JSONB, Itree etc.)
  * with aggregate expression
* Pagination
* limit
* offset
* order_by
* aggregate
* Simple Object query
  * Scalar integer & Text
  * Scalar JSON field
  * Nested objects 
  * Aggregate nested object
* Bool expressions
* Introspection
* Run SQL
  * Ability to run DB queries, we solve this by using dbeaver
* Permission
  * Column SELECT, UPDATE, DELETE Permisisons
  * Permission rules

### Actions

* Support Functions (Query)
* Support Procedures (mutations)
* Ability to define types
  * Custom types 
  * Custom scalar types
* HTTP configuration
  * Configure endpoints
  * URL templating (kriti - from context)
  * Request type
  * Headers - Static
  * Headers - Dynamic (Kriti)
  * Forward Client headers on target HTTP request
* Async Actions
* Open API Import
  * Single Import
  * bulk import 
  * error handling during import
  * support yaml & JSON
* Code Gen
  * Support multiple programming languages & frameworks
* Derive Actions
* Action Permissions
  * Basic permissions
* Action Relationships
  * Action to DB
  * Action to RS
  * Action to Action
* Transforms
  * Request Transforms
    * suppport all context objects
    * Request Method
    * Request URL
    * Request Body
  * Response Transforms
    * Support all context objects
    * Debugging mode (include transorm errors)
  * Payload Transform
    * Response Body Transforms
* Async Action logs cleanup (?)

### Remote Schemas

* Configuration
  * URL & from env
  * Timeout
  * Headers
  * Headers Dynamic
* Relationships
  * RS to DB
  * RS to RS
  * Arg Presets
* Permissions
* Bypass Auth

### Event Triggers

* Trigger type
  * Insert
  * Update
  * Delete
  * Via console
  * Column specific updates
* Trigger Protocols
  * HTTP webhook
* Auto Cleanup Event Logs
  * Make sure event logs are cleared periodically
* Request Transform
  * Headers
  * Request Method
  * Request URL
  * Request Body
  * Context 
* Response Transform
  * Transform response data
* Payload Transform
  * Response Body Transforms
* Retry Logic
  * Configuration
* Simulation
  * Event triggers can be tested by manually triggering from console
* Codegen
* Observability
* Performance

### Authentication

* JWT
  * JWT secret
  * JWK
* Webhook
* Admin Role
* Unauthorised Access
* Multiple Admin Secrets
* Multiple JWTs

### Authz

* Role from Session variable
* Role based scheme generation
* Configure Row Permissions
* Column Permissions
* Aggregation permissions
* Row Fetch Limit
* Root Field Visibility
* Column Presets
* Backend Only Mutations
* Permissions Operators
* Permissions summary (on Console)

### Relay

* Query
  * Connection Objects, Nodes, Edges
* Subscription
* Mutation

### RESTified Endpoints

* Basic Graphql
* Arguments from payload body
* Arguments from URL params
* Different HTTP methods
* RESTified endpoint playground (like GraphiQL UI, to test REST endpoints)

### Caching

* Basic Caching with @cached
* With Custom TTL @cached(ttl: 120)
* forcing a refresh @cached(refresh: true)

### Allow LIst

* Configuer Query Collections with Role access
* Multiple Query Colections & Roles
* Caching Metrics
* EE Suppport (suppport custom Redis configurations)

### API Limits

* Rate limits
* Depth limits
* Node limits
* Time limits
* Batch Limits

### Dynamic Env Vars

* Secret resolution with zero downtime

### Other

* Metadata Consistency
* Schema Registry
* Federation
  * Apollo Federation

## Telemetry Detected Features

### Subscriptions		

* Cannot be determined by metadata or introspection 	

### Observability	Open Telemetry	Traces URL

* Metrics URL
* Trace Propogation (b3 or tracecontext)
* Headers
* Attrubutes
* APM Integrations	Prometheus
  * Datadog
  * New Relic
  * Azure Monitor
* Query Tags (?)


## Executive Summary

Your project is partially supported. In order to use V3 you will need to write some custom TS.

The following features that you are using are not currently supported:

* TODO

{{if .V3Directory }}
A copy of this report has been saved to {{.V3Directory}}/v2_upgrade_report.md
{{else}}
To save a copy of this report specify --v3-directory [V3_DIRECTORY]
{{end}}
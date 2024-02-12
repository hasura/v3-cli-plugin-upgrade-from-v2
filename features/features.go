package features

// Taken from https://docs.google.com/spreadsheets/d/1CfTEmDsjnuh_flWhSFVZ6wqWR7m7egeyVD6BXppgCyE/edit#gid=1074788665

type Feature struct {
	Type       string
	Category   string
	Feature    string
	SubFeature string
}

var features = []Feature{
	{"Metadata based", "Datasource", "URL from Environment/ Secrets", ""},
	{"Metadata based", "Datasource", "Advanced Connection Configurations", "SSL, Max Connections, TImeout etc."},
	{"Metadata based", "Datasource", "GQL Customisations (DataSource level)", ""},
	{"Metadata based", "Datasource", "Read Replicas", ""},
	{"Metadata based", "Datasource", "Dynamic Routing", ""},
	{"Metadata based", "Datasource", "Latency Test (From Console)", ""},
	{"Metadata based", "Datasource", "Reload Datasource", "Re introspect datasource and generate the DB schema"},

	{"Metadata based", "Tables/Models", "Track or Untrack Models(& Bulk)", ""},
	{"Metadata based", "Tables/Models", "Track or Untrack Foreign Key Relationships(& Bulk)", ""},
	{"Metadata based", "Tables/Models", "Track or Untrack Functions (& Bulk)", ""},
	{"Metadata based", "Tables/Models", "Comments on Models (tables)", ""},
	{"Metadata based", "Tables/Models", "Enum Tables", ""},
	{"Metadata based", "Tables/Models", "Apollo federation", ""},

	{"Metadata based", "Relationships", "Local Relationship with Foreign key", ""},
	{"Metadata based", "Relationships", "Local Relationship without FK", ""},
	{"Metadata based", "Relationships", "Remote Relationship ", "DB to DB"},
	{"Metadata based", "Relationships", "Remote Relationship", "DB to RS"},
	{"Metadata based", "Relationships", "Remote Relationship", "RS to DB"},
	{"Metadata based", "Relationships", "Remote Relationship", "RS to RS"},

	{"Metadata based", "Queries & Mutations", "Auto-generated Query Features", ""},
	{"Metadata based", "Queries & Mutations", "Query List", ""},
	{"Metadata based", "Queries & Mutations", "Query by PK", ""},
	{"Metadata based", "Queries & Mutations", "distinct_on", ""},
	{"Metadata based", "Queries & Mutations", "where", "Basic Usage"},
	{"Metadata based", "Queries & Mutations", "where", "With Bool Expression"},
	{"Metadata based", "Queries & Mutations", "where", "All Operators -( Generic, Text, JSONB, Itree etc.)"},
	{"Metadata based", "Queries & Mutations", "where", "with aggregate expression"},

	{"Metadata based", "Queries & Mutations", "Pagination", ""},
	{"Metadata based", "Queries & Mutations", "limit", ""},
	{"Metadata based", "Queries & Mutations", "offset", ""},
	{"Metadata based", "Queries & Mutations", "order_by", ""},
	{"Metadata based", "Queries & Mutations", "aggregate", "All Aggregations"},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "scalar integer & Text"},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "scalar JSON field"},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "nested objects "},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "aggregate nested object"},
	{"Metadata based", "Queries & Mutations", "Bool expressions", ""},
	{"Metadata based", "Queries & Mutations", "Introspection", ""},
	{"Metadata based", "Queries & Mutations", "Run SQL", "Ability to run DB queries, we solve this by using dbeaver"},
	{"Metadata based", "Queries & Mutations", "Permission", "Column SELECT, UPDATE, DELETE Permisisons"},
	{"Metadata based", "Queries & Mutations", "Permission", "Permission rules"},

	{"Metadata based", "Permissions", "", ""},
	{"Metadata based", "Actions", "", ""},
	{"Metadata based", "Remote Schemas", "Configuration", "URL & from env"},
	{"Metadata based", "Event Triggers", "Trigger type", "Insert"},

	{"Metadata based", "Authentication", "JWT", "JWT secret"},
	{"Metadata based", "Authentication", "JWT", "JWK"},
	{"Metadata based", "Authentication", "Webhook", ""},
	{"Metadata based", "Authentication", "Admin Role", ""},
	{"Metadata based", "Authentication", "Unauthorised Access", ""},
	{"Metadata based", "Authentication", "Multiple Admin Secrets", ""},
	{"Metadata based", "Authentication", "Multiple JWTs", ""},

	{"Metadata based", "Authz", "Role from Session variable", ""},

	{"Metadata based", "Relay", "Query", "Connection Objects, Nodes, Edges"},
	{"Metadata based", "Relay", "Subscription", ""},
	{"Metadata based", "Relay", "Mutation", ""},

	{"Metadata based", "RESTified Endpoints", "Basic Graphql", ""},
	{"Metadata based", "RESTified Endpoints", "Arguments from payload body", ""},
	{"Metadata based", "RESTified Endpoints", "Arguments from URL params", ""},
	{"Metadata based", "RESTified Endpoints", "Different HTTP methods", ""},
	{"Metadata based", "RESTified Endpoints", "RESTified endpoint playground (like GraphiQL UI, to test REST endpoints)", ""},

	{"Metadata based", "Caching ", "Basic Caching with @cached", ""},
	{"Metadata based", "Caching ", "With Custom TTL @cached(ttl: 120)", ""},
	{"Metadata based", "Caching ", "forcing a refresh @cached(refresh: true)", ""},

	{"Metadata based", "Allow List", "Configuer Query Collections with Role access", ""},
	{"Metadata based", "Allow List", "Multiple Query Colections & Roles", ""},
	{"Metadata based", "Allow List", "Caching Metrics ", ""},
	{"Metadata based", "Allow List", "EE Suppport (suppport custom Redis configurations)", ""},

	{"Metadata based", "API Limits", "Rate limits", ""},
	{"Metadata based", "API Limits", "Depth limits", ""},
	{"Metadata based", "API Limits", "Node limits", ""},
	{"Metadata based", "API Limits", "Time limits", ""},
	{"Metadata based", "API Limits", "Batch Limits", ""},

	{"Metadata based", "Dynamic Env Vars", "Secret resolution with zero downtime", ""},

	{"Metadata based", "Metadata Consistency", "", ""},

	{"Metadata based", "Schema Registry", "", ""},

	{"Metadata based", "Federation", "Apollo Federation", ""},

	{"Behaviour Based", "", "https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/", ""},

	{"Less Priority (Telemetry Based)", "Subscriptions", "Cannot be determined by metadata or introspection ", ""},

	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Traces URL"},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Metrics URL"},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Trace Propogation (b3 or tracecontext)"},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Headers"},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Attrubutes"},

	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Prometheus"},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Datadog"},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "New Relic"},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Azure Monitor"},

	{"Less Priority (Telemetry Based)", "Observability", "Query Tags", ""},
}

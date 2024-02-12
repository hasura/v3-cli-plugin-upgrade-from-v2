package features

// Taken from https://docs.google.com/spreadsheets/d/1CfTEmDsjnuh_flWhSFVZ6wqWR7m7egeyVD6BXppgCyE/edit#gid=1074788665

type Feature struct {
	Type       string
	Category   string
	Feature    string
	SubFeature string
	Supported  bool
	Present    bool
}

var features = []Feature{
	{"Metadata based", "Datasource", "URL from Environment/ Secrets", "", false, false},
	{"Metadata based", "Datasource", "Advanced Connection Configurations", "SSL, Max Connections, TImeout etc.", false, false},
	{"Metadata based", "Datasource", "GQL Customisations (DataSource level)", "", false, false},
	{"Metadata based", "Datasource", "Read Replicas", "", false, false},
	{"Metadata based", "Datasource", "Dynamic Routing", "", false, false},
	{"Metadata based", "Datasource", "Latency Test (From Console)", "", false, false},
	{"Metadata based", "Datasource", "Reload Datasource", "Re introspect datasource and generate the DB schema", false, false},

	{"Metadata based", "Tables/Models", "Track or Untrack Models(& Bulk)", "", false, false},
	{"Metadata based", "Tables/Models", "Track or Untrack Foreign Key Relationships(& Bulk)", "", false, false},
	{"Metadata based", "Tables/Models", "Track or Untrack Functions (& Bulk)", "", false, false},
	{"Metadata based", "Tables/Models", "Comments on Models (tables)", "", false, false},
	{"Metadata based", "Tables/Models", "Enum Tables", "", false, false},
	{"Metadata based", "Tables/Models", "Apollo federation", "", false, false},

	{"Metadata based", "Relationships", "Local Relationship with Foreign key", "", false, false},
	{"Metadata based", "Relationships", "Local Relationship without FK", "", false, false},
	{"Metadata based", "Relationships", "Remote Relationship ", "DB to DB", false, false},
	{"Metadata based", "Relationships", "Remote Relationship", "DB to RS", false, false},
	{"Metadata based", "Relationships", "Remote Relationship", "RS to DB", false, false},
	{"Metadata based", "Relationships", "Remote Relationship", "RS to RS", false, false},

	{"Metadata based", "Queries & Mutations", "Auto-generated Query Features", "", false, false},
	{"Metadata based", "Queries & Mutations", "Query List", "", false, false},
	{"Metadata based", "Queries & Mutations", "Query by PK", "", false, false},
	{"Metadata based", "Queries & Mutations", "distinct_on", "", false, false},
	{"Metadata based", "Queries & Mutations", "where", "Basic Usage", false, false},
	{"Metadata based", "Queries & Mutations", "where", "With Bool Expression", false, false},
	{"Metadata based", "Queries & Mutations", "where", "All Operators -( Generic, Text, JSONB, Itree etc.)", false, false},
	{"Metadata based", "Queries & Mutations", "where", "with aggregate expression", false, false},

	{"Metadata based", "Queries & Mutations", "Pagination", "", false, false},
	{"Metadata based", "Queries & Mutations", "limit", "", false, false},
	{"Metadata based", "Queries & Mutations", "offset", "", false, false},
	{"Metadata based", "Queries & Mutations", "order_by", "", false, false},
	{"Metadata based", "Queries & Mutations", "aggregate", "All Aggregations", false, false},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "scalar integer & Text", false, false},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "scalar JSON field", false, false},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "nested objects ", false, false},
	{"Metadata based", "Queries & Mutations", "Simple Object query", "aggregate nested object", false, false},
	{"Metadata based", "Queries & Mutations", "Bool expressions", "", false, false},
	{"Metadata based", "Queries & Mutations", "Introspection", "", false, false},
	{"Metadata based", "Queries & Mutations", "Run SQL", "Ability to run DB queries, we solve this by using dbeaver", false, false},
	{"Metadata based", "Queries & Mutations", "Permission", "Column SELECT, UPDATE, DELETE Permisisons", false, false},
	{"Metadata based", "Queries & Mutations", "Permission", "Permission rules", false, false},

	{"Metadata based", "Permissions", "", "", false, false},
	{"Metadata based", "Actions", "", "", false, false},
	{"Metadata based", "Remote Schemas", "Configuration", "URL & from env", false, false},
	{"Metadata based", "Event Triggers", "Trigger type", "Insert", false, false},

	{"Metadata based", "Authentication", "JWT", "JWT secret", false, false},
	{"Metadata based", "Authentication", "JWT", "JWK", false, false},
	{"Metadata based", "Authentication", "Webhook", "", false, false},
	{"Metadata based", "Authentication", "Admin Role", "", false, false},
	{"Metadata based", "Authentication", "Unauthorised Access", "", false, false},
	{"Metadata based", "Authentication", "Multiple Admin Secrets", "", false, false},
	{"Metadata based", "Authentication", "Multiple JWTs", "", false, false},

	{"Metadata based", "Authz", "Role from Session variable", "", false, false},

	{"Metadata based", "Relay", "Query", "Connection Objects, Nodes, Edges", false, false},
	{"Metadata based", "Relay", "Subscription", "", false, false},
	{"Metadata based", "Relay", "Mutation", "", false, false},

	{"Metadata based", "RESTified Endpoints", "Basic Graphql", "", false, false},
	{"Metadata based", "RESTified Endpoints", "Arguments from payload body", "", false, false},
	{"Metadata based", "RESTified Endpoints", "Arguments from URL params", "", false, false},
	{"Metadata based", "RESTified Endpoints", "Different HTTP methods", "", false, false},
	{"Metadata based", "RESTified Endpoints", "RESTified endpoint playground (like GraphiQL UI, to test REST endpoints)", "", false, false},

	{"Metadata based", "Caching ", "Basic Caching with @cached", "", false, false},
	{"Metadata based", "Caching ", "With Custom TTL @cached(ttl: 120)", "", false, false},
	{"Metadata based", "Caching ", "forcing a refresh @cached(refresh: true)", "", false, false},

	{"Metadata based", "Allow List", "Configuer Query Collections with Role access", "", false, false},
	{"Metadata based", "Allow List", "Multiple Query Colections & Roles", "", false, false},
	{"Metadata based", "Allow List", "Caching Metrics ", "", false, false},
	{"Metadata based", "Allow List", "EE Suppport (suppport custom Redis configurations)", "", false, false},

	{"Metadata based", "API Limits", "Rate limits", "", false, false},
	{"Metadata based", "API Limits", "Depth limits", "", false, false},
	{"Metadata based", "API Limits", "Node limits", "", false, false},
	{"Metadata based", "API Limits", "Time limits", "", false, false},
	{"Metadata based", "API Limits", "Batch Limits", "", false, false},

	{"Metadata based", "Dynamic Env Vars", "Secret resolution with zero downtime", "", false, false},

	{"Metadata based", "Metadata Consistency", "", "", false, false},

	{"Metadata based", "Schema Registry", "", "", false, false},

	{"Metadata based", "Federation", "Apollo Federation", "", false, false},

	{"Behaviour Based", "", "https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/", "", false, false},

	{"Less Priority (Telemetry Based)", "Subscriptions", "Cannot be determined by metadata or introspection ", "", false, false},

	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Traces URL", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Metrics URL", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Trace Propogation (b3 or tracecontext)", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Headers", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "Open Telemetry", "Attrubutes", false, false},

	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Prometheus", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Datadog", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "New Relic", false, false},
	{"Less Priority (Telemetry Based)", "Observability", "APM Integrations", "Azure Monitor", false, false},

	{"Less Priority (Telemetry Based)", "Observability", "Query Tags", "", false, false},
}

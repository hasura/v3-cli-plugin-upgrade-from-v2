package features

import (
	"reflect"
)

// Taken from https://docs.google.com/spreadsheets/d/1CfTEmDsjnuh_flWhSFVZ6wqWR7m7egeyVD6BXppgCyE/edit#gid=1074788665

// Default global checklist.
// Helper template function for randomized testing.
var List = Checklist{}

func SumCategories(checklist interface{}) map[string]int {
	support := make(map[string]int)

	checklistValue := reflect.ValueOf(checklist)
	checklistType := checklistValue.Type()

	for i := 0; i < checklistType.NumField(); i++ {
		field := checklistType.Field(i)
		fieldValue := checklistValue.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			// Nested struct, recursively process
			// Don't need a supported check, since they already have "Used"
			nestedCounts := SumCategories(fieldValue.Interface())
			for category, count := range nestedCounts {
				support[category] += count
			}
		} else {
			// Check for the "category" tag and increment count if true
			if category, ok := field.Tag.Lookup("supported"); ok && fieldValue.Bool() {
				support[category]++
			} else {
				support["unknown"]++
			}
		}
	}

	return support
}

// The feature checklist structure.
// Fields are marked with yes,no,partial for supported status
// This is used for summarizing v2 feature usage
type Checklist struct {
	// TODO: Check this with Rahul? Do we care about backends for this report?
	Sources struct {
		Used      bool `supported:"yes"`
		FromEnv   bool `supported:"yes"`
		PG        bool `supported:"yes"`
		SQLServer bool `supported:"no"`
		MySQL     bool `supported:"no"`
		Mongo     bool `supported:"no"`
	}

	Tables struct {
		Used bool `supported:"yes"`
		/* // These are actions on MD, not features:
		TrackModels bool
		TrackForeignKeys bool
		TrackFunctions bool
		*/
		CommentsOnModels bool `supported:"yes"`
		EnumTables       bool `supported:"no"` // TODO: Check
		ApolloFederation bool `supported:"no"`
	}

	Relationships struct {
		Used                               bool `supported:"yes"`
		LocalRelationshipWithForeignKey    bool `supported:"yes"`
		LocalRelationshipWithoutForeignKey bool `supported:"yes"` // TODO: Check
		RemoteRelationship                 bool `supported:"yes"`
	}

	RemoteSchemas struct {
		Used          bool `supported:"no"`
		Configuration struct {
			Used           bool `supported:"no"`
			FromEnv        bool `supported:"no"`
			Timeout        bool `supported:"no"`
			Headers        bool `supported:"no"`
			DynamicHeaders bool `supported:"no"`
		}
		Relationships struct {
			Used            bool `supported:"no"`
			ToDatabase      bool `supported:"no"`
			ToRemoteSchema  bool `supported:"no"`
			ArgumentPresets bool `supported:"no"`
		}
		Permissions bool `supported:"no"`
		BypassAuth  bool `supported:"no"`
	}

	Queries struct {
		Used bool
		// AutoGeneratedQueryFeatures bool // N/A?
		QueryList         bool `supported:"yes"` // TODO: These are implicit in V2 right?
		QueryByPrimaryKey bool `supported:"yes"`
		DistinctOn        bool `supported:"yes"`
		// It may be difficult to determine this from metadata...
		Where struct {
			Used                 bool `supported:"yes"`
			BasicUsage           bool `supported:"yes"`
			WithBoolExpression   bool `supported:"yes"`
			Operators            bool `supported:"partial"`
			AggregateExpressions bool `supported:"no"`
		}
		Pagination          bool `supported:"yes"`
		Limit               bool `supported:"yes"`
		Offset              bool `supported:"yes"`
		OrderBy             bool `supported:"yes"`
		Aggregate           bool `supported:"no"`
		SimpleObjectQueries struct {
			Used                   bool `supported:"yes"`
			ScalarIntegerAndText   bool `supported:"yes"`
			ScalarJSON             bool `supported:"no"`
			NestedObjects          bool `supported:"no"` // TODO: Check
			AggregateNestedObjects bool `supported:"no"`
		}
		BoolExpressions bool `supported:"partial"`
		Introspection   bool `supported:"yes"` // TODO: Check
		// RunSQL          bool // N/A?
		Permissions struct {
			Used              bool `supported:"yes"`
			ColumnPermisisons bool `supported:"yes"`
			PermissionRules   bool // TODO: Check?
		}
		QueryCollection bool `supported:"no"`
	}

	Mutations struct {
		Used    bool `supported:"no"`
		Simple  bool `supported:"no"`
		Complex bool `supported:"no"`
	}

	LogicalModels struct {
		Used bool `supported:"yes"`
	}

	NativeQueries struct {
		Used    bool `supported:"yes"`
		Queries struct {
			Used        bool `supported:"yes"`
			Permissions bool `supported:"yes"`
		}
		Mutations struct {
			Used        bool `supported:"yes"`
			Permissions bool `supported:"no"`
		}
	}

	Actions struct {
		Used      bool `supported:"yes"`
		Queries   bool `supported:"yes"`
		Mutations bool `supported:"yes"`
		Types     struct {
			Used              bool `supported:"yes"`
			CustomTypes       bool `supported:"yes"`
			CustomScalarTypes bool `supported:"yes"`
		}
		HttpConfiguration struct {
			Used                 bool `supported:"yes"`
			ConfiguredEndpoints  bool `supported:"yes"`
			URLTemplating        bool `supported:"manual"`
			RequestMethod        bool `supported:"partial"`
			StaticHeaders        bool `supported:"partial"`
			DynamicHeaders       bool `supported:"partial"`
			ForwardClientHeaders bool `supported:"no"` // TODO: Check
		}
		AsyncActions bool `supported:"no"`
		/* /// Not part of configuration
		* Open API Import
			* Single Import
			* bulk import
			* error handling during import
			* support yaml & JSON
		* Code Gen
			* Support multiple programming languages & frameworks
		* Derive Actions
		*/
		BasicPermissions bool `supported:"yes"`
		Relationships    struct {
			Used           bool `supported:"yes"`
			ActionToDB     bool `supported:"yes"`
			ActionToRS     bool `supported:"no"` // No RS is supported anyway
			ActionToAction bool `supported:"yes"`
		}
		Transforms struct {
			Used              bool `supported:"manual"` // Subsumes Kriti Feature.
			RequestTransforms struct {
				Used           bool `supported:"manual"`
				ContextObjects bool `supported:"manual"`
				RequestMethod  bool `supported:"manual"`
				RequestURL     bool `supported:"manual"`
				RequestBody    bool `supported:"manual"`
			}
			ResponseTransforms struct {
				Used           bool `supported:"manual"`
				ContextObjects bool `supported:"manual"`
				DebuggingMode  bool `supported:"manual"`
			}
			PayloadTransforms struct {
				Used                   bool `supported:"manual"`
				ResponseBodyTransforms bool `supported:"manual"`
			}
		}
		// * Async Action logs cleanup (?) // What's this?
	}

	EventTriggers struct {
		Used  bool `supported:"no"`
		Types struct {
			Insert bool `supported:"no"`
			Update bool `supported:"no"`
			Delete bool `supported:"no"`
			// Via console: N/A?
			ColumnSpecificUpdates bool `supported:"no"`
		}
		TriggerProtocols struct {
			Webhook bool `supported:"no"`
		}
		// * Auto Cleanup Event Logs (Make sure event logs are cleared periodically): N/A?
		RequestTransforms struct {
			UsesRequestTransforms bool `supported:"no"`
			Headers               bool `supported:"no"`
			RequestMethod         bool `supported:"no"`
			RequestURL            bool `supported:"no"`
			RequestBody           bool `supported:"no"`
			Context               bool `supported:"no"`
		}
		ResponseTransforms struct {
			Used                   bool `supported:"no"`
			UsesResponseTransforms bool `supported:"no"`
		}
		PayloadTransform struct {
			Used                   bool `supported:"no"`
			ResponseBodyTransforms bool `supported:"no"`
		}
		RetryLogicConfiguration bool `supported:"no"`
		/* // N/A?
		* Simulation
			* Event triggers can be tested by manually triggering from console
		* Codegen
		* Observability
		* Performance
		*/
	}

	Authentication struct {
		Used bool `supported:"yes"`
		JWTs struct {
			Used   bool `supported:"yes"`
			Secret bool `supported:"yes"`
			JWK    bool `supported:"yes"`
		}
		Webhook              bool `supported:"yes"`
		AdminRole            bool `supported:"yes"`
		UnauthorisedAccess   bool // TODO: Check
		MultipleAdminSecrets bool `supported:"no"`
		MultipleJWTs         bool `supported:"no"`
	}

	Authorization struct {
		Used                      bool `supported:"yes"`
		RoleFromSessionVariable   bool `supported:"yes"`
		RoleBasedSchemeGeneration bool // TODO: Check
		ConfigureRowPermissions   bool `supported:"yes"`
		ColumnPermissions         bool `supported:"yes"`
		AggregationPermissions    bool `supported:"no"`
		RowFetchLimit             bool // TODO: Check
		RootFieldVisibility       bool `supported:"yes"` // TODO: Check
		ColumnPresets             bool // TODO: Check
		BackendOnlyMutations      bool `supported:"no"`
		PermissionsOperators      bool `supported:"partial"`
		// PermissionsSummary (on Console): N/A?
	}

	Relay struct {
		Used    bool `supported:"yes"`
		Queries struct {
			Used                        bool `supported:"yes"`
			GlobalID                    bool `supported:"yes"`
			ConnectionObjectsNodesEdges bool // TODO: Check
			ConnectionSpec              bool `supported:"no"`
		}
		Subscriptions bool `supported:"no"`
		Mutations     bool `supported:"no"`
	}

	RESTifiedEndpoints struct {
		Used                     bool `supported:"no"`
		BasicGraphql             bool `supported:"no"`
		ArgumentsFromPayloadBody bool `supported:"no"`
		ArgumentsFromURLParams   bool `supported:"no"`
		DifferentHTTPMethods     bool `supported:"no"`
		// * RESTified endpoint playground (like GraphiQL UI, to test REST endpoints): N/A?
	}

	/* // Not able to be detected from metadata:
	### Caching
	* Basic Caching with @cached
	* With Custom TTL @cached(ttl: 120)
	* forcing a refresh @cached(refresh: true)
	*/

	AllowLists struct {
		Used                                    bool `supported:"no"`
		ConfigureQueryCollectionsWithRoleAccess bool `supported:"no"`
		MultipleQueryColectionsAndRoles         bool `supported:"no"`
		CachingMetrics                          bool `supported:"no"`
		EnterpriseSuppport                      struct {
			Used                      bool `supported:"no"`
			CustomRedisConfigurations bool `supported:"no"`
		}
	}

	APILimits struct {
		Used        bool `supported:"no"`
		RateLimits  bool `supported:"no"`
		DepthLimits bool `supported:"no"`
		NodeLimits  bool `supported:"no"`
		TimeLimits  bool `supported:"no"`
		BatchLimits bool `supported:"no"`
	}

	// How to detect this?
	DynamicEnvironmentVariables struct {
		Used                         bool `supported:"no"`
		ZeroDowntimeSecretResolution bool `supported:"no"`
	}

	/* // N/A?
	### Other
	* Metadata Consistency
	* Schema Registry
	* Federation
		* Apollo Federation
	*/

	// The following may need to be detected in telemetry data:
	// --------------------------------------------------------

	Subscriptions struct {
		Used bool `supported:"no"`
	}

	// TODO: Check on this.
	Observability struct {
		Used             bool
		OpenTelemetry    bool
		MetricsURL       bool
		TracePropogation bool
		Headers          bool
		Attrubutes       bool
		APMIntegrations  struct {
			Prometheus   bool
			Datadog      bool
			NewRelic     bool
			AzureMonitor bool
		}
		QueryTags bool
	}
}

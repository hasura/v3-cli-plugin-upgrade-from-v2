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
		Citus     bool `supported:"no"`
		Cockroach bool `supported:"no"`
		SQLServer bool `supported:"no"`
		BigQuery  bool `supported:"no"`
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
		Used      bool `supported:"yes"` // PROGRESS: DONE
		Queries   bool `supported:"yes"` // PROGRESS: DONE
		Mutations bool `supported:"yes"` // PROGRESS: DONE
		Types     struct {
			Used              bool `supported:"yes"` // PROGRESS: DONE
			CustomTypes       bool `supported:"yes"` // PROGRESS: DONE - Don't all actions have custom types?
			CustomScalarTypes bool `supported:"yes"` // PROGRESS: DONE - How to check this?
		}
		HttpConfiguration struct {
			Used                 bool `supported:"yes"`     // PROGRESS: DONE
			ConfiguredEndpoints  bool `supported:"yes"`     // PROGRESS: DONE - What is this? Don't all actions have endpoints configured?
			URLTemplating        bool `supported:"manual"`  // PROGRESS: DONE - Should we check that templating features are actually being used?
			RequestMethod        bool `supported:"partial"` // PROGRESS: DONE - Does this mean anything other than POST?
			StaticHeaders        bool `supported:"partial"` // PROGRESS: DONE
			DynamicHeaders       bool `supported:"partial"` // PROGRESS: DONE - Check this means value_from_env
			ForwardClientHeaders bool `supported:"no"`
		}
		AsyncActions bool `supported:"no"` // PROGRESS: DONE
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
		BasicPermissions bool `supported:"yes"` // PROGRESS: DONE
		Relationships    struct {
			Used           bool `supported:"yes"` // PROGRESS: DONE - Detected on custom types
			ActionToDB     bool `supported:"yes"` // PROGRESS: DONE - Detected on custom types
			ActionToRS     bool `supported:"no"`  // PROGRESS: DONE - This doesn't seem possible in V2
			ActionToAction bool `supported:"yes"` // PROGRESS: DONE - This doesn't seem possible in V2
		}
		Transforms struct {
			// Subsumes Kriti Feature.
			Used              bool `supported:"manual"` // PROGRESS: DONE
			RequestTransforms struct {
				Used           bool `supported:"manual"` // PROGRESS: DONE
				ContextObjects bool `supported:"manual"` // PROGRESS: Unsure - What is this? Does this refer to context variables?
				RequestMethod  bool `supported:"manual"` // PROGRESS: DONE - But how is this different from HttpConfiguration?
				RequestURL     bool `supported:"manual"` // PROGRESS: DONE
				RequestBody    bool `supported:"manual"` // PROGRESS: DONE
			}
			ResponseTransforms struct {
				Used           bool `supported:"manual"` // PROGRESS: DONE
				ContextObjects bool `supported:"manual"` // PROGRESS: Unsure
				DebuggingMode  bool `supported:"manual"` // PROGRESS: Unsure - What does this mean in a live project?
			}
			PayloadTransforms struct {
				Used                   bool `supported:"manual"` // PROGRESS: Unsure - is this the same as RequestBody?
				ResponseBodyTransforms bool `supported:"manual"` // PROGRESS: DONE - Using response_transform.body ?
			}
		}
		// * Async Action logs cleanup (?) // What's this?
	}

	EventTriggers struct {
		Used  bool `supported:"no"` // PROGRESS: DONE
		Types struct {
			Insert                bool `supported:"no"` // PROGRESS: DONE
			Update                bool `supported:"no"` // PROGRESS: DONE
			Delete                bool `supported:"no"` // PROGRESS: DONE
			Console               bool `supported:"no"` // PROGRESS: DONE
			ColumnSpecificUpdates bool `supported:"no"` // PROGRESS: DONE
		}
		TriggerProtocols struct {
			Webhook bool `supported:"no"` // PROGRESS: Unsure - What is this?
		}
		// * Auto Cleanup Event Logs (Make sure event logs are cleared periodically): N/A?
		RequestTransforms struct {
			Used          bool `supported:"no"` // PROGRESS: DONE
			Headers       bool `supported:"no"` // PROGRESS: Unsure - Doesn't seem like this is part of the transofrm, jus the trigger
			RequestMethod bool `supported:"no"` // PROGRESS: DONE - Does this default to POST?
			RequestURL    bool `supported:"no"` // PROGRESS: DONE
			RequestBody   bool `supported:"no"` // PROGRESS: DONE
			Context       bool `supported:"no"` // PROGRESS: Unsure - What is this? Just the variables used for testing?
		}
		ResponseTransforms      bool `supported:"no"` // PROGRESS: DONE
		PayloadTransform        bool `supported:"no"` // PROGRESS: Unsure - Isn't this the same as the request transform body?
		RetryLogicConfiguration bool `supported:"no"` // PROGRESS: DONE
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

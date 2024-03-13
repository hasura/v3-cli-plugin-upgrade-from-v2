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
		Used      bool `supported:"yes"` // PROGRESS: DONE
		FromEnv   bool `supported:"yes"` // PROGRESS: DONE
		PG        bool `supported:"yes"` // PROGRESS: DONE
		Citus     bool `supported:"no"`  // PROGRESS: DONE
		Cockroach bool `supported:"no"`  // PROGRESS: DONE
		SQLServer bool `supported:"no"`  // PROGRESS: DONE
		BigQuery  bool `supported:"no"`  // PROGRESS: DONE
		MySQL     bool `supported:"no"`  // PROGRESS: Skipped
		Mongo     bool `supported:"no"`  // PROGRESS: Skipped
	}

	Tables struct {
		Used bool `supported:"yes"`
		/* // These are actions on MD, not features:
		TrackModels bool
		TrackForeignKeys bool
		TrackFunctions bool
		*/
		CommentsOnModels bool `supported:"yes"` // PROGRESS: DONE
		EnumTables       bool `supported:"no"`  // PROGRESS: DONE
		ApolloFederation bool `supported:"no"`  // PROGRESS: DONE - This isn't specified on the table level
	}

	Relationships struct {
		Used                               bool `supported:"yes"` // PROGRESS: DONE
		LocalRelationshipWithForeignKey    bool `supported:"yes"` // PROGRESS: DONE
		LocalRelationshipWithoutForeignKey bool `supported:"yes"` // PROGRESS: Unknown - Can this be left to the else block?
		RemoteRelationships                bool `supported:"yes"` // PROGRESS: DONE
	}

	RemoteSchemas struct {
		Used          bool `supported:"no"` // PROGRESS: DONE
		Configuration struct {
			Used           bool `supported:"no"`
			FromEnv        bool `supported:"no"` // PROGRESS: DONE
			Timeout        bool `supported:"no"` // PROGRESS: DONE
			Headers        bool `supported:"no"` // PROGRESS: DONE
			DynamicHeaders bool `supported:"no"` // PROGRESS: DONE
		}
		Relationships struct {
			Used            bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
			ToDatabase      bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
			ToRemoteSchema  bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
			ArgumentPresets bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
		}
		Permissions bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
		BypassAuth  bool `supported:"no"` // PROGRESS: Unsure - Not visible in SuperApp currently
	}

	// This will have to be detected by telemetry, there isn't a metadata feature that matches this
	Queries struct {
		Used bool // PROGRESS: Unsure
		// AutoGeneratedQueryFeatures bool // N/A?
		QueryList         bool `supported:"yes"` // PROGRESS: Unsure // TODO: These are implicit in V2 right?
		QueryByPrimaryKey bool `supported:"yes"` // PROGRESS: Unsure
		DistinctOn        bool `supported:"yes"` // PROGRESS: Unsure
		// It may be difficult to determine this from metadata...
		Where struct {
			Used                 bool `supported:"yes"`     // PROGRESS: Unsure
			BasicUsage           bool `supported:"yes"`     // PROGRESS: Unsure
			WithBoolExpression   bool `supported:"yes"`     // PROGRESS: Unsure
			Operators            bool `supported:"partial"` // PROGRESS: Unsure
			AggregateExpressions bool `supported:"no"`      // PROGRESS: Unsure
		}
		Pagination          bool `supported:"yes"` // PROGRESS: Unsure
		Limit               bool `supported:"yes"` // PROGRESS: Unsure
		Offset              bool `supported:"yes"` // PROGRESS: Unsure
		OrderBy             bool `supported:"yes"` // PROGRESS: Unsure
		Aggregate           bool `supported:"no"`  // PROGRESS: Unsure
		SimpleObjectQueries struct {
			Used                   bool `supported:"yes"` // PROGRESS: Unsure
			ScalarIntegerAndText   bool `supported:"yes"` // PROGRESS: Unsure
			ScalarJSON             bool `supported:"no"`  // PROGRESS: Unsure
			NestedObjects          bool `supported:"no"`  // PROGRESS: Unsure
			AggregateNestedObjects bool `supported:"no"`  // PROGRESS: Unsure
		}
		BoolExpressions bool `supported:"partial"` // PROGRESS: Unsure
		Introspection   bool `supported:"yes"`     // TODO: Check  // PROGRESS: Unsure
		// RunSQL          bool // N/A?
		Permissions struct {
			Used              bool `supported:"yes"` // PROGRESS: Unsure
			ColumnPermisisons bool `supported:"yes"` // PROGRESS: Unsure
			PermissionRules   bool // TODO: Check?  // PROGRESS: Unsure
		}
		QueryCollection bool `supported:"no"` // PROGRESS: Unsure
	}

	// Like Queries, hard to tell if these features are being used without looking at telelmetry
	Mutations struct {
		Used    bool `supported:"no"` // PROGRESS: Unsure
		Simple  bool `supported:"no"` // PROGRESS: Unsure
		Complex bool `supported:"no"` // PROGRESS: Unsure
	}

	LogicalModels struct {
		Used bool `supported:"yes"` // PROGRESS: Unsure
	}

	NativeQueries struct {
		Used    bool `supported:"yes"` // PROGRESS: DONE
		Queries struct {
			Used        bool `supported:"yes"` // PROGRESS: DONE
			Permissions bool `supported:"yes"` // PROGRESS: DONE // These are actually configured against the Logical Model not the NQ.
		}
		Mutations struct {
			Used        bool `supported:"yes"` // PROGRESS: DONE // This doesn't seem to be selectable in permissions section of LM on V2.
			Permissions bool `supported:"no"`  // PROGRESS: DONE
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

	// Can't be detected from metadata, or internal state API...
	Authentication struct {
		Used bool `supported:"yes"` // PROGRESS: DONE - Check that this can't be detected from MD
		JWTs struct {
			Used   bool `supported:"yes"` // PROGRESS: DONE
			Secret bool `supported:"yes"` // PROGRESS: DONE
			JWK    bool `supported:"yes"` // PROGRESS: DONE
		}
		Webhook              bool `supported:"yes"` // PROGRESS: DONE
		AdminRole            bool `supported:"yes"` // PROGRESS: DONE
		UnauthorisedAccess   bool // PROGRESS: DONE
		MultipleAdminSecrets bool `supported:"no"` // PROGRESS: DONE
		MultipleJWTs         bool `supported:"no"` // PROGRESS: DONE
	}

	// TODO: Perform more advanced checks for visibility properties.
	//       Markng progress as done for now
	Authorization struct {
		// Question: If there are any roles or session based filters used anywhere this counts?
		Used                      bool `supported:"yes"` // PROGRESS: DONE
		RoleFromSessionVariable   bool `supported:"yes"` // PROGRESS: Unsure - This seems like it refers to session based permissions, not derived roles
		RoleBasedSchemeGeneration bool // This will require checking if schema visibility restrictions are defined for the role - Could be tricky
		ConfigureRowPermissions   bool `supported:"yes"` // PROGRESS: Omitted
		ColumnPermissions         bool `supported:"yes"` // PROGRESS: Omitted
		AggregationPermissions    bool `supported:"no"`  // PROGRESS: Omitted
		RowFetchLimit             bool // PROGRESS: Omitted // TODO: Check
		RootFieldVisibility       bool `supported:"yes"` // PROGRESS: Omitted // TODO: Check
		ColumnPresets             bool // PROGRESS: Omitted // TODO: Check
		BackendOnlyMutations      bool `supported:"no"`      // PROGRESS: Omitted
		PermissionsOperators      bool `supported:"partial"` // PROGRESS: Omitted
		// PermissionsSummary (on Console): N/A?
	}

	// There's no real way to determine that the user is using Relay features since this isn't saved in the project
	Relay struct {
		Used    bool `supported:"yes"` // PROGRESS: DONE - Verify that this is ok to skip since we can't detect relay
		Queries struct {
			Used                        bool `supported:"yes"` // PROGRESS: DONE
			GlobalID                    bool `supported:"yes"` // PROGRESS: DONE
			ConnectionObjectsNodesEdges bool // PROGRESS: DONE - TODO: Check
			ConnectionSpec              bool `supported:"no"` // PROGRESS: DONE
		}
		Subscriptions bool `supported:"no"` // PROGRESS: DONE
		Mutations     bool `supported:"no"` // PROGRESS: DONE
	}

	RESTifiedEndpoints struct {
		Used                     bool `supported:"no"` // PROGRESS: DONE
		BasicGraphql             bool `supported:"no"` // PROGRESS: DONE - How is this different from General USED?
		ArgumentsFromPayloadBody bool `supported:"no"` // PROGRESS: DONE - No good way to detect this. Check.
		ArgumentsFromURLParams   bool `supported:"no"` // PROGRESS: DONE - No good way to detect this
		DifferentHTTPMethods     bool `supported:"no"` // PROGRESS: DONE
		// * RESTified endpoint playground (like GraphiQL UI, to test REST endpoints): N/A?
	}

	/* // Not able to be detected from metadata:
	### Caching
	* Basic Caching with @cached
	* With Custom TTL @cached(ttl: 120)
	* forcing a refresh @cached(refresh: true)
	*/

	AllowLists struct {
		Used                                    bool `supported:"no"` // PROGRESS: DONE - Should this be separate from the REST feature?
		ConfigureQueryCollectionsWithRoleAccess bool `supported:"no"` // PROGRESS: Unsure - Where is this configured?
		MultipleQueryColectionsAndRoles         bool `supported:"no"` // PROGRESS: Unsure - How to detect this?
		CachingMetrics                          bool `supported:"no"` // PROGRESS: Unsure - Does this just refere to the use of @cached?
		EnterpriseSuppport                      struct {
			Used                      bool `supported:"no"` // PROGRESS: Unsure - How to detect this?
			CustomRedisConfigurations bool `supported:"no"` // PROGRESS: Unsure - How to detect this?
		}
	}

	// TODO: Check if disabled flag is active
	APILimits struct {
		Used        bool `supported:"no"` // PROGRESS: DONE
		RateLimits  bool `supported:"no"` // PROGRESS: DONE
		DepthLimits bool `supported:"no"` // PROGRESS: DONE
		NodeLimits  bool `supported:"no"` // PROGRESS: DONE
		TimeLimits  bool `supported:"no"` // PROGRESS: DONE
		BatchLimits bool `supported:"no"` // PROGRESS: DONE
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

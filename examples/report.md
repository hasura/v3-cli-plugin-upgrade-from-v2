# V2 to V3 Upgrade Report

This report outlines the compatibility of your Hasura V2 project with Hasura V3.

These are the features used by your project:

## Sources

### FromENV [Supported]

### Postgres [Supported]

## Tables

### Comments On Models [Supported]

### Enum Tables [Unsupported]

## Relationships

### Local Relationships with Foreign Keys [Supported]

### Remote Relationships [Supported]

## Actions

### Query Actions [Supported]

### Mutation Actions [Supported]

### Types [Supported]

#### Custom Types [Supported]

*Note: All V2 actions use custom types.*

#### Custom Scalar Types [Supported]

### Http Configuration [Supported]

#### Endpoints [Supported]

#### URL Templating [Manual Configuration Required]

#### Dynamic Headers [Partially Supported]

#### Forward Client Headers [Unsupported]

### Async Actions [Unsupported]

### Basic Permissions [Supported]

### Transforms [Manual Configuration Required]

#### Request Transforms [Manual Configuration Required]

##### Request URL [Manual Configuration Required]

##### Request Body [Manual Configuration Required]

## Event Triggers [Unsupported]

Further breakdown of the feature is not included as there is currently no V3 support for event triggers.

## Authorization

## RESTified Endpoints [Unsupported]

### Basic Graphql [Unsupported]

### Multiple HTTP Methods [Unsupported]

## Allow Lists [Unsupported]

## Caching Metrics [Unsupported]

# Executive Summary

The following categories of feature support were detected in your Hasura V2 project:

| Support | Count |
| --- | --- |
| manual | 5 |
| no | 20 |
| partial | 1 |
| unknown | 123 |
| yes | 18 |



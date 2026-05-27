// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package store

const (
	AND_CONDITION = "$AND"
	OR_CONDITION  = "$OR"
)

var opToSql map[string]string

func init() {
	opToSql = map[string]string{
		"$gt":   ">",
		"$lt":   "<",
		"$gte":  ">=",
		"$lte":  "<=",
		"$eq":   "=",
		"$ne":   "!=",
		"$like": "like",
	}
}

// fieldMapper maps the given field name to the expression to be passed in the sql
type fieldMapper func(string) (string, error)

type queryOptions struct {
	stringifyMappedParams bool
}

func sqliteFieldMapper(field string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func postgresFieldMapper(field string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func postgresSortFieldMapper(field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Protect against sql injection, even though this is the column name rather than value

// Use jsonb semantics for Postgres sorting so numeric fields are ordered numerically.

func postgresIndexFieldMapper(field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Postgres index expressions must be parenthesized.

func jsonFieldMapper(field string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Protect against sql injection, even though this is the column name rather than value

func parseQuery(query map[string]any, mapper fieldMapper) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func parseQueryWithOptions(query map[string]any, mapper fieldMapper, options queryOptions) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Sort the keys, mainly for easily testing the generated query

func parseCondition(field string, value any, mapper fieldMapper, options queryOptions) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Check if the map represents a logical operator or multiple conditions

// Simple equality condition

func parseLogicalOperator(operator string, query []map[string]any, mapper fieldMapper, options queryOptions) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func parseFieldCondition(field string, query map[string]any, mapper fieldMapper, options queryOptions) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Sort the keys, mainly for easily testing the generated query

// Check if the map represents a logical operator or multiple conditions

func parseFieldLogicalOperator(field string, operator string, query []map[string]any, mapper fieldMapper, options queryOptions) (string, []any, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func normalizeMappedValue(field string, value any, mapper fieldMapper, options queryOptions) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func isLogicalOperator(operator string) bool { _ = "STUB: not implemented"; return false }

func getOperator(field string) string { _ = "STUB: not implemented"; return "" }

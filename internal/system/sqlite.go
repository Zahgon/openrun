// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package system

import (
	"database/sql"

	"github.com/openrundev/openrun/internal/types"
)

const (
	DB_CONNECTION_CONFIG = "db_connection"
)

func SQLItePragmas(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

func CheckConnectString(connStr string, invoker string, supportedDBs []DBType) (DBType, string, error) {
	_ = "STUB: not implemented"
	return *new(DBType), "", nil
}

type DBType string

const (
	DB_TYPE_SQLITE   DBType = "sqlite"
	DB_TYPE_POSTGRES DBType = "postgres"
)

var (
	DB_SQLITE_POSTGRES = []DBType{DB_TYPE_SQLITE, DB_TYPE_POSTGRES}
	DB_SQLITE          = []DBType{DB_TYPE_SQLITE}
	DRIVER_MAP         = map[DBType]string{
		DB_TYPE_SQLITE:   "sqlite",
		DB_TYPE_POSTGRES: "pgx",
	}
)

func InitDBConnection(connectString string, invoker string, supportedDBs []DBType) (*sql.DB, DBType, error) {
	_ = "STUB: not implemented"
	return nil, *new(DBType), nil
}

//nolint:staticcheck

// Configure connection pool settings for Postgres
// Maximum number of open connections
// Maximum number of idle connections
// Maximum time a connection can be idle
// Maximum lifetime of a connection

// Test the connection

//nolint:errcheck

func telemetryDBSystem(dbType DBType) string { _ = "STUB: not implemented"; return "" }

func GetConnectString(pluginContext *types.PluginContext) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func PostgresRebind(q string) string { _ = "STUB: not implemented"; return "" }

func RebindQuery(dbType DBType, q string) string { _ = "STUB: not implemented"; return "" }

func MapDataType(dbType DBType, dataType string) string { _ = "STUB: not implemented"; return "" }

func FuncNow(dbType DBType) string { _ = "STUB: not implemented"; return "" }

func InsertIgnorePrefix(dbType DBType) string { _ = "STUB: not implemented"; return "" }

func InsertIgnoreSuffix(dbType DBType) string { _ = "STUB: not implemented"; return "" }

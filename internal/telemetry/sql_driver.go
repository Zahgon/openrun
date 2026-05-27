// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"database/sql/driver"
	"sync"
)

const (
	DBSystemSQLite   = "sqlite"
	DBSystemPostgres = "postgresql"
)

// registeredSQLDrivers maps the wrapped driver name to a sync.Once guarding
// the corresponding sql.Register call. Storing the once (rather than the
// driver) ensures that "name is in the map" is equivalent to "registration
// has run to completion": a concurrent caller cannot observe the name as
// present and race ahead to sql.Open before sql.Register has actually
// completed.
var registeredSQLDrivers sync.Map

func SQLDriverName(driverName, dbSystem, invoker string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func baseDriver(driverName string) driver.Driver {
	_ = "STUB: not implemented"
	return *new(driver.Driver)
}

type sqlDriver struct {
	driver   driver.Driver
	dbSystem string
	invoker  string
}

func (d *sqlDriver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

type sqlConn struct {
	driver.Conn
	dbSystem string
	invoker  string
}

func (c *sqlConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (c *sqlConn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (c *sqlConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (c *sqlConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (c *sqlConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

func (c *sqlConn) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *sqlConn) CheckNamedValue(value *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *sqlConn) ResetSession(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *sqlConn) IsValid() bool { _ = "STUB: not implemented"; return false }

type sqlStmt struct {
	driver.Stmt
	query    string
	dbSystem string
	invoker  string
}

func (s *sqlStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (s *sqlStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (s *sqlStmt) CheckNamedValue(value *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

type sqlTx struct {
	driver.Tx
	dbSystem string
	invoker  string
}

func (t *sqlTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (t *sqlTx) Rollback() error { _ = "STUB: not implemented"; return nil }

func queryOperation(query string) string { _ = "STUB: not implemented"; return "" }

func safeDriverName(value string) string { _ = "STUB: not implemented"; return "" }

// Copyright (c) ClaceIO, LLC
// SPDX-License-Identifier: Apache-2.0

package bindings

import (
	"context"
	"database/sql"

	"github.com/openrundev/openrun/internal/types"
)

// mysqlErrNoSuchTable is the MySQL error number for ER_NO_SUCH_TABLE
// (returned when a table referenced in a statement does not exist).
const mysqlErrNoSuchTable = 1146

// mysqlErrNonExistingGrant is the MySQL error number for
// ER_NONEXISTING_TABLE_GRANT (returned by REVOKE when no matching grant
// exists for the user on the table).
const mysqlErrNonExistingGrant = 1147

// MySQL identifier length limits (8.0):
//   - user name: 32 characters
//   - database name: 64 characters
//
// A bindingId is `bnd_` + 27-char ksuid = 31 chars, which already exceeds
// the user limit when any prefix is added. We strip the `bnd_` prefix from
// the binding id when forming the user name to stay within the 32-char
// budget: `cl_p_` (5) + ksuid (27) = 32.
const (
	mysqlUserPrefixProd = "cl_p_"
	mysqlUserPrefixStg  = "cl_s_"
	mysqlDBPrefixProd   = "cl_db_prd_"
	mysqlDBPrefixStg    = "cl_db_stg_"
	mysqlBindingIDTrim  = "bnd_"
	mysqlDefaultHost    = "%"

	// Privileges granted/revoked for `full:*` and `full:tbl`. Chosen to mirror
	// Postgres' `full` semantics (ALL ON TABLES = SELECT/INSERT/UPDATE/DELETE/
	// TRUNCATE/REFERENCES/TRIGGER, plus the schema-level CREATE family) while
	// deliberately *not* using `ALL PRIVILEGES`: that would also revoke the
	// SHOW VIEW baseline granted in GenerateAccount, breaking the account's
	// ability to connect to its default database after a `full:*` revoke.
	//
	// Intentionally excluded:
	//   - GRANT OPTION (matches Postgres binding policy)
	//   - SHOW VIEW (reserved as the connect-time baseline; see GenerateAccount)
	//   - CREATE/ALTER ROUTINE, EXECUTE, EVENT (stored procs and scheduled
	//     events are out of scope for v1; users can grant them via custom SQL)
	//
	// Table-scoped GRANTs cannot include the database-only privileges
	// (CREATE TEMPORARY TABLES, LOCK TABLES) — MySQL rejects them at the
	// table level. mysqlFullPrivilegesTable is the safe subset for `full:tbl`.
	mysqlFullPrivilegesDB    = "SELECT, INSERT, UPDATE, DELETE, CREATE, ALTER, INDEX, DROP, REFERENCES, TRIGGER, CREATE VIEW, CREATE TEMPORARY TABLES, LOCK TABLES"
	mysqlFullPrivilegesTable = "SELECT, INSERT, UPDATE, DELETE, CREATE, ALTER, INDEX, DROP, REFERENCES, TRIGGER, CREATE VIEW"
)

// isMysqlErrNo reports whether err is a MySQL error with the given numeric code.
func isMysqlErrNo(err error, code uint16) bool { _ = "STUB: not implemented"; return false }

type MysqlServiceBinding struct {
	*types.Logger
	serviceConfig map[string]string
	hostPattern   string  // The @host part used when creating users (default '%')
	adminConn     *sql.DB // Admin connection to the MySQL server, available after InitService
}

func init() {
	RegisterServiceBinding("mysql", NewMysqlServiceBinding)
}

var _ ServiceBinding = (*MysqlServiceBinding)(nil)

func NewMysqlServiceBinding() ServiceBinding {
	_ = "STUB: not implemented"
	return *new(ServiceBinding)
}

func (b *MysqlServiceBinding) InitializeService(ctx context.Context, logger *types.Logger, serviceConfig map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (b *MysqlServiceBinding) CloseService(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type MysqlContextKey string

const MYSQL_TX_STATE_KEY MysqlContextKey = "mysql_sb_tx_state"

// mysqlTxState tracks objects created during a logical "transaction" so they
// can be cleaned up on rollback. MySQL DDL auto-commits, so the binding
// interface's transaction methods cannot give us true atomicity; instead we
// record compensating actions and run them if RollbackTransaction is called
// without a matching CommitTransaction.
type mysqlTxState struct {
	createdUsers     []string // user identities in form `'name'@'host'` (already SQL-quoted)
	createdDatabases []string // database identifiers (already backtick-quoted)
	committed        bool
}

func (b *MysqlServiceBinding) BeginTransaction(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	// MySQL DDL (CREATE USER, GRANT, CREATE DATABASE, ...) implicitly commits
	// any open transaction, so we cannot use sql.Tx for atomicity. Stash a
	// tracker in the context so RollbackTransaction can issue compensating
	// DROP statements for anything we created in this logical transaction.
	return *new(context.Context), nil
}

func (b *MysqlServiceBinding) CommitTransaction(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *MysqlServiceBinding) RollbackTransaction(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Commit already ran — nothing to undo.

// Best-effort cleanup. We use a fresh context.Background() so that
// cancellation of the outer context does not abort the cleanup. Failures
// are logged but not propagated; an operator may need to clean up
// manually if a DROP also fails.

func (b *MysqlServiceBinding) GenerateAccount(ctx context.Context, bindingId, bindingPath string, bindingMetadata types.BindingMetadata, derivedFromMetadata *types.BindingMetadata, isStaging bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Derived binding, reuse the base binding's database

// Base binding: create database, then user, then grant ALL on the database.

// ALL on the new database. A database-level grant in MySQL covers
// every current and future table in that database, so this is also
// how the base user owns objects later created by derived bindings.

// Derived binding: create the user in the base binding's database.
// Application privileges are still assigned only by ApplyGrants.

// Let the derived account select the database so negative permission
// checks fail on the intended table/DDL privilege instead of during
// connection setup. MySQL USAGE is not enough to pass the default
// database access check, so use SHOW VIEW as a minimal database-level
// privilege and keep full:* grants from revoking it.

func (b *MysqlServiceBinding) ApplyGrants(ctx context.Context, account map[string]string, bindingMetadata types.BindingMetadata,
	derivedFromMetadata types.BindingMetadata, reapplyAll bool) ([]types.BindingGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *MysqlServiceBinding) processGrants(ctx context.Context, user, host, database string,
	bindingMetadata types.BindingMetadata, reapplyAll bool) ([]types.BindingGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply all grants, can help when new tables are present which need to be granted to the role

// applyPerms runs GRANT or REVOKE statements for binding grants on the MySQL
// admin connection. operation must be "grant" or "revoke".
//
// Future-table behavior: in MySQL, a database-level grant (`ON db.*`) covers
// every current and future table in that database, so `*`-target grants do
// not need a "default privileges" follow-up like Postgres requires. Only
// table-specific grants need to be deferred when the target table does not
// yet exist.
func (b *MysqlServiceBinding) applyPerms(ctx context.Context, operation string,
	grants []types.BindingGrant, quotedDB, database, userRef string) ([]types.BindingGrant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In Postgres the role-on-schema CREATE privilege also lets the
// role drop/alter objects it owns. MySQL splits these into
// distinct privileges, so to match that semantic we bundle the
// table-lifecycle privileges together for `create:*`.

// Use explicit privileges rather than ALL PRIVILEGES so
// revoking full:* does not remove the baseline SHOW VIEW
// privilege needed for the account to connect to its default
// database. There is no separate "sequence" privilege space in
// MySQL (AUTO_INCREMENT columns are part of the table).

// Use the same explicit list as full:* (minus the DB-only
// privileges MySQL rejects at the table level), so the meaning
// of "full" is consistent regardless of target and avoids the
// SHOW VIEW-revoking pitfall of ALL PRIVILEGES.

// trySoftGrant runs a single GRANT or REVOKE on a specific table. Because
// MySQL DDL auto-commits we cannot use SAVEPOINT to recover from a missing
// table the way the Postgres binding does. Instead we precheck the table's
// existence: if it is missing we report the grant as deferred (returns
// false, nil). For revokes we additionally swallow ER_NONEXISTING_TABLE_GRANT
// so removing a grant from a metadata list does not fail if it was never
// applied (e.g. it was deferred).
func (b *MysqlServiceBinding) trySoftGrant(ctx context.Context, isGrant bool, database, table, stmt string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Revoking something that was never granted is harmless.

// Lost a race with a DROP TABLE between precheck and grant; treat
// as deferred so the caller logs and moves on instead of aborting.

// tableExists reports whether the given database/table is present in
// information_schema. Used to precheck table-level GRANT/REVOKE.
func (b *MysqlServiceBinding) tableExists(ctx context.Context, database, table string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// quoteMysqlIdent quotes an identifier (database, table, column) using
// backticks. Embedded backticks are doubled, per MySQL's identifier rules.
func quoteMysqlIdent(name string) string { _ = "STUB: not implemented"; return "" }

// quoteMysqlString quotes a value as a SQL string literal. Embedded single
// quotes are doubled and backslashes escaped, to defend against both standard
// quoting and the (default) NO_BACKSLASH_ESCAPES-off mode where backslash is
// also an escape character.
func quoteMysqlString(s string) string { _ = "STUB: not implemented"; return "" }

// mysqlUserRef returns a `'user'@'host'` reference suitable for use in
// CREATE USER, GRANT, REVOKE, DROP USER statements.
func mysqlUserRef(user, host string) string { _ = "STUB: not implemented"; return "" }

// mysqlURLToDSN converts a mysql:// URL to a go-sql-driver DSN. If
// overrideDB is non-empty it replaces the URL's path-derived database.
func mysqlURLToDSN(rawURL, overrideDB string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Allow callers to pass a bare host:port or a mysql:// URL; reject
// other schemes outright so misconfiguration fails loudly.

// go-sql-driver expects host:port; URLs like mysql://host/db (no port)
// or mysql://[::1]/db (bracketed IPv6, no port) would otherwise emit
// `tcp(host)` which the driver's parser handles inconsistently across
// versions. Default to :3306 when no port is present. net.SplitHostPort
// is used so IPv6 literals are not misclassified as host:port by a naive
// strings.Contains(":") check.

// Map common URL-style flags to go-sql-driver Config fields where
// possible; otherwise fall through to the generic Params bag so the
// driver still sees them.

// loc is a tz name; deferred to the driver via Params to avoid
// pulling time.LoadLocation here.

// buildMysqlAccountURL constructs a new mysql:// URL using the admin URL's
// host/port and the supplied user, password and default database. The result
// is what we store in BindingMetadata.Account["url"] and hand back to apps.
func buildMysqlAccountURL(adminURL, user, password, database string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *MysqlServiceBinding) RunCommand(ctx context.Context, bindingMetadata types.BindingMetadata, command string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Heuristic: a statement that returns rows is treated as a query; otherwise
// we use Exec so callers can issue DDL/DML and still see rows_affected.

//nolint:errcheck

// The mysql driver returns text-protocol values as []byte; convert
// to string so JSON marshalling produces something useful for
// callers.

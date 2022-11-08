package sql

import (
	"context"
	dbsql "database/sql"
	"log"
)

// A wrapper around database/sql. Provide methods that allow higher-level code to
// perform various action on a SQL database.
type SqlDatabase struct {
	Dialect  *SqlDialect
	ConnPool *dbsql.DB
}

// Begin a transaction, and return a pointer to TransactionManager
func (sqldb *SqlDatabase) StartTransaction(ctx context.Context) (*TransactionManager, error) {
	tx, err := sqldb.ConnPool.BeginTx(ctx, &dbsql.TxOptions{})
	if err != nil {
		return nil, err
	}
	return &TransactionManager{Transaction: tx}, nil
}

// Execute SQL statement that does not return anything.
// Return the number of rows affected as well as any error encountered during the process.
func (sqldb *SqlDatabase) ExecuteSQL(stmt *SqlStmt, ctx context.Context) (int64, error) {
	return executeSQLWithCtx(nil, sqldb.ConnPool, stmt, ctx)
}

// Execute a SQL query that returns data from the database.
// Return key-value pairs of column names and their corresponding value
func (sqldb *SqlDatabase) ExecuteQuery(stmt *SqlStmt, ctx context.Context) ([]map[string]interface{}, error) {
	return executeQueryWithCtx(nil, sqldb.ConnPool, stmt, ctx)
}

// Represent a SQL transaction. Allow users to execute a series of SQL query as a single transaction.
type TransactionManager struct {
	Transaction *dbsql.Tx
}

// Execute SQL statement that does not return anything inside a transaction.
// Return the number of rows affected as well as any error encountered during the process.
// Rollback if encountered errors.
func (txManager *TransactionManager) ExecuteSQL(stmt *SqlStmt, ctx context.Context) (int64, error) {
	result, err := executeSQLWithCtx(txManager.Transaction, nil, stmt, ctx)
	if err != nil {
		txManager.Rollback()
		return 0, err
	}
	return result, nil
}

// Execute a SQL query that returns data from the database.
// Return key-value pairs of column names and their corresponding value
// Rollback if encountered errors.
func (txManager *TransactionManager) ExecuteQuery(stmt *SqlStmt, ctx context.Context) ([]map[string]interface{}, error) {
	result, err := executeQueryWithCtx(txManager.Transaction, nil, stmt, ctx)
	if err != nil {
		txManager.Rollback()
		return nil, err
	}
	return result, nil
}

// Rollback all the changes made to the database.
func (txManager *TransactionManager) Rollback() {
	err := txManager.Transaction.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}

// Commit all the changes made to the database.
func (txManager *TransactionManager) Commit() error {
	err := txManager.Transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func executeSQLWithCtx(tx *dbsql.Tx, conn *dbsql.DB, stmt *SqlStmt, ctx context.Context) (int64, error) {
	var result dbsql.Result
	var err error
	if tx == nil {
		result, err = conn.ExecContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params))
	} else {
		result, err = tx.ExecContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params))
	}
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func executeQueryWithCtx(tx *dbsql.Tx, conn *dbsql.DB, stmt *SqlStmt, ctx context.Context) ([]map[string]interface{}, error) {
	var rows *dbsql.Rows
	var err error
	if tx == nil {
		rows, err = conn.QueryContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params)...)
	} else {
		rows, err = tx.QueryContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params)...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make(map[string]interface{})
		vals := make([]interface{}, len(cols))
		for index, colName := range cols {
			row[colName] = &vals[index]
		}
		if err := rows.Scan(vals...); err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return res, nil
}

func castStringListToAnyList(arr []string) []any {
	params := make([]interface{}, len(arr))
	for i, v := range arr {
		params[i] = v
	}
	return params
}

package sql

import (
	"context"
	dbsql "database/sql"
)

type SqlDatabase struct {
	Dialect  *SqlDialect
	ConnPool *dbsql.DB
}

func (sqldb *SqlDatabase) StartTransaction() error {
	return nil
}

func (sqldb *SqlDatabase) ExecuteSQL(stmt *SqlStmt, ctx context.Context) (int64, error) {
	result, err := sqldb.ConnPool.ExecContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params))
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (sqldb *SqlDatabase) ExecuteQuery(stmt *SqlStmt, ctx context.Context) ([]map[string]interface{}, error) {
	rows, err := sqldb.ConnPool.QueryContext(ctx, stmt.Stmt, castStringListToAnyList(stmt.Params)...)
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

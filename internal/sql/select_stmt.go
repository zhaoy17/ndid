package sql

import (
	"errors"
	"fmt"
	"strings"
)

// Generator for SQL SELECT Statement
type SelectStmt struct {
	Dialect        SqlDialect
	ColumnsToQuery []string
	Tables         []string
	QueryCondition SqlQueryFunction
}

// Generate parameterized SELECT SQL statement with FROM clause and WHERE clause based on
// the dialect provided, the tables to query from, as well as the condition while fethcing
// the data.
func (stmt *SelectStmt) GenerateStmt() (res *SqlStmt, err error) {
	var sb strings.Builder

	sb.WriteString("SELECT ")
	selectCluse, err := generateSelectClause(stmt.ColumnsToQuery)
	if err != nil {
		return &SqlStmt{}, err
	}
	sb.WriteString(selectCluse)

	sb.WriteString("\nFROM ")
	fromClause, err := generateFromClause(stmt.Tables)
	if err != nil {
		return &SqlStmt{}, err
	}
	sb.WriteString(fromClause)

	if stmt.QueryCondition == nil {
		return &SqlStmt{
			Stmt:   sb.String(),
			Params: nil,
		}, nil
	} else {
		// generate WHERE clause if query conditions are specified
		sb.WriteString("\nWHERE ")
		whereClause, params, err := stmt.QueryCondition.ToSQLParameterizedQuery(stmt.Dialect, 1)
		if err != nil {
			return &SqlStmt{}, err
		}
		sb.WriteString(whereClause)
		sb.WriteString(";")
		return &SqlStmt{
			Stmt:   sb.String(),
			Params: params,
		}, nil
	}
}

// Generate SELECT clause and validate each columns selected
func generateSelectClause(columns []string) (string, error) {
	numOfCols := len(columns)
	if numOfCols == 0 {
		return "*", nil
	}
	var sb strings.Builder
	for i, s := range columns {
		if !validateToken(s) {
			return "", fmt.Errorf("column validation failed")
		}
		sb.WriteString(s)
		if i < numOfCols-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String(), nil
}

// Generate FROM clause and validate each table that will be queried from
func generateFromClause(tables []string) (string, error) {
	numOfTables := len(tables)

	var sb strings.Builder
	if numOfTables == 0 {
		return "", errors.New("must select from at least one table")
	}

	for i, s := range tables {
		if !validateToken(s) {
			return "", fmt.Errorf("table validation failed")
		}
		sb.WriteString(s)
		if i < numOfTables-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String(), nil
}

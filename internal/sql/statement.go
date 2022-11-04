package sql

import (
	"errors"
	"fmt"
	"strings"
)

// Contains a list of valid SQL dialect supported by the sql generator
type SqlDialect int64

// SQL dialects currently supported are PostgreSQL, MySQL, SQLite
// and MS SQL Server
const (
	Psql SqlDialect = iota
	MySql
	SqlLite
	SqlServer
)

type TableSchema struct {
	TableName string
	Columns   map[string]SqlDataType
}

// Result returned by the SQL generators ready to be passed in for the SQL driver.
// The struct contains a parameterized SQL statement and parameters.
type SqlStmt struct {
	Stmt   string
	Params []string
}

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

// Generate parameterized CREATE TABLE statement based on the dialect provided, and the data type of each specified columns
type CreateTableStmt struct {
	Dialect     SqlDialect
	TableSchema TableSchema
}

func (stmt *CreateTableStmt) GenerateStmt() (res *SqlStmt, err error) {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	if !validateToken(stmt.TableSchema.TableName) {
		return &SqlStmt{}, fmt.Errorf("column validation failed")
	}
	sb.WriteString(stmt.TableSchema.TableName)
	sb.WriteString(" (\n\t")

	index := 0
	for col, dataType := range stmt.TableSchema.Columns {
		if !validateToken(col) {
			return &SqlStmt{}, fmt.Errorf("column validation failed")
		}
		sb.WriteString(col)
		sb.WriteString(" ")

		sqlDataType, err := dataType.ToSqlDataType(stmt.Dialect)
		if err != nil {
			return &SqlStmt{}, err
		}
		sb.WriteString(sqlDataType)
		if index < len(stmt.TableSchema.Columns)-1 {
			sb.WriteString(",\n\t")
		}
		index += 1
	}
	sb.WriteString("\n);")
	return &SqlStmt{
		Stmt:   sb.String(),
		Params: []string{},
	}, nil
}

// Get placeholders in the parameterized SQL statement for the given dialect.
// Different dialect uses different placeholder. PostgreSQL uses $1...$N; SQLLite
// uses ? and MS SQL Server uses %1...%N.
func getPlaceholderForSqlDialect(dialect SqlDialect, index int) (string, error) {
	switch dialect {
	case Psql:
		if index <= 0 {
			return "", errors.New("index has to be greater than 0")
		}
		return fmt.Sprintf("$%d", index), nil
	case MySql, SqlLite:
		return "?", nil
	case SqlServer:
		if index <= 0 {
			return "", errors.New("index has to be greater than 0")
		}
		return fmt.Sprintf("@p%d", index), nil
	default:
		return "", errors.New("unknown dialect or dialect not supported")
	}
}

// Validate if the given token contains only alphabet or number. Implemented to mitigate
// SQL injection attack in parameterized query.
func validateToken(token string) bool {
	if token == "" {
		return false
	}
	for _, r := range token {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

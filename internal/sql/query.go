package sql

import (
	"fmt"
	"strings"
)

// Represent query that can be converted into a SQL WHERE clause statement
type SqlQueryFunction interface {

	// Generate WHERE parameterized query based on the SQLDialect passed in.
	// index represents the starting index of the query placeholder (only applicable to SQLServer and Postgres dialects)
	ToSQLParameterizedQuery(dialect SqlDialect, index int) (stmt string, params []string, err error)
}

// SQL WHERE clause with a single condition
type SqlSingleQueryFunction struct {
	table     string
	column    string
	valueEqTo string
	operator  string
}

func (function *SqlSingleQueryFunction) ToSQLParameterizedQuery(dialect SqlDialect, index int) (string, []string, error) {
	if function.table != "" {
		if !validateToken(function.table) {
			return "", nil, fmt.Errorf("token validation failed")
		}
	}
	if !validateToken(function.column) {
		return "", nil, fmt.Errorf("token validation failed")
	}
	replacement, err := getPlaceholderForSqlDialect(dialect, index)
	if err != nil {
		return "", nil, err
	}
	if function.table == "" {
		return fmt.Sprintf("%s%s%s", function.column, function.operator, replacement), []string{function.valueEqTo}, nil
	}
	return fmt.Sprintf("%s.%s%s%s", function.table, function.column, function.operator, replacement), []string{function.valueEqTo}, nil
}

// SQL WHERE clause with multiple conditions chained by AND or OR
type SqlCompositeQueryFunction struct {
	operator  string
	functions []SqlQueryFunction
}

func (function *SqlCompositeQueryFunction) ToSQLParameterizedQuery(dialect SqlDialect, index int) (string, []string, error) {
	var stmtSb strings.Builder
	var params []string
	for i, f := range function.functions {
		stmt, paramsForF, err := f.ToSQLParameterizedQuery(dialect, index)
		if err != nil {
			return "", nil, err
		}
		index += cap(paramsForF)

		// check for nested composite query
		_, isCompositeQuery := f.(*SqlCompositeQueryFunction)
		if !isCompositeQuery {
			stmtSb.WriteString(stmt)
		} else {
			stmtSb.WriteString(fmt.Sprintf("(%s)", stmt))
		}
		if i < cap(function.functions)-1 {
			stmtSb.WriteString(fmt.Sprintf(" %s ", function.operator))
		}
		params = append(params, paramsForF...)
	}
	return stmtSb.String(), params, nil
}

type SqlColumnEqualQueryFunction struct {
	lTable  string
	lColumn string
	rTable  string
	rColumn string
}

func (function *SqlColumnEqualQueryFunction) ToSQLParameterizedQuery(dialect SqlDialect, index int) (string, []string, error) {
	if !validateToken(function.lTable) {
		return "", nil, fmt.Errorf("token validation failed")
	}
	if !validateToken(function.lColumn) {
		return "", nil, fmt.Errorf("token validation failed")
	}
	if !validateToken(function.rTable) {
		return "", nil, fmt.Errorf("token validation failed")
	}
	if !validateToken(function.rColumn) {
		return "", nil, fmt.Errorf("token validation failed")
	}
	return fmt.Sprintf("%s.%s=%s.%s", function.lTable, function.lColumn, function.rTable, function.rColumn), []string{}, nil
}

// Represent SQL Equal (=) Operator
func SQLEqual(table string, col string, val string) SqlQueryFunction {
	return &SqlSingleQueryFunction{
		table:     table,
		column:    col,
		valueEqTo: val,
		operator:  "=",
	}
}

// Represent SQL Column Equal (tableA.col = tableB.col) when joining two tables
func SQLColumnEqual(lTable string, lCol string, rTable string, rColumn string) SqlQueryFunction {
	return &SqlColumnEqualQueryFunction{
		lTable:  lTable,
		lColumn: lCol,
		rTable:  rTable,
		rColumn: rColumn,
	}
}

// Represent SQL Less Than (<) Operato
func SQLLessThan(table string, col string, val float64) SqlQueryFunction {
	return &SqlSingleQueryFunction{
		table:     table,
		column:    col,
		valueEqTo: fmt.Sprintf("%.2f", val),
		operator:  "<",
	}
}

// Represent SQL Greater Than (>) Operator
func SQLGreaterThan(table string, col string, val float64) SqlQueryFunction {
	return &SqlSingleQueryFunction{
		table:     table,
		column:    col,
		valueEqTo: fmt.Sprintf("%.2f", val),
		operator:  ">",
	}
}

// Represent REGEX comparison for SQL
func SQLRegex(table string, col string, pattern string) SqlQueryFunction {
	return &SqlSingleQueryFunction{
		table:     table,
		column:    col,
		valueEqTo: pattern,
		operator:  "REGEX",
	}
}

// Represent SUBSTRING comparison for SQL
func SQLSubstring(table string, col string, substr string) SqlQueryFunction {
	return &SqlSingleQueryFunction{
		table:     table,
		column:    col,
		valueEqTo: substr,
		operator:  "LIKE",
	}
}

// Represent SQL AND
func SQLAnd(funcList ...SqlQueryFunction) SqlQueryFunction {
	return &SqlCompositeQueryFunction{
		functions: funcList,
		operator:  "AND",
	}
}

// Represent SQL OR
func SQLOr(funcList ...SqlQueryFunction) SqlQueryFunction {
	return &SqlCompositeQueryFunction{
		functions: funcList,
		operator:  "OR",
	}
}

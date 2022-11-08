package sql

import (
	"errors"
	"fmt"
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

// Result returned by the SQL generators ready to be passed in for the SQL driver.
// The struct contains a parameterized SQL statement and parameters.
type SqlStmt struct {
	Stmt   string
	Params []string
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

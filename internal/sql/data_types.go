package sql

import (
	"fmt"
	"strings"
)

type SqlDataType interface {
	//Convert to the correct data types based on the dialect
	ToSqlDataType(dialect SqlDialect) (string, error)
}

// Represent String SQL data type
// Set Len to 0 or negative value for string with unlimited length
type SqlText struct {
	Len     int
	NotNull bool
}

func (dataType *SqlText) ToSqlDataType(dialect SqlDialect) (string, error) {
	var sb strings.Builder
	switch dialect {
	case Psql, MySql, SqlServer:
		if dataType.Len <= 0 {
			sb.WriteString(fmt.Sprintf("VARCHAR(%d)", dataType.Len))
		} else {
			sb.WriteString("TEXT")
		}
	case SqlLite:
		sb.WriteString("TEXT")
	default:
		return "", fmt.Errorf("dialect not supported")
	}
	appendNotNull(&sb, dataType.NotNull)
	return sb.String(), nil
}

// Represent SQL Integer Type
type SqlInteger struct {
	NotNull bool
}

func (dataType *SqlInteger) ToSqlDataType(dialect SqlDialect) (string, error) {
	var sb strings.Builder
	switch dialect {
	case SqlLite, SqlServer, MySql, Psql:
		sb.WriteString("INTEGER")
	default:
		return "", fmt.Errorf("dialect not supported")
	}

	appendNotNull(&sb, dataType.NotNull)
	return sb.String(), nil
}

// Represent SQL Float Type
type SqlFloat struct {
	NotNull bool
}

func (dataType *SqlFloat) ToSqlDataType(dialect SqlDialect) (string, error) {
	var sb strings.Builder
	switch dialect {
	case SqlLite, SqlServer, MySql, Psql:
		sb.WriteString("FLOAT")
	default:
		return "", fmt.Errorf("dialect not supported")
	}

	appendNotNull(&sb, dataType.NotNull)
	return sb.String(), nil
}

// Represent SQL Datetime Type
type SqlDateTime struct {
	NotNull bool
}

func (dataType *SqlFloat) ToDateTimeType(dialect SqlDialect) (string, error) {
	var sb strings.Builder
	switch dialect {
	case SqlLite:
		sb.WriteString("TEXT")
	case Psql:
		sb.WriteString("TIMESTAMP")
	case SqlServer, MySql:
		sb.WriteString("DATETIME")
	default:
		return "", fmt.Errorf("dialect not supported")
	}
	appendNotNull(&sb, dataType.NotNull)
	return sb.String(), nil
}

// Append NOT NULL to the end if needed
func appendNotNull(sb *strings.Builder, notNull bool) {
	if notNull {
		sb.WriteString(" NOT NULL")
	}
}

package sql

import (
	"fmt"
	"strings"
)

type TableSchema struct {
	TableName string
	Columns   map[string]SqlDataType
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

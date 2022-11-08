package schema

import (
	"context"

	sql "github.com/zhaoy17/ndid/internal/sql"
)

const SCHEMA_TABLE_NAME = "ndischema"

type SQLSchemaRepository struct {
	db sql.SqlDatabase
}

func (schemaRepository *SQLSchemaRepository) InsertSchemas(schemas []*NDISchema) error {
	return nil
}

func (schemaRepository *SQLSchemaRepository) Setup(ctx context.Context) error {
	stmt := sql.CreateTableStmt{
		Dialect: *schemaRepository.db.Dialect,
		TableSchema: sql.TableSchema{
			TableName: SCHEMA_TABLE_NAME,
			Columns: map[string]sql.SqlDataType{
				"table_name":        &sql.SqlText{},
				"schema_name":       &sql.SqlText{},
				"schema_definition": &sql.SqlText{},
			},
		},
	}
	sqlStmt, err := stmt.GenerateStmt()
	if err != nil {
		return err
	}
	_, err = schemaRepository.db.ExecuteSQL(sqlStmt, ctx)
	return err
}

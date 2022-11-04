package main

import (
	"fmt"

	sql "github.com/zhaoy17/ndid/internal/sql"
)

func main() {
	s := &sql.SelectStmt{
		Dialect:        sql.Psql,
		ColumnsToQuery: []string{"id", "name", "author"},
		Tables:         []string{"books", "john"},
		QueryCondition: sql.SQLOr(
			sql.SQLEqual("", "id", "1659"),
			sql.SQLColumnEqual("joe", "id", "books", "id"),
			sql.SQLAnd(
				sql.SQLEqual("book", "idb", "1234"),
				sql.SQLOr(
					sql.SQLLessThan("", "joe", 123.345),
					sql.SQLGreaterThan("", "kate", 145),
				),
				sql.SQLEqual("bookid", "abc", "123456"),
			),
		),
	}
	res, err := s.GenerateStmt()
	//res, err := ct.GenerateStmt()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%s\n", res.Stmt)
		fmt.Printf("%s\n", res.Params)
	}
	/**
	ct := &sql.CreateTableStmt{
		Dialect: sql.Psql,
		TableSchema: sql.TableSchema{
			TableName: "exampletable",
			Columns: map[string]sql.SqlDataType{
				"col1": &sql.SqlInteger{NotNull: true},
				"col2": &sql.SqlFloat{NotNull: false},
			},
		},
	}
	**/
}

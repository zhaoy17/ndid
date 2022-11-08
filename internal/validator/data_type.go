package validator

import sqldb "github.com/zhaoy17/ndid/internal/sql"

type NDIDataType interface {
	// Convert DIDDataType to SQLDataType, allowing appropriate SQL token to be generated
	ToSqlDataType() (sqldb.SqlDataType, error)

	// Validate the string value pass in against the data type
	Validate(string) error
}

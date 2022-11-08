package validator

import sqldb "github.com/zhaoy17/ndid/internal/sql"

// Represent string type for DIDDocument's field
type NDIString struct {
	// the maximum length of the string
	MaxLen int

	// Weather or not the string can be null
	NotNull bool

	// Default value, ignored if NotNull has been set to true
	DefaultValue string

	// Enforce the string to have a certain regex pattern
	MustHaveRegexPattern string
}

func (str *NDIString) ToSqlDataType() (sqldb.SqlDataType, error) {
	return &sqldb.SqlText{Len: 0, NotNull: false}, nil
}

func (str *NDIString) Validate(string) error {
	// TODO: need to actually implement it
	return nil
}

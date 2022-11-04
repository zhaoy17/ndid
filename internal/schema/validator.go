package schema

import (
	"fmt"
	"strconv"

	sqldb "github.com/zhaoy17/ndid/internal/sql"
)

type NDIDataType interface {
	// Convert DIDDataType to SQLDataType, allowing appropriate SQL token to be generated
	ToSqlDataType() (sqldb.SqlDataType, error)

	// Validate the string value pass in against the data type
	Validate(string) error
}

// Represent string type for DIDDocument's field

// DefaultValue
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

type NDIInteger struct {
	Max  int
	Min  int
	Enum map[int]bool
}

func (ineger *NDIInteger) ToSqlDataType() (sqldb.SqlDataType, error) {
	return &sqldb.SqlInteger{NotNull: false}, nil
}

func (integer *NDIInteger) Validate(val string) error {
	num, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("%s is not int", val)
	}
	if integer.Enum != nil || len(integer.Enum) == 0 {
		if num > integer.Max {
			return fmt.Errorf("%d cannot be greater than %d", num, integer.Max)
		}
		if num < integer.Min {
			return fmt.Errorf("%d cannot be less than %d", num, integer.Min)
		}
	} else {
		if _, ok := integer.Enum[num]; !ok {
			return fmt.Errorf("%d is not among the list of possible values specified", num)
		}
	}
	return nil
}

type NDIFloat struct {
	Max float64
	Min float64
}

func (float *NDIFloat) ToSqlDataType() (sqldb.SqlDataType, error) {
	return &sqldb.SqlFloat{NotNull: false}, nil
}

func (float *NDIFloat) Validate(val string) error {
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("%s is not float", val)
	}
	if num > float.Max {
		return fmt.Errorf("%f cannot be greater than %f", num, float.Max)
	}
	if num < float.Min {
		return fmt.Errorf("%f cannot be less than %f", num, float.Min)
	}
	return nil
}

package validator

import (
	"fmt"
	"strconv"

	sqldb "github.com/zhaoy17/ndid/internal/sql"
)

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

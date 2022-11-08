package validator

import (
	"fmt"
	"strconv"

	sqldb "github.com/zhaoy17/ndid/internal/sql"
)

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

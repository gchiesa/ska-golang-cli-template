// Package convert provides type conversion utilities.
package convert

import (
	"errors"
	"fmt"
	"strconv"
)

// ErrConversion is returned when a value cannot be converted to integer.
var ErrConversion = errors.New("conversion error")

// ToInteger converts the given value to an integer.
func ToInteger(v any) (int, error) {
	i, err := strconv.Atoi(fmt.Sprintf("%v", v))
	if err != nil {
		return i, fmt.Errorf("%w: cannot convert value %v (type %T) to integer", ErrConversion, v, v)
	}

	return i, nil
}

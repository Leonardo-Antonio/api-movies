package helpers

import "errors"

var (
	ErrStars     = errors.New("must be in the range of 0 to 5")
	ErrIDInvalid = errors.New("the id is invalid")
)

func ErrNull(null string) error {
	var ErrNull = errors.New("column '" + null + "' cannot be null")
	return ErrNull
}

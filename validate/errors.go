package validate

import "errors"

var (
	ErrInvalid    = errors.New("invalid")
	ErrProhibited = errors.New("prohibited")
)

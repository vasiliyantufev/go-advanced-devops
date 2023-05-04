package errors

import "errors"

var (
	ErrTypeIsMissing  = errors.New("the query parameter type is missing")
	ErrTypeIncorrect  = errors.New("the type incorrect")
	ErrNameIsMissing  = errors.New("the query parameter name is missing")
	ErrValueIsMissing = errors.New("the query parameter value is missing")
	ErrHashSum        = errors.New("hash sum does not match calculated")
)

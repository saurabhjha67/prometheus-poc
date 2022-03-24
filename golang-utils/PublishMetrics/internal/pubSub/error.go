package pubSub

// ErrorType represents defined type for Service errors
type ErrorType int

// Custom error types with their error type IDs
const (
	DuplicateUser ErrorType = iota + 1
	UserNotFound
)

// ApiError represent a Api error response
type ApiError struct {
	eType ErrorType
	err   string
}

func (e *ApiError) Error() string {
	return e.err
}

// Type returns type of the Api error
func (e *ApiError) Type() ErrorType {
	return e.eType
}

// SetType sets type for a Api error
func (e *ApiError) SetType(eType ErrorType) {
	e.eType = eType
}

func (e ErrorType) String() string {
	return toString[e]
}

var toString = map[ErrorType]string{
	DuplicateUser: "DuplicateUser",
	UserNotFound:  "UserNotFound",
}

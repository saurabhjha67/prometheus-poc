package contracts

// ErrorResponse describes common error response for all of the API's endpoints
// swagger:model ErrorResponse
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

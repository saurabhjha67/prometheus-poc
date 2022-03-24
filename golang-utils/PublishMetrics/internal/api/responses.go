package api

import (
	"encoding/json"
	"net/http"

	"com.publish.api/internal/contracts"
)

//writeConflictResponse writes not found response
func writeConflictResponse(w http.ResponseWriter, message string) {
	msg := &contracts.ErrorResponse{
		Status:  http.StatusConflict,
		Message: message,
	}
	writeResponse(w, http.StatusConflict, msg)
}

//writeNotFoundResponse writes not found response
func writeNotFoundResponse(w http.ResponseWriter, message string) {
	msg := &contracts.ErrorResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
	writeResponse(w, http.StatusNotFound, msg)
}

//writeErrorResponse writes an error response
func writeErrorResponse(w http.ResponseWriter, message string) {
	msg := &contracts.ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
	writeResponse(w, http.StatusInternalServerError, msg)
}

//writeBadRequestResponse writes an bad request response
func writeBadRequestResponse(w http.ResponseWriter, message string) {
	msg := &contracts.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	writeResponse(w, http.StatusBadRequest, msg)
}

//writeResponse writes reponse headers, code and body.
func writeResponse(w http.ResponseWriter, code int, output interface{}) {
	//convert the output to json
	response, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Add the response code and response body.
	w.WriteHeader(code)

	if output != nil {
		// Set the content type to json for browsers
		w.Header().Set("Content-Type", "application/json")

		if _, err := w.Write(response); err != nil {
			writeErrorResponse(w, "could not write the correct response")
		}
	}
}

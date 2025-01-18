package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type StatusError struct {
	Code   int
	Err    error
	Caller string
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) StatusCode() int {
	return se.Code
}

func writeErrorResponse(w http.ResponseWriter, se StatusError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(se.Code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":   se.Err.Error(),
		"details": se.Caller,
	})
}

func NewError(err error, code int) StatusError {
	pc, _, line, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)

	return StatusError{
		Code:   code,
		Err:    err,
		Caller: fmt.Sprintf("%s#%d", details.Name(), line),
	}
}

func internalServerError(err error) StatusError {
	return NewError(err, http.StatusInternalServerError)
}

func notFoundError(err error) StatusError {
	return NewError(err, http.StatusNotFound)
}

func conflictError(err error) StatusError {
	return NewError(err, http.StatusConflict)
}

func methodNotAllowedError(err error) StatusError {
	return NewError(err, http.StatusMethodNotAllowed)
}

func badRequestError(err error) StatusError {
	return NewError(err, http.StatusBadRequest)
}

func editConflictError(err error) StatusError {
	return NewError(err, http.StatusConflict)
}

func unauthorizedError(err error) StatusError {
	return NewError(err, http.StatusUnauthorized)
}

func unprocessableEntityError(err error) StatusError {
	return NewError(err, http.StatusUnprocessableEntity)
}

package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorCode int

const (
	ErrInternal ErrorCode = iota
)

func (e ErrorCode) ToError(meta error, status int) Error {
	err := Error{Meta: meta.Error(), Code: fmt.Sprint(e)}
	switch e {
	case ErrInternal:
		err.Status = fmt.Sprint(http.StatusInternalServerError)
		err.Title = http.StatusText(http.StatusInternalServerError)
		err.Meta = ""
	}
	return err
}

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Errors   []Error        `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Errors)
}

type Error struct {
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Meta   string `json:"meta,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v: error %v: %v(%v): %v",
		e.Status, e.Code, e.Title, e.Detail, e.Meta)
}

func Err(w http.ResponseWriter, code ErrorCode, err error, status int) {
	e := code.ToError(err, status)
	errors := []Error{e}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&ErrorResponse{Errors: errors})
}

func Errors(w http.ResponseWriter, status int, errors ...Error) {
	var errs []Error
	for _, e := range errors {
		errs = append(errs, e)
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&ErrorResponse{Errors: errs})
}

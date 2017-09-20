package main

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error          error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	AppCode    int    `json:"code"`
	ErrorText  string `json:"error"`
	StatusText string `json:"status"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func InvalidRequestError(err error) *ErrorResponse {
	return &ErrorResponse{
		Error:          err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorText:      err.Error(),
		StatusText:     http.StatusText(http.StatusBadRequest),
	}
}

func ServerError(err error) *ErrorResponse {
	return &ErrorResponse{
		Error:          err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     http.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

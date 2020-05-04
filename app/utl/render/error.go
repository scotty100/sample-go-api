package render

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Errors interface{} `json:"error"`          // the errors to return
	Output string 		`json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {

	// log outgoing error

	render.Status(r, e.HTTPStatusCode)

	return nil
}

func NotFoundRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
	}
}

func ErrRender(err error, httpStatus int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatus,
	}
}

func ErrRenderWithBody(err error, httpStatus int, body interface{}) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatus,
		Errors:body,
	}
}

type FieldError struct {
	field string `json:"field"`
	message string `json:"message"`
}

type ErrInvalidRequest struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Errors []FieldError `json:"fieldErrors"`          // the errors to return
}

func (e *ErrInvalidRequest) Render(w http.ResponseWriter, r *http.Request) error {

	// log outgoing error
	render.Status(r, e.HTTPStatusCode)
	return nil
}



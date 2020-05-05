package render

import (
	"fmt"
	"github.com/BenefexLtd/onehub-go-base/pkg/errors"
	"github.com/go-chi/render"
	logger2 "github.com/teltech/logger"
	"net/http"
)

type ErrResponse struct {
	HTTPStatusCode int   `json:"-"` // http response status code

	//Error       interface{} `json:"error"` // the errors to return (include when we have an error with a body)
	FieldErrors interface{} `json:"fieldErrors"` // any fieldErrors to return
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {

	render.Status(r, e.HTTPStatusCode)
	return nil
}

// could we pull a default logger from the context instead of creating new struct?
type ErrorRenderer struct {
	logger *logger2.Log
}

func NewErrorRenderer(logger *logger2.Log) *ErrorRenderer {
	return &ErrorRenderer{logger: logger}

}

func (er *ErrorRenderer) ErrRender(err errors.OneHubError) render.Renderer {

	// update to only lof server errors and not client errors
	if err.HttpStatus() >= 500 {
		er.logger.Error(fmt.Sprintf("%+v\n", err))
		if err.Err() != nil {
			er.logger.Error(fmt.Sprintf("%+v\n", err.Err))
		}
	} else {
		er.logger.Debug(fmt.Sprintf("%+v\n", err))
	}

	resp := &ErrResponse{
		HTTPStatusCode: err.HttpStatus(),
	}

	if bre, ok := err.(*errors.BadRequestError); ok {
		resp.FieldErrors=bre.FieldErrors
	}

	return resp
}

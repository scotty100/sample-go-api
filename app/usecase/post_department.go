package usecase

import (
	"errors"
	utlerrors "github.com/BenefexLtd/onehub-go-base/pkg/errors"
	"net/http"
)

type PostDepartment struct {
	Name string `json:"name" binding:"required" bson:"name"`
}

func (p *PostDepartment) Bind(r *http.Request) error {

	// look at converting to an attribute based approach

	if p.Name == "" {
		fieldErrors := make([]utlerrors.FieldError, 1)
		fieldErrors[0] = utlerrors.FieldError{Field: "name", Message: "name is required"}

		return utlerrors.BadRequest(errors.New("invalid request"), fieldErrors).(error)
	}

	return nil
}

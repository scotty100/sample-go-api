package usecase

import (
	"github.com/go-chi/render"
	"net/http"
)

// DepartmentPagedResult
// swagger:model departmentPagedResponse
type DepartmentPagedResult struct {
	Content          []DepartmentInfo `json:"content"`
	TotalElements    int64            `json:"totalElements"`
	NumberOfElements int              `json:"numberOfElements"`
	First            bool             `json:"first"`
	Last             bool             `json:"last"`
	Sort             string           `json:"sort"`
	Page             int64            `json:"page"`
	Size             int64            `json:"size"`
}

func (dpr *DepartmentPagedResult) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	render.Status(r, 200)
	return nil
}

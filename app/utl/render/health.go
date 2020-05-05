package render

import (
	"github.com/go-chi/render"
	"net/http"
)

type HealthResponse struct{

	HTTPStatusCode int  `json:"-"`
	Alive          bool `json:"alive"`
}

func (e *HealthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}


func HealthFailureRender() render.Renderer{
	return &HealthResponse{HTTPStatusCode:500, Alive:false}
}

func HealthSuccessRender() render.Renderer{
	return &HealthResponse{HTTPStatusCode:200, Alive:true}
}
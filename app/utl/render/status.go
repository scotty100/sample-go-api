package render

import (
	"github.com/go-chi/render"
	"net/http"
)

func OK(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, 200)
	render.DefaultResponder(w, r, &v)
}

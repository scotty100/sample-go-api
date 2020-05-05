package http

import (
	"context"
	resp "github.com/BenefexLtd/onehub-go-base/pkg/render"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type HealthHandler struct {
	client *mongo.Client
}

func NewHealthHandler(client *mongo.Client) *HealthHandler {
	return &HealthHandler{ client:client}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := h.client.Ping(ctx, readpref.Primary())
	if err != nil {
		render.Render(w, r, resp.HealthFailureRender())

	} else {
		render.Render(w, r, resp.HealthSuccessRender())
	}

}


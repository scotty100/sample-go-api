package app

import (
	adapterhttp "github.com/BenefexLtd/departments-api-refactor/app/adapter/http"
	data "github.com/BenefexLtd/departments-api-refactor/app/adapter/persistence"
	"github.com/BenefexLtd/departments-api-refactor/app/usecase"
	"github.com/BenefexLtd/onehub-go-base/pkg/config"
	"github.com/BenefexLtd/onehub-go-base/pkg/mongo"
	utlrender "github.com/BenefexLtd/onehub-go-base/pkg/render"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/teltech/logger"
	"net/http"
	"time"
)

// wire up the app
// new name of this package should match to the subfolder in the cmd folder

func Start(config *config.Configuration) error {

	r := chi.NewRouter()

	// get this from secret manager -> load from main.go
	connectionString := "mongodb://localhost:27017/departments"

	logger := logger.New()
	mongo := mongo.New(connectionString, config.DB.Timeout, config.DB.Database, logger)
	departmentHandler := getDepartmentHandler(logger, mongo)
	healthHandler := getHealthHandler(mongo)

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(time.Second * 5))

	r.Get("/health", healthHandler.HealthCheck)
	r.Get("/companies/{companyId}/departments", departmentHandler.GetCompanyDepartments)
	r.Get("/companies/{companyId}/departments/{departmentId}", departmentHandler.GetDepartment)
	r.Post("/companies/{companyId}/departments", departmentHandler.CreateDepartment)

	// add paginnate middlewear



	return http.ListenAndServe(config.Server.Port, r)
}

func getDepartmentHandler(logger *logger.Log, mongo *mongo.Datastore) *adapterhttp.DepartmentHandler{

	departmentRepository := data.DepartmentRepositoryImpl{Store: mongo, Logger: logger}
	useCaseService := usecase.NewDepartmentUseCase(&departmentRepository)
	errRenderer := utlrender.NewErrorRenderer(logger)
	return adapterhttp.NewDepartmentHandler(logger, useCaseService, errRenderer)
}

func getHealthHandler( mongo *mongo.Datastore) *adapterhttp.HealthHandler{

	return adapterhttp.NewHealthHandler(mongo.Session)
}
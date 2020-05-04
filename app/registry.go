package app

import (
	"github.com/BenefexLtd/departments-api-refactor/app/adapter/http"
	data "github.com/BenefexLtd/departments-api-refactor/app/adapter/persistence"
	"github.com/BenefexLtd/departments-api-refactor/app/usecase"
	"github.com/BenefexLtd/departments-api-refactor/app/utl/config"
	"github.com/BenefexLtd/departments-api-refactor/app/utl/mongo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/teltech/logger"
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
	r.Use(middleware.Timeout(time.Second * 30))

	r.Get("/health", healthHandler.HealthCheck)
	r.Get("/companies/{companyId}/departments", departmentHandler.GetCompanyDepartments)

	// add paginnate middlewear



	return nil
}

func getDepartmentHandler(logger *logger.Log, mongo *mongo.Datastore) *http.DepartmentHandler{

	departmentRepository := data.DepartmentRepositoryImpl{Store: mongo, Logger: logger}
	//publisher := messaging.Publisher{Topic: "test"}
	//departmentService := service.DepartmentServiceImpl{Repository:&departmentRepository, Publisher: &publisher}
	useCaseService := usecase.NewDepartmentUseCase(&departmentRepository)
	return http.NewDepartmentHandler(&departmentRepository, useCaseService)
}

func getHealthHandler( mongo *mongo.Datastore) *http.HealthHandler{

	return http.NewHealthHandler(mongo.Session)
}
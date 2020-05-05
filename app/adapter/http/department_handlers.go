package http

import (
	"github.com/BenefexLtd/departments-api-refactor/app/usecase"
	onehuberrors "github.com/BenefexLtd/onehub-go-base/pkg/errors"
	httputl "github.com/BenefexLtd/onehub-go-base/pkg/http"
	resp "github.com/BenefexLtd/onehub-go-base/pkg/render"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/teltech/logger"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	logger         *logger.Log
	useCaseService usecase.DepartmentUseCase
	errRender      *resp.ErrorRenderer
}

func NewDepartmentHandler(logger *logger.Log, departmentUseCase usecase.DepartmentUseCase, errRender *resp.ErrorRenderer) *DepartmentHandler {
	return &DepartmentHandler{
		logger:         logger,
		useCaseService: departmentUseCase,
		errRender:      errRender,
	}
}

// GetCompanyDepartments swagger:route GET /companies/{companyId}/departments departments getCompanyDepartments
//
// Get paged list of departments for a company.
//
// Responses:
// 		200: departmentPagedResponse
//
func (h *DepartmentHandler) GetCompanyDepartments(w http.ResponseWriter, req *http.Request) {

	companyId := chi.URLParam(req, "companyId")
	sort := httputl.DefaultQuery(req, "sort", "_id, asc")
	page, _ := strconv.Atoi(httputl.DefaultQuery(req, "page", "0"))
	size, _ := strconv.Atoi(httputl.DefaultQuery(req, "size", "20"))

	res, err := h.useCaseService.GetPagedDepartmentsForCompany(req.Context(), companyId, sort, int64(page), int64(size))
	if err != nil {
		render.Render(w, req, h.errRender.ErrRender(err))
		return
	}

	resp.OK(w, req, &res)
}

// GetDepartment swagger:route GET /companies/{companyId}/departments/{id} departments getDepartment
//
// Get a department for a company by id.
//
// Responses:
// 		200: department
//
func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, req *http.Request) {

	companyId := chi.URLParam(req, "companyId")
	departmentId := chi.URLParam(req, "departmentId")

	department, err := h.useCaseService.GetDepartment(req.Context(), companyId, departmentId)
	if err != nil {
		render.Render(w, req, h.errRender.ErrRender(err))
		return
	}

	resp.OK(w, req, &department)
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, req *http.Request) {
	companyId := chi.URLParam(req, "companyId")
	postDept := &usecase.PostDepartment{}
	if err := render.Bind(req, postDept); err != nil {
		badReqError := err.(*onehuberrors.BadRequestError)
		render.Render(w, req, h.errRender.ErrRender(badReqError))
		return
	}

	dept, err := h.useCaseService.CreateDepartment(req.Context(), companyId, postDept.Name)
	if err != nil {
		render.Render(w, req, h.errRender.ErrRender(err))
		return
	}

	resp.OK(w, req, dept)
}

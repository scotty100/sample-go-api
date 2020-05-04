package http

import (
	"github.com/BenefexLtd/departments-api-refactor/app/domain/service"
	"github.com/BenefexLtd/departments-api-refactor/app/usecase"
	httputl "github.com/BenefexLtd/departments-api-refactor/app/utl/http"
	resp "github.com/BenefexLtd/departments-api-refactor/app/utl/render"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	repository     service.DepartmentRepository
	useCaseService usecase.DepartmentUseCase
}

func NewDepartmentHandler(repository service.DepartmentRepository, departmentUseCase usecase.DepartmentUseCase) *DepartmentHandler {
	return &DepartmentHandler{
		repository:     repository,
		useCaseService: departmentUseCase,
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
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	res, err := h.useCaseService.GetPagedDepartmentsForCompany(req.Context(), companyId, sort, int64(page), int64(size))
	if err != nil {
		render.Render(w, req, resp.ErrRender(err, 500))
		return
	}

	resp.OK(w, req, &res)
}
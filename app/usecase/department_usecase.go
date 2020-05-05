package usecase

import (
	"context"
	"fmt"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/service"
	onehuberrors "github.com/BenefexLtd/departments-api-refactor/app/utl/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type DepartmentUseCase interface {
	GetPagedDepartmentsForCompany(ctx context.Context, companyId string, sort string, page int64, size int64) (*DepartmentPagedResult, onehuberrors.OneHubError)
	GetDepartment(ctx context.Context, companyId, departmentId string) (*DepartmentDto, onehuberrors.OneHubError)
}

type departmentUseCase struct {
	repo service.DepartmentRepository
}

func NewDepartmentUseCase(repo service.DepartmentRepository) *departmentUseCase {
	return &departmentUseCase{
		repo: repo,
	}
}

func (d *departmentUseCase) GetPagedDepartmentsForCompany(ctx context.Context, companyId string, sort string, page int64, size int64) (*DepartmentPagedResult, onehuberrors.OneHubError) {

	depts, err := d.repo.FindCompanyDepartments(ctx, companyId, sort, page, size)
	if err != nil {
		return nil,  onehuberrors.NewOneHubError(err,   fmt.Sprintf("error getting pages companies for %s", companyId), 500)

	}

	count, cErr := d.repo.CountDepartmentsForCompany(ctx, companyId)
	if cErr != nil {
		return nil, onehuberrors.NewOneHubError(cErr,  fmt.Sprintf("error department count for company for %s", companyId), 500)
	}

	return &DepartmentPagedResult{
		Content:          mapToDocumentInfos(depts),
		TotalElements:    count,
		NumberOfElements: len(depts),
		First:            page == 0,
		Last:             (page+1)*size > count,
		Sort:             sort,
		Page:             page,
		Size:             size}, nil

}

func (d *departmentUseCase) GetDepartment(ctx context.Context, companyId, departmentId string) (*DepartmentDto,  onehuberrors.OneHubError) {

	department, err := d.repo.FindByCompanyIdAndId(ctx, companyId, departmentId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, onehuberrors.NotFound(err)
		} else {
			return nil, onehuberrors.NewOneHubError(err, fmt.Sprintf("Error finding department %s for company %s", departmentId, companyId), 500)
		}
	}

	dto := DepartmentDto(*department)
	return &dto, nil
}

//func (d *departmentUseCase) CreateDepartment(ctx context.Context, postDepartment PostDepartment) (*DepartmentDto,  *onehuberrors.OneHubError) {
//
//}

func mapToDocumentInfos(departments []model.Department) []DepartmentInfo {
	var departmentInfos = make([]DepartmentInfo, 0)
	i := 0
	for range departments {
		departmentInfo := DepartmentInfo(departments[i])
		// do not pull all userIds and get the userCount from a database-scripts aggregate query
		departmentInfo.UserCount = len(departmentInfo.UserIds)
		departmentInfos = append(departmentInfos, departmentInfo)
	}

	return departmentInfos
}

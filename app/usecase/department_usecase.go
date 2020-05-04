package usecase

import (
	"context"
	"fmt"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/service"
)

type DepartmentUseCase interface {

	GetPagedDepartmentsForCompany(ctx context.Context, companyId string, sort string, page int64, size int64) (*DepartmentPagedResult,error)
}

type departmentUseCase struct {
	repo service.DepartmentRepository
}

func NewDepartmentUseCase(repo service.DepartmentRepository ) *departmentUseCase {
	return &departmentUseCase{
		repo:    repo,
	}
}

func (d * departmentUseCase) GetPagedDepartmentsForCompany(ctx context.Context, companyId string, sort string, page int64, size int64) (*DepartmentPagedResult,error) {

	depts, err := d.repo.FindCompanyDepartments(ctx, companyId, sort, page, size)
	if err != nil {
		return nil, fmt.Errorf("error getting pages companies for %s : %v", companyId, err)
	}

	count, cErr := d.repo.CountDepartmentsForCompany(ctx,companyId)
	if cErr != nil {
		return nil, fmt.Errorf("error department count for company for %s : %v", companyId, cErr)
	}

	return &DepartmentPagedResult{
		Content:          mapToDocumentInfos(depts),
		TotalElements:    count,
		NumberOfElements: len(depts),
		First:            page == 0,
		Last:             (page+1)*size > count,
		Sort:             sort,
		Page:             page,
		Size:             size}, err

}

func mapToDocumentInfos( departments []model.Department) []DepartmentInfo {
	var departmentInfos = make([]DepartmentInfo, 0)
	i :=0
	for range departments {
		 departmentInfo := DepartmentInfo(departments[i])
		// do not pull all userIds and get the userCount from a database-scripts aggregate query
		departmentInfo.UserCount = len(departmentInfo.UserIds)
		departmentInfos = append(departmentInfos, departmentInfo)
	}

	return departmentInfos
}


package event

import (
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"time"
)

type DepartmentCreatedEvent struct {
	Id          string    `json:"id"`
	CompanyId   string    `json:"companyId"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
}

func NewDepartmentCreatedEvent(department *model.Department) *DepartmentCreatedEvent {
	return &DepartmentCreatedEvent{
		Id:          department.Id,
		CompanyId:   department.CompanyId,
		Name:        department.Name,
		CreatedDate: department.CreatedDate,
	}
}

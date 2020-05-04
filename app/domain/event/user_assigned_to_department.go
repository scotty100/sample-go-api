package event

import "github.com/BenefexLtd/departments-api-refactor/app/domain/model"

type UserAssignedToDepartmentEvent struct {
	Id         string `json:"id"`
	CompanyId  string `json:"companyId"`
	Department string `json:"department"`
	UserId     string `json:"userId"`
}

func NewUserAssignedToDepartmentEvent(department *model.Department, userId string) *UserAssignedToDepartmentEvent {
	return &UserAssignedToDepartmentEvent{
		Id:         department.Id,
		CompanyId:  department.CompanyId,
		Department: department.Name,
		UserId:     userId,
	}
}
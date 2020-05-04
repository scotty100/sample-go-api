package service

import (
	"context"
	"fmt"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/event"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"github.com/BenefexLtd/departments-api-refactor/app/utl/messaging"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DepartmentServiceImpl struct {
	Repository DepartmentRepository
	Publisher  MessagePublisher
}

func (ds *DepartmentServiceImpl) AddUserToDepartment(ctx context.Context, companyId, departmentName, userId string) error {

	department, err := ds.getDepartmentOrCreate(ctx, companyId, departmentName)
	if err != nil {
		return err
	}

	dbErr := ds.Repository.AddUserToDepartment(ctx, department.Id, userId)
	if dbErr != nil {
		return fmt.Errorf("error adding user to department: %v", err)
	}

	userAssignedToDepartmentEvent := event.NewUserAssignedToDepartmentEvent(department, userId)
	ds.Publisher.Publish(ctx, messaging.OneHubEvent{Content: userAssignedToDepartmentEvent}, "UserAssignedToDepartmentEvent")

	return nil
}

func (ds *DepartmentServiceImpl) RemoveUserFromDepartment(ctx context.Context, companyId, departmentName, userId string) error {
	department, err := ds.Repository.FindDepartmentByName(ctx, companyId, departmentName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		} else {
			return fmt.Errorf("error finding department: %s. %v", departmentName, err)
		}
	}

	if department == nil {
		return nil
	}

	dbErr := ds.Repository.RemoveUserFromDepartment(ctx, department.Id, userId)
	if dbErr != nil {
		return fmt.Errorf("error removing user %s from department %s : %v", userId, department.Id, err)
	}

	if len(department.UserIds) == 1 {
		dbErr = ds.Repository.DeleteDepartment(ctx, department.Id)
		if dbErr != nil {
			return fmt.Errorf("error deleting department %s : %v", department.Id, err)
		}
	}

	return nil
}

func (ds *DepartmentServiceImpl) getDepartmentOrCreate(ctx context.Context, companyId, name string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByName(ctx, companyId, name)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error finding department: %v", err)
	}

	if department == nil {
		departmentToAdd := model.Department{
			Id:          uuid.NewV4().String(),
			CompanyId:   companyId,
			Name:        name,
			CreatedDate: time.Now(),
			UserIds:     make([]string, 0),
		}

		newDepartment, err := ds.Repository.AddDepartment(ctx, &departmentToAdd)
		if err != nil {
			return nil, fmt.Errorf("error creating new department %v", err)
		}

		department = newDepartment

	}

	return department, nil
}

func (ds *DepartmentServiceImpl) GetDepartmentForUserId(ctx context.Context, userId string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByUserId(ctx, userId)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error finding department for user %s : %v", userId, err)
	}

	return department, nil
}

func (ds *DepartmentServiceImpl) GetDepartmentByCompanyAndName(ctx context.Context, companyId, name string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByName(ctx, companyId, name)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("error finding department %s for company %s: %v", name, companyId, err)
	}

	return department, nil
}

func (ds *DepartmentServiceImpl) ReplaceDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	createdDepartment, err := ds.Repository.ReplaceDepartment(ctx, department)
	if err != nil {
		return nil, fmt.Errorf("error replacing department : %v", err)
	}
	return createdDepartment, nil
}

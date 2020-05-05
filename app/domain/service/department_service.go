package service

import (
	"context"
	"fmt"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/event"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"github.com/BenefexLtd/onehub-go-base/pkg/messaging"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DepartmentServiceImpl struct {
	Repository DepartmentRepository
	Publisher  MessagePublisher
}

func (ds *DepartmentServiceImpl) AddUserToDepartment(ctx context.Context, companyId, departmentName, userId string) error {

	department, err := ds.GetDepartmentOrCreate(ctx, companyId, departmentName)
	if err != nil {
		return err
	}

	dbErr := ds.Repository.AddUserToDepartment(ctx, department.Id, userId)
	if dbErr != nil {
		return  errors.Wrap(err, "error adding user to department")
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
			return errors.Wrap(err, fmt.Sprintf("error finding department: %s.", departmentName))
		}
	}

	if department == nil {
		return nil
	}

	dbErr := ds.Repository.RemoveUserFromDepartment(ctx, department.Id, userId)
	if dbErr != nil {
		return errors.Wrap(err, fmt.Sprintf("error removing user %s from department %s", userId, department.Id))
	}

	if len(department.UserIds) == 1 {
		dbErr = ds.Repository.DeleteDepartment(ctx, department.Id)
		if dbErr != nil {
			return errors.Wrap(err, fmt.Sprintf("error deleting department %s", department.Id))
		}
	}

	return nil
}

func (ds *DepartmentServiceImpl) GetDepartmentOrCreate(ctx context.Context, companyId, name string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByName(ctx, companyId, name)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.Wrap(err, fmt.Sprintf("error finding department: %s", name))
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
			return nil, errors.Wrap(err, "error creating new department")
		}

		department = newDepartment

	}

	return department, nil
}

func (ds *DepartmentServiceImpl) GetDepartmentForUserId(ctx context.Context, userId string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByUserId(ctx, userId)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.Wrap(err, fmt.Sprintf("error finding department for user %s", userId))
	}

	return department, nil
}

func (ds *DepartmentServiceImpl) GetDepartmentByCompanyAndName(ctx context.Context, companyId, name string) (*model.Department, error) {
	department, err := ds.Repository.FindDepartmentByName(ctx, companyId, name)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.Wrap(err, fmt.Sprintf("error finding department %s for company %s", name, companyId))
	}

	return department, nil
}

func (ds *DepartmentServiceImpl) ReplaceDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	createdDepartment, err := ds.Repository.ReplaceDepartment(ctx, department)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error replacing department"))
	}
	return createdDepartment, nil
}

package data

import (
	"context"
	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
	"github.com/BenefexLtd/departments-api-refactor/app/utl"
	"github.com/teltech/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	utlmongo "github.com/BenefexLtd/departments-api-refactor/app/utl/mongo"

)

const DepartmentsCollection = "departments"

type DepartmentRepositoryImpl struct {
	Store  *utlmongo.Datastore
	Logger *logger.Log
}

func (r *DepartmentRepositoryImpl) FindByCompanyIdAndId(ctx context.Context, companyId string, id string) (*model.Department, error) {

	res := r.Store.Db.Collection(DepartmentsCollection).FindOne(ctx, bson.M{"companyId": companyId, "_id": id})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var department model.Department
	res.Decode(&department)
	department.UserCount = len(department.UserIds)

	return &department, nil
}

func (r *DepartmentRepositoryImpl) FindCompanyDepartments(ctx context.Context, companyId string, sort string, page int64, size int64) ([]model.Department, error) {
	options := options.Find()
	ss := strings.Split(sort, ",")
	dir := 1
	if (len(ss)) == 2 {
		dir = utl.IfThenElse(strings.ToLower(ss[1]) == "desc", -1, 1).(int)
	}
	options.SetSort(bson.D{{translateToMongo(ss[0]), dir}})
	options.SetLimit(int64(size))
	options.SetSkip(int64(page * size))

	cur, err := r.Store.Db.Collection(DepartmentsCollection).Find(ctx, bson.M{"companyId": companyId}, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var departments = mapToDocuments(ctx, *r.Logger, cur)

	return departments, nil
}

func (r *DepartmentRepositoryImpl) CountDepartmentsForCompany(ctx context.Context, companyId string) (int64, error) {
	return r.Store.Db.Collection(DepartmentsCollection).CountDocuments(ctx, bson.M{"companyId": companyId})
}

func (r *DepartmentRepositoryImpl) FindDepartmentByName(ctx context.Context, companyId, name string) (*model.Department, error) {

	res := r.Store.Db.Collection(DepartmentsCollection).FindOne(ctx, bson.M{"companyId": companyId, "name": name})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var department model.Department
	res.Decode(&department)

	return &department, nil
}

func (r *DepartmentRepositoryImpl) FindDepartmentByUserId(ctx context.Context, userId string) (*model.Department, error) {

	res := r.Store.Db.Collection(DepartmentsCollection).FindOne(ctx, bson.M{"userIds": userId})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var department model.Department
	res.Decode(&department)
	department.UserCount = len(department.UserIds)

	return &department, nil
}

func (r *DepartmentRepositoryImpl) AddDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	insertResult, err := r.Store.Db.Collection(DepartmentsCollection).InsertOne(ctx, department)
	if err != nil {
		return nil, err
	}
	department.Id = insertResult.InsertedID.(string)
	return department, nil
}

func (r *DepartmentRepositoryImpl) ReplaceDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	_, err := r.Store.Db.Collection(DepartmentsCollection).ReplaceOne(ctx, bson.M{"_id": department.Id}, department)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (r *DepartmentRepositoryImpl) AddUserToDepartment(ctx context.Context, id string, userId string) error {
	_, err := r.Store.Db.Collection(DepartmentsCollection).UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$addToSet", bson.M{"userIds": userId}},
		},
	)
	return err
}

func (r *DepartmentRepositoryImpl) RemoveUserFromDepartment(ctx context.Context, id string, userId string) error {
	_, err := r.Store.Db.Collection(DepartmentsCollection).UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$pull", bson.M{"userIds": userId}},
		},
	)
	return err
}

func (r *DepartmentRepositoryImpl) DeleteDepartment(ctx context.Context, id string) error {
	_, err := r.Store.Db.Collection(DepartmentsCollection).DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
	return err
}

func (r *DepartmentRepositoryImpl) FindDepartmentsForCompanyWithoutHeadUsersSet(ctx context.Context, companyId string) ([]model.Department, error) {
	cur, err := r.Store.Db.Collection(DepartmentsCollection).Find(ctx, bson.M{"companyId": companyId, "headUsersUpdatedDate": new(time.Time)})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var departments = mapToDocuments(ctx, *r.Logger, cur)

	return departments, nil
}

func mapToDocuments(ctx context.Context, logger logger.Log, cur *mongo.Cursor) []model.Department {
	var departments = make([]model.Department, 0)
	for cur.Next(ctx) {
		var department model.Department
		err := cur.Decode(&department)
		if err != nil {
			logger.Errorf("Department decode error: %s", err.Error())
			continue
		}
		// do not pull all userIds and get the userCount from a database-scripts aggregate query
		department.UserCount = len(department.UserIds)
		departments = append(departments, department)
	}

	return departments
}

//func mapToDocumentInfos(ctx context.Context, logger logger.Log, cur *mongo.Cursor) []model.DepartmentInfo {
//	var departments = make([]model.DepartmentInfo, 0)
//	for cur.Next(ctx) {
//		var department model.DepartmentInfo
//		err := cur.Decode(&department)
//		if err != nil {
//			logger.Errorf("Department decode error: %s", err.Error())
//			continue
//		}
//		// do not pull all userIds and get the userCount from a database-scripts aggregate query
//		department.UserCount = len(department.UserIds)
//		departments = append(departments, department)
//	}
//
//	return departments
//}

func translateToMongo(objectField string) string {
	if objectField == "id" {
		return "_id"
	}

	return objectField
}

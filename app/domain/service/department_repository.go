package service


import
(
	"context"

	"github.com/BenefexLtd/departments-api-refactor/app/domain/model"
)

type DepartmentRepository interface {
	FindByCompanyIdAndId(ctx context.Context, companyId string, id string) (*model.Department, error)
	FindCompanyDepartments(ctx context.Context, companyId string, sort string, page int64, size int64) ([]model.Department, error)
	CountDepartmentsForCompany(ctx context.Context, companyId string) (int64, error)
	FindDepartmentByName(ctx context.Context, companyId string, name string) (*model.Department, error)
	FindDepartmentByUserId(ctx context.Context, userId string) (*model.Department, error)
	FindDepartmentsForCompanyWithoutHeadUsersSet(ctx context.Context, companyId string) ([]model.Department, error)
	AddDepartment(ctx context.Context, department *model.Department) (*model.Department, error)
	ReplaceDepartment(ctx context.Context, department *model.Department) (*model.Department, error)
	AddUserToDepartment(ctx context.Context, id string, userId string) error
	RemoveUserFromDepartment(ctx context.Context, id string, userId string) error
	DeleteDepartment(ctx context.Context, id string) error
}

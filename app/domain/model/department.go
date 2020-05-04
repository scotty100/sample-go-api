package model

import "time"

// A Department
// swagger:model department
type Department struct {
	Id                  string    `json:"id" bson:"_id"`
	CompanyId           string    `json:"companyId" bson:"companyId"`
	Name                string    `json:"name" bson:"name"`
	CreatedDate         time.Time `json:"createdDate" bson:"createdDate"`
	UserIds             []string  `json:"userIds" bson:"userIds"`
	HeadUserIds         []string  `json:"headUserIds" bson:"headUserIds"`
	HeadUserUpdatedDate time.Time `json:"-" bson:"headUsersUpdatedDate"`
	UserCount           int       `json:"userCount" bson:",omitempty"`
}

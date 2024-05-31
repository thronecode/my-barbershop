package admin

import (
	"backend/infra/admin"
	"backend/utils"
)

// IAdmin is the interface that defines the methods that the repository should implement
type IAdmin interface {
	List(params *utils.RequestParams) (*admin.PagAdmin, error)
	Get(id *int) (*admin.Admin, error)
	GetByUsername(username *string) (*admin.Admin, error)
	Add(adm *admin.Admin) (*int, error)
	Update(id *int, password *string) error
	Delete(id *int) error
}

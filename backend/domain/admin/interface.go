package admin

import (
	"backend/utils"
)

// IAdmin is the interface that defines the methods that the repository should implement
type IAdmin interface {
	List(params *utils.RequestParams) (*Pag, error)
	Get(id *int) (*Admin, error)
	GetByUsername(username *string) (*Admin, error)
	Add(adm *Admin) (*int, error)
	Update(id *int, password *string) error
	Delete(id *int) error
}

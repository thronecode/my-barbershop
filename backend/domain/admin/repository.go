package admin

import (
	"backend/config/database"
	"backend/infra/admin"
	"backend/infra/admin/postgres"
	"backend/utils"
)

type repository struct {
	pg *postgres.PGAdmin
}

// New creates a new instance of the repository
func New(tx *database.DBTransaction) IAdmin {
	return &repository{pg: &postgres.PGAdmin{DB: tx}}
}

// List returns a paginated list of admins
func (r *repository) List(params *utils.RequestParams) (*admin.Pag, error) {
	return r.pg.List(params)
}

// Get returns an admin by its id
func (r *repository) Get(id *int) (*admin.Admin, error) {
	return r.pg.Get(id)
}

// GetByUsername returns an admin by its username
func (r *repository) GetByUsername(username *string) (*admin.Admin, error) {
	return r.pg.GetByUsername(username)
}

// Add adds a new admin
func (r *repository) Add(adm *admin.Admin) (*int, error) {
	return r.pg.Add(adm)
}

// Update updates the password of an admin
func (r *repository) Update(id *int, password *string) error {
	return r.pg.Update(id, password)
}

// Delete deletes an admin
func (r *repository) Delete(id *int) error {
	return r.pg.Delete(id)
}

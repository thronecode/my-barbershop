package admin

// Admin is the struct that represents the admin user
type Admin struct {
	ID       *int    `converter:"id" sql:"adm.id"`
	Username *string `converter:"username" sql:"adm.username"`
	Password *string `converter:"password" sql:"adm.password"`
}

// FilterListAdmins is the struct that represents the filter for the list of admins
type FilterListAdmins struct {
	Username *string `converter:"username"`
}

// PagAdmin is the struct that represents the paginated list of admins
type PagAdmin struct {
	Data  []Admin
	Next  *bool
	Count *int
}

package admin

// Admin is the struct that represents the admin user
type Admin struct {
	ID       *int    `converter:"id"`
	Username *string `converter:"username"`
	Password *string `converter:"password"`
}

// FilterListAdmins is the struct that represents the filter for the list of admins
type FilterListAdmins struct {
	Username *string `converter:"username"`
}

// PagAdmin is the struct that represents the paginated list of admins
type PagAdmin struct {
	Data  []Admin
	Next  *bool
	Count *int64
}

package admin

import "time"

// Admin is the struct that represents the admin user
type Admin struct {
	ID        *int       `converter:"id" sql:"adm.id"`
	Username  *string    `converter:"username" sql:"adm.username"`
	Password  *string    `converter:"password" sql:"adm.password"`
	DeletedAt *time.Time `converter:"deleted_at" sql:"adm.deleted_at"`
}

// FilterList is the struct that represents the filter for the list of admins
type FilterList struct {
	Username *string `converter:"username"`
}

// Pag is the struct that represents the paginated list of admins
type Pag struct {
	Data  []Admin
	Next  *bool `converter:"next"`
	Count *int  `converter:"count"`
}

package admin

import "time"

// Admin is the struct that represents the admin user
type Admin struct {
	ID        *int       `converter:"id"`
	Username  *string    `converter:"username"`
	Password  *string    `converter:"password"`
	DeletedAt *time.Time `converter:"deleted_at"`
}

// Pag is the struct that represents the paginated list of admins
type Pag struct {
	Data  []Admin
	Next  *bool
	Count *int
}

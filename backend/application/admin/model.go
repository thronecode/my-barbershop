package admin

// AdminInput is the struct that represents the input for the admin
type AdminInput struct {
	Username *string `json:"username" converter:"username" binding:"required"`
	Password *string `json:"password" converter:"password" binding:"required"`
}

// AdminUpdateInput is the struct that represents the input for the admin update
type AdminUpdateInput struct {
	Password *string `json:"password" converter:"password" binding:"required"`
}

// AdminOutput is the struct that represents the output for the admin
type AdminOutput struct {
	ID       *int    `json:"id" converter:"id"`
	Username *string `json:"username" converter:"username"`
}

// AdminPagOutput is the struct that represents the paginated list of admins
type AdminPagOutput struct {
	Data  []AdminOutput `json:"data,omitempty" converter:"data"`
	Next  *bool         `json:"next,omitempty" converter:"next"`
	Count *int64        `json:"count,omitempty" converter:"count"`
}

package admin

// Input is the struct that represents the input for the admin
type Input struct {
	Username *string `json:"username" converter:"username" binding:"required"`
	Password *string `json:"password" converter:"password" binding:"required"`
}

// UpdateInput is the struct that represents the input for the admin update
type UpdateInput struct {
	Password *string `json:"password" converter:"password" binding:"required"`
}

// Output is the struct that represents the output for the admin
type Output struct {
	ID       *int    `json:"id" converter:"id"`
	Username *string `json:"username" converter:"username"`
}

// PagOutput is the struct that represents the paginated list of admins
type PagOutput struct {
	Data  []Output `json:"data,omitempty" converter:"data"`
	Next  *bool    `json:"next,omitempty" converter:"next"`
	Count *int64   `json:"count,omitempty" converter:"count"`
}

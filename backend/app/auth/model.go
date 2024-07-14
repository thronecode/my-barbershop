package auth

// LoginInput is the struct that represents the input for the login
type LoginInput struct {
	Username *string `json:"username" converter:"username" binding:"required"`
	Password *string `json:"password" converter:"password" binding:"required"`
}

// LoginOutput is the struct that represents the output for the login
type LoginOutput struct {
	Token *string `json:"token"`
}

// Admin is the struct that represents the admin user
type Admin struct {
	ID       *int    `converter:"id"`
	Username *string `converter:"username"`
	Password *string `converter:"password"`
}

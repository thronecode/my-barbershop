package service

// Service is a struct that represents the service
type Service struct {
	ID             *int     `converter:"id"`
	Name           *string  `converter:"name"`
	Description    *string  `converter:"description"`
	Duration       *int     `converter:"duration"`
	Price          *float64 `converter:"price"`
	CommissionRate *float64 `converter:"commission_rate"`
	IsCombo        *bool    `converter:"is_combo"`
	Kinds          []string `converter:"kinds"`
	DeletedAt      *string  `converter:"deleted_at"`
}

// Pag is a struct that represents the paginated list of services
type Pag struct {
	Data  []Service
	Next  *bool `converter:"next"`
	Count *int  `converter:"count"`
}

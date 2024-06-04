package service

// Output represents the output structure for the service entity
type Output struct {
	ID             *int     `json:"id" converter:"id"`
	Name           *string  `json:"name" converter:"name"`
	Description    *string  `json:"description" converter:"description"`
	Duration       *int     `json:"duration" converter:"duration"`
	Price          *float64 `json:"price" converter:"price"`
	CommissionRate *float64 `json:"commission_rate" converter:"commission_rate"`
	IsCombo        *bool    `json:"is_combo" converter:"is_combo"`
	Kinds          []string `json:"kinds" converter:"kinds"`
}

// PagOutput represents the paginated output structure for the service entity
type PagOutput struct {
	Data  []Output `json:"data"`
	Next  *bool    `json:"next" converter:"next"`
	Count *int     `json:"count" converter:"count"`
}

// Input represents the input structure for creating/updating a service
type Input struct {
	Name           *string  `json:"name" validate:"required" converter:"name"`
	Description    *string  `json:"description" converter:"description"`
	Duration       *int     `json:"duration" validate:"required" converter:"duration" validate:"required"`
	Price          *float64 `json:"price" validate:"required" converter:"price" validate:"required"`
	CommissionRate *float64 `json:"commission_rate" converter:"commission_rate"`
	IsCombo        *bool    `json:"is_combo" converter:"is_combo"`
	Kinds          []string `json:"kinds" converter:"kinds"`
}

package service

import "github.com/jackc/pgx/v5/pgtype"

// Service is a struct that represents the service
type Service struct {
	ID             *int                  `converter:"id" sql:"ser.id"`
	Name           *string               `converter:"name" sql:"ser.name"`
	Description    *string               `converter:"description" sql:"ser.description"`
	Duration       *int                  `converter:"duration" sql:"ser.duration"`
	Price          *float64              `converter:"price" sql:"ser.price"`
	CommissionRate *float64              `converter:"commission_rate" sql:"ser.commission_rate"`
	IsCombo        *bool                 `converter:"is_combo" sql:"ser.is_combo"`
	Kinds          *pgtype.Array[string] `converter:"kinds" sql:"ser.kinds"`
	DeletedAt      *string               `converter:"deleted_at" sql:"ser.deleted_at"`
}

// FilterList is a struct that represents the filter for the list of services
type FilterList struct {
	Name    *string  `converter:"name"`
	Kinds   []string `converter:"kinds"`
	IsCombo *bool    `converter:"is_combo"`
}

// Pag is a struct that represents the paginated list of services
type Pag struct {
	Data  []Service
	Next  *bool
	Count *int
}

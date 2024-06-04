package barber

import "time"

// Barber is the struct that represents the barber
type Barber struct {
	ID             *int       `converter:"id"`
	Name           *string    `converter:"name"`
	PhotoURL       *string    `converter:"photo_url" `
	CommissionRate *float64   `converter:"commission_rate"`
	DeletedAt      *time.Time `converter:"deleted_at"`
}

// Pag is the struct that represents the paginated list of barbers
type Pag struct {
	Data  []Barber
	Next  *bool
	Count *int
}

// Checkin is the struct that represents the checkin of a barber
type Checkin struct {
	ID       *int       `converter:"id"`
	BarberID *int       `converter:"barber_id"`
	DateTime *time.Time `converter:"date_time"`
}

// PagCheckin is the struct that represents the paginated list of checkins
type PagCheckin struct {
	Data  []Checkin
	Next  *bool
	Count *int
}

package barber

import "time"

// Barber is the struct that represents the barber user
type Barber struct {
	ID             *int       `converter:"id" sql:"bar.id"`
	Name           *string    `converter:"name" sql:"bar.name"`
	PhotoURL       *string    `converter:"photo_url" sql:"bar.photo_url"`
	CommissionRate *float64   `converter:"commission_rate" sql:"bar.commission_rate"`
	DeletedAt      *time.Time `converter:"deleted_at" sql:"bar.deleted_at"`
}

// FilterList is the struct that represents the filter for the list of barbers
type FilterList struct {
	Name *string `converter:"name"`
}

// Pag is the struct that represents the paginated list of barbers
type Pag struct {
	Data  []Barber
	Next  *bool
	Count *int
}

// Checkin is the struct that represents the checkin of a barber
type Checkin struct {
	ID       *int       `converter:"id" sql:"bch.id"`
	BarberID *int       `converter:"barber_id" sql:"bch.barber_id"`
	DateTime *time.Time `converter:"date_time" sql:"bch.date_time"`
}

// FilterCheckinList is the struct that represents the filter for the list of checkins
type FilterCheckinList struct {
	InitialDate *time.Time `converter:"initial_date"`
	FinalDate   *time.Time `converter:"final_date"`
}

// PagCheckin is the struct that represents the paginated list of checkins
type PagCheckin struct {
	Data  []Checkin
	Next  *bool
	Count *int
}

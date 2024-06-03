package barber

// Input is the struct that represents the input for the barber
type Input struct {
	Name           *string  `json:"name" converter:"name" binding:"required"`
	PhotoURL       *string  `json:"photo_url" converter:"photo_url" binding:"required"`
	CommissionRate *float64 `json:"commission_rate" converter:"commission_rate"`
}

// Output is the struct that represents the output for the barber
type Output struct {
	ID             *int     `json:"id" converter:"id"`
	Name           *string  `json:"name" converter:"name"`
	PhotoURL       *string  `json:"photo_url" converter:"photo_url"`
	CommissionRate *float64 `json:"commission_rate" converter:"commission_rate"`
}

// PagOutput is the struct that represents the paginated list of barbers
type PagOutput struct {
	Data  []Output `json:"data,omitempty" converter:"data"`
	Next  *bool    `json:"next,omitempty" converter:"next"`
	Count *int64   `json:"count,omitempty" converter:"count"`
}

// CheckinInput is the struct that represents the input for the checkin
type CheckinInput struct {
	BarberID *int    `converter:"barber_id"`
	DateTime *string `json:"date_time" converter:"date_time" binding:"required"`
}

// CheckinOutput is the struct that represents the output for the checkin
type CheckinOutput struct {
	ID       *int    `json:"id" converter:"id"`
	BarberID *int    `json:"barber_id" converter:"barber_id"`
	DateTime *string `json:"date_time" converter:"date_time"`
}

// PagCheckinOutput is the struct that represents the paginated list of checkins
type PagCheckinOutput struct {
	Data  []CheckinOutput `json:"data,omitempty" converter:"data"`
	Next  *bool           `json:"next,omitempty" converter:"next"`
	Count *int64          `json:"count,omitempty" converter:"count"`
}

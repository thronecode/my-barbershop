package barber

import (
	"backend/utils"
)

type IBarber interface {
	List(params *utils.RequestParams) (*Pag, error)
	Get(id *int) (*Barber, error)
	Add(bar *Barber) (*int, error)
	Update(id *int, bar *Barber) error
	Delete(id *int) error
	AddCheckin(checkin *Checkin) (*int, error)
	GetCheckins(barberID *int, params *utils.RequestParams) (*PagCheckin, error)
}

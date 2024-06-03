package barber

import (
	"backend/infra/barber"
	"backend/utils"
)

type IBarber interface {
	List(params *utils.RequestParams) (*barber.Pag, error)
	Get(id *int) (*barber.Barber, error)
	Add(bar *barber.Barber) (*int, error)
	Update(id *int, bar *barber.Barber) error
	Delete(id *int) error
	AddCheckin(checkin *barber.Checkin) (*int, error)
	GetCheckins(barberID *int, params *utils.RequestParams) (*barber.PagCheckin, error)
}

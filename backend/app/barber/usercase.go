package barber

import (
	"backend/config/database"
	"backend/domain/barber"
	"backend/domain/service"
	"backend/sorry"
	"backend/utils"
	"strconv"
)

// List is the function that lists all barbers
func List(params *utils.RequestParams) (*PagOutput, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	barbers, err := barberRepo.List(params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	res := new(PagOutput)
	if err = utils.ConvertStruct(barbers, res); err != nil {
		return nil, sorry.Err(err)
	}

	res.Data = make([]Output, len(barbers.Data))
	for i := range barbers.Data {
		if err = utils.ConvertStruct(&barbers.Data[i], &res.Data[i]); err != nil {
			return nil, sorry.Err(err)
		}
	}

	return res, nil
}

// Get is the function that gets a barber by its ID
func Get(id *int) (*Output, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar, err := barberRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if bar == nil {
		return nil, sorry.NewErr("barber not found")
	}

	res := new(Output)
	if err = utils.ConvertStruct(bar, res); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Add is the function that adds a barber
func Add(input *Input) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar := new(barber.Barber)
	if err = utils.ConvertStruct(input, bar); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(bar, res); err != nil {
		return nil, sorry.Err(err)
	}

	if res.ID, err = barberRepo.Add(bar); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Update is the function that updates a barber
func Update(id *int, input *Input) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar, err := barberRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if bar == nil {
		return nil, sorry.NewErr("barber not found")
	}

	data := new(barber.Barber)
	if err = utils.ConvertStruct(input, data); err != nil {
		return nil, sorry.Err(err)
	}

	if err = barberRepo.Update(id, bar); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(data, res); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Delete is the function that deletes a barber
func Delete(id *int) error {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	if err = barberRepo.Delete(id); err != nil {
		return sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return sorry.Err(err)
	}

	return nil
}

// AddCheckin is the function that adds a check-in for a barber
func AddCheckin(input *CheckinInput) (*CheckinOutput, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar, err := barberRepo.Get(input.BarberID)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if bar == nil {
		return nil, sorry.NewErr("barber not found")
	}

	checkin := new(barber.Checkin)
	if err = utils.ConvertStruct(input, checkin); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(CheckinOutput)
	if err = utils.ConvertStruct(checkin, res); err != nil {
		return nil, sorry.Err(err)
	}

	if res.ID, err = barberRepo.AddCheckin(checkin); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// GetCheckins is the function that gets check-ins for a barber
func GetCheckins(barberID *int, params *utils.RequestParams) (*PagCheckinOutput, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar, err := barberRepo.Get(barberID)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if bar == nil {
		return nil, sorry.NewErr("barber not found")
	}

	checkins, err := barberRepo.GetCheckins(barberID, params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	res := new(PagCheckinOutput)
	if err = utils.ConvertStruct(checkins, res); err != nil {
		return nil, sorry.Err(err)
	}

	res.Data = make([]CheckinOutput, len(checkins.Data))
	for i := range checkins.Data {
		if err = utils.ConvertStruct(&checkins.Data[i], &res.Data[i]); err != nil {
			return nil, sorry.Err(err)
		}
	}

	return res, nil
}

// AddService is the function that adds a service for a barber
func AddService(barberID *int, services []int) error {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)
	serviceRepo := service.New(tx)

	bar, err := barberRepo.Get(barberID)
	if err != nil {
		return sorry.Err(err)
	}

	if bar == nil {
		return sorry.NewErr("barber not found")
	}

	for i := range services {
		serv, err := serviceRepo.Get(&services[i])
		if err != nil {
			return sorry.Err(err)
		}

		if serv == nil {
			return sorry.NewErr("service with id:" + strconv.Itoa(services[i]) + "  not found")
		}
	}

	if err = barberRepo.AddService(barberID, services); err != nil {
		return sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return sorry.Err(err)
	}

	return nil
}

// DeleteService is the function that deletes a service for a barber
func DeleteService(barberID, serviceiD *int) error {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return sorry.Err(err)
	}
	defer tx.Rollback()

	barberRepo := barber.New(tx)

	bar, err := barberRepo.Get(barberID)
	if err != nil {
		return sorry.Err(err)
	}

	if bar == nil {
		return sorry.NewErr("barber not found")
	}

	if err = barberRepo.DeleteService(barberID, serviceiD); err != nil {
		return sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return sorry.Err(err)
	}

	return nil
}

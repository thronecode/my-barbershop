package service

import (
	"backend/config/database"
	"backend/domain/service"
	"backend/sorry"
	"backend/utils"
)

// List is the function that lists all services
func List(params *utils.RequestParams) (*PagOutput, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	serviceRepo := service.New(tx)

	services, err := serviceRepo.List(params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	res := new(PagOutput)
	if err = utils.ConvertStruct(services, res); err != nil {
		return nil, sorry.Err(err)
	}

	res.Data = make([]Output, len(services.Data))
	for i := range services.Data {
		if err = utils.ConvertStruct(&services.Data[i], &res.Data[i]); err != nil {
			return nil, sorry.Err(err)
		}
	}

	return res, nil
}

// Get is the function that gets a service by its ID
func Get(id *int) (*Output, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	serviceRepo := service.New(tx)

	ser, err := serviceRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if ser == nil {
		return nil, sorry.NewErr("service not found")
	}

	res := new(Output)
	if err = utils.ConvertStruct(ser, res); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Add is the function that adds a service
func Add(input *Input) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	serviceRepo := service.New(tx)

	exists, err := serviceRepo.GetByName(input.Name)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if exists != nil {
		return nil, sorry.NewErr("service already exists")
	}

	ser := new(service.Service)
	if err = utils.ConvertStruct(input, ser); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(ser, res); err != nil {
		return nil, sorry.Err(err)
	}

	if res.ID, err = serviceRepo.Add(ser); err != nil {
		return nil, sorry.Err(err)
	}

	if err = serviceRepo.AddPriceHistory(res.ID, res.Price); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Update is the function that updates a service
func Update(id *int, input *Input) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	serviceRepo := service.New(tx)

	ser, err := serviceRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if ser == nil {
		return nil, sorry.NewErr("service not found")
	}

	if ser.Price != nil && input.Price != nil && *ser.Price != *input.Price {
		if err = serviceRepo.AddPriceHistory(id, input.Price); err != nil {
			return nil, sorry.Err(err)
		}
	}

	if err = utils.ConvertStruct(input, ser); err != nil {
		return nil, sorry.Err(err)
	}

	if err = serviceRepo.Update(id, ser); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(ser, res); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Delete is the function that deletes a service
func Delete(id *int) error {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return sorry.Err(err)
	}
	defer tx.Rollback()

	serviceRepo := service.New(tx)

	if err = serviceRepo.Delete(id); err != nil {
		return sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return sorry.Err(err)
	}

	return nil
}

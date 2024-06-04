package admin

import (
	"backend/config/database"
	"backend/domain/admin"
	"backend/sorry"
	"backend/utils"
)

// List is the function that lists all admins
func List(params *utils.RequestParams) (*PagOutput, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	admins, err := adminRepo.List(params)
	if err != nil {
		return nil, sorry.Err(err)
	}

	res := new(PagOutput)
	if err = utils.ConvertStruct(admins, res); err != nil {
		return nil, sorry.Err(err)
	}

	res.Data = make([]Output, len(admins.Data))
	for i := range admins.Data {
		if err = utils.ConvertStruct(&admins.Data[i], &res.Data[i]); err != nil {
			return nil, sorry.Err(err)
		}
	}

	return res, nil
}

// Get is the function that gets an admin by its ID
func Get(id *int) (*Output, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	adm, err := adminRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if adm == nil {
		return nil, sorry.NewErr("admin not found")
	}

	res := new(Output)
	if err = utils.ConvertStruct(adm, res); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Add is the function that adds an admin
func Add(input *Input) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	exists, err := adminRepo.GetByUsername(input.Username)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if exists != nil {
		return nil, sorry.NewErr("admin already exists")
	}

	adm := new(admin.Admin)
	if err = utils.ConvertStruct(input, adm); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(adm, res); err != nil {
		return nil, sorry.Err(err)
	}

	if *adm.Password, err = utils.HashPassword(*input.Password); err != nil {
		return nil, sorry.Err(err)
	}

	if res.ID, err = adminRepo.Add(adm); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Update is the function that updates an admin
func Update(id *int, input *UpdateInput) (*Output, error) {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	adm, err := adminRepo.Get(id)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if adm == nil {
		return nil, sorry.NewErr("admin not found")
	}

	hashedPassword, err := utils.HashPassword(*input.Password)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if err = adminRepo.Update(id, &hashedPassword); err != nil {
		return nil, sorry.Err(err)
	}

	res := new(Output)
	if err = utils.ConvertStruct(adm, res); err != nil {
		return nil, sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, sorry.Err(err)
	}

	return res, nil
}

// Delete is the function that deletes an admin
func Delete(id *int) error {
	tx, err := database.NewTransaction(false)
	if err != nil {
		return sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	if err = adminRepo.Delete(id); err != nil {
		return sorry.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return sorry.Err(err)
	}

	return nil
}

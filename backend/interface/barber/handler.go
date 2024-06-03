package barber

import (
	"backend/app/barber"
	"backend/sorry"
	"backend/utils"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// add godoc
// @Summary Add barber
// @Description Add a new barber
// @Tags barber
// @Accept  json
// @Produce  json
// @Param input body barber.Input true "Barber input"
// @Success 201 {object} barber.Output
// @Router /barber [post]
func add(c *gin.Context) {
	var (
		input  = new(barber.Input)
		output *barber.Output
		err    error
	)

	if err = c.ShouldBindJSON(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.Add(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// list godoc
// @Summary List barbers
// @Description List all barbers
// @Tags barber
// @Accept  json
// @Produce  json
// @Param name query string false "full or partial barber name"
// @Success 200 {object} barber.PagOutput
// @Router /barber [get]
func list(c *gin.Context) {
	var (
		params utils.RequestParams
		output *barber.PagOutput
		err    error
	)

	if params, err = utils.ParseParams(c); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.List(&params); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// get godoc
// @Summary Get barber
// @Description Get barber by ID
// @Tags barber
// @Accept  json
// @Produce  json
// @Param id path int true "Barber ID"
// @Success 200 {object} barber.Output
// @Router /barber/{id} [get]
func get(c *gin.Context) {
	var (
		id     int
		output *barber.Output
		err    error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.Get(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// update godoc
// @Summary Update barber
// @Description Update barber by ID
// @Tags barber
// @Accept  json
// @Produce  json
// @Param id path int true "Barber ID"
// @Param input body barber.Input true "Barber update input"
// @Success 200 {object} barber.Output
// @Router /barber/{id} [put]
func update(c *gin.Context) {
	var (
		id     int
		input  = new(barber.Input)
		output *barber.Output
		err    error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if err = c.ShouldBindJSON(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.Update(&id, input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// remove godoc
// @Summary Delete barber
// @Description Delete barber by ID
// @Tags barber
// @Accept  json
// @Produce  json
// @Param id path int true "Barber ID"
// @Success 204
// @Router /barber/{id} [delete]
func remove(c *gin.Context) {
	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if err = barber.Delete(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// addCheckin godoc
// @Summary Add check-in
// @Description Add a check-in for a barber
// @Tags barber
// @Accept  json
// @Produce  json
// @Param id path int true "Barber ID"
// @Param input body barber.CheckinInput true "Check-in input"
// @Success 201 {object} barber.CheckinOutput
// @Router /barber/{id}/checkin [post]
func addCheckin(c *gin.Context) {
	var (
		input  = new(barber.CheckinInput)
		output *barber.CheckinOutput
		err    error
	)

	if err = c.ShouldBindJSON(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.AddCheckin(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// getCheckins godoc
// @Summary Get check-ins
// @Description Get check-ins for a barber
// @Tags barber
// @Accept  json
// @Produce  json
// @Param id path int true "Barber ID"
// @Param initial_date query string false "Initial date"
// @Param final_date query string false "Final date"
// @Success 200 {object} barber.PagCheckinOutput
// @Router /barber/{id}/checkin [get]
func getCheckins(c *gin.Context) {
	var (
		id     int
		params utils.RequestParams
		output *barber.PagCheckinOutput
		err    error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if params, err = utils.ParseParams(c); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = barber.GetCheckins(&id, &params); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

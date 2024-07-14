package service

import (
	"backend/app/service"
	"backend/sorry"
	"backend/utils"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// add godoc
// @Summary Add service
// @Description Add a new service
// @Tags service
// @Accept  json
// @Produce  json
// @Param input body service.Input true "Service input"
// @Success 201 {object} service.Output
// @Router /service [post]
func add(c *gin.Context) {
	var (
		input  = new(service.Input)
		output *service.Output
		err    error
	)

	if err = c.ShouldBindJSON(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = service.Add(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// list godoc
// @Summary List services
// @Description List all services
// @Tags service
// @Accept  json
// @Produce  json
// @Param name query string false "full or partial service name"
// @Param barber_id query int false "barber ID"
// @Param is_combo query bool false "is combo"
// @Param kinds query []string false "kinds"
// @Success 200 {object} service.PagOutput
// @Router /service [get]
func list(c *gin.Context) {
	var (
		params utils.RequestParams
		output *service.PagOutput
		err    error
	)

	if params, err = utils.ParseParams(c); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = service.List(&params); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// get godoc
// @Summary Get service
// @Description Get service by ID
// @Tags service
// @Accept  json
// @Produce  json
// @Param id path int true "Service ID"
// @Success 200 {object} service.Output
// @Router /service/{id} [get]
func get(c *gin.Context) {
	var (
		id     int
		output *service.Output
		err    error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = service.Get(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// update godoc
// @Summary Update service
// @Description Update service by ID
// @Tags service
// @Accept  json
// @Produce  json
// @Param id path int true "Service ID"
// @Param input body service.Input true "Service update input"
// @Success 200 {object} service.Output
// @Router /service/{id} [put]
func update(c *gin.Context) {
	var (
		id     int
		input  = new(service.Input)
		output *service.Output
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

	if output, err = service.Update(&id, input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// remove godoc
// @Summary Delete service
// @Description Delete service by ID
// @Tags service
// @Accept  json
// @Produce  json
// @Param id path int true "Service ID"
// @Success 204
// @Router /service/{id} [delete]
func remove(c *gin.Context) {
	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if err = service.Delete(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

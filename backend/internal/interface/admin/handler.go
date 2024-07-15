package admin

import (
	"github.com/thronecode/my-barbershop/backend/internal/app/admin"
	"github.com/thronecode/my-barbershop/backend/internal/sorry"
	"github.com/thronecode/my-barbershop/backend/internal/utils"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// add godoc
// @Summary Add admin
// @Description Add a new admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param input body admin.Input true "Admin input"
// @Success 201 {object} admin.Output
// @Router /admin [post]
func add(c *gin.Context) {
	var (
		input  = new(admin.Input)
		output *admin.Output
		err    error
	)

	if err = c.ShouldBindJSON(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = admin.Add(input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

// list godoc
// @Summary List admins
// @Description List all admins
// @Tags admin
// @Accept  json
// @Produce  json
// @Param username query string false "full or partial username"
// @Success 200 {object} admin.PagOutput
// @Router /admin [get]
func list(c *gin.Context) {
	var (
		params utils.RequestParams
		output *admin.PagOutput
		err    error
	)

	if params, err = utils.ParseParams(c); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = admin.List(&params); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// get godoc
// @Summary Get admin
// @Description Get admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success 200 {object} admin.Output
// @Router /admin/{id} [get]
func get(c *gin.Context) {
	var (
		id     int
		output *admin.Output
		err    error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if output, err = admin.Get(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// update godoc
// @Summary Update admin
// @Description Update admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Param input body admin.UpdateInput true "Admin update input"
// @Success 200 {object} admin.Output
// @Router /admin/{id} [put]
func update(c *gin.Context) {
	var (
		id     int
		input  = new(admin.UpdateInput)
		output *admin.Output
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

	if output, err = admin.Update(&id, input); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

// remove godoc
// @Summary Delete admin
// @Description Delete admin by ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param id path int true "Admin ID"
// @Success 204
// @Router /admin/{id} [delete]
func remove(c *gin.Context) {
	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		sorry.Handling(c, err)
		return
	}

	if err = admin.Delete(&id); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

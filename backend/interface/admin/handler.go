package admin

import (
	"backend/application/admin"
	"backend/sorry"
	"backend/utils"

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
// @Param input body admin.AdminInput true "Admin input"
// @Success 201 {object} admin.AdminOutput
// @Router /admin [post]
func add(c *gin.Context) {
	var (
		input  = new(admin.AdminInput)
		output *admin.AdminOutput
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
// @Success 200 {object} admin.AdminPagOutput
// @Router /admin [get]
func list(c *gin.Context) {
	var (
		params utils.RequestParams
		output *admin.AdminPagOutput
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
// @Success 200 {object} admin.AdminOutput
// @Router /admin/{id} [get]
func get(c *gin.Context) {
	var (
		id     int
		output *admin.AdminOutput
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
// @Param input body admin.AdminUpdateInput true "Admin update input"
// @Success 200 {object} admin.AdminOutput
// @Router /admin/{id} [put]
func update(c *gin.Context) {
	var (
		id     int
		input  = new(admin.AdminUpdateInput)
		output *admin.AdminOutput
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

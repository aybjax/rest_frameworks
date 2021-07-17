package controllers

import (
	"net/http"
	"strconv"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Train struct {
	ID int `json:id`
	DriverName string `json:driver_name`
	OperatingStatus bool `json:operating_status`
}

func (c *App) GetTrain() revel.Result {
	var train Train

	// id := c.Params.Route.Get("train-id")

	train.ID, _ = strconv.Atoi("1")
	train.DriverName = "Logan" // typa dbdan keled
	train.OperatingStatus = true // typa dbdan keled

	c.Response.Status = http.StatusOK

	return c.RenderJSON(train)
}

func (c *App) CreateTrain() revel.Result {
	var train Train

	c.Params.BindJSON(&train)

	// DB stuff

	train.ID = 2

	c.Response.Status = http.StatusCreated

	return c.RenderJSON(train)
}


func (c *App) RemoveTrain() revel.Result {
	var train Train

	id := c.Params.Route.Get("train-id")

	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "Logan" // typa dbdan keled
	train.OperatingStatus = true // typa dbdan keled

	c.Response.Status = http.StatusOK

	return c.RenderText("")
}


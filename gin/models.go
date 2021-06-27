package main

import (
	"database/sql"
)

var DB *sql.DB

type Train struct {
	ID int
	DriverName string
	OperatingStatus bool
}

type Station struct {
	ID int `json: "id"`
	Name string `json: "name"`
	OpeningTime string `json: "opening_time"`
	ClosingTime string `json: "closing_time"`
}

type Schedule struct {
	ID int
	TrainID int
	StationID int
	ArrivalTime string
}
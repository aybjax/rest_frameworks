package main

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type Train struct {
	ID int
	DriverName string
	OperatingStatus bool
}

type Station struct {
	ID int
	Name string
	OpeningTime time.Time
	ClosingTime time.Time
}

type Schedule struct {
	ID int
	TrainID int
	StationID int
	ArrivalTime time.Time
}
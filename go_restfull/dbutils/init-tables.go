package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverErr := dbDriver.Prepare(train)

	if driverErr != nil {
		log.Fatal(driverErr)
	}

	//create train
	_, statementErr := statement.Exec()

	if statementErr != nil {
		log.Println("train Table alredy exists")
	}


	statement, driverErr = dbDriver.Prepare(station)

	if driverErr != nil {
		log.Fatal(driverErr)
	}

	//create station
	_, statementErr = statement.Exec()

	if statementErr != nil {
		log.Println("station Table alredy exists")
	}


	statement, driverErr = dbDriver.Prepare(schedule)

	if driverErr != nil {
		log.Fatal(driverErr)
	}

	//create schedule
	_, statementErr = statement.Exec()

	if statementErr != nil {
		log.Println("schedule Table alredy exists")
	}

	log.Println("All tables created/initialized")
}
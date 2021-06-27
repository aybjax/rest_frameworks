package main

import (
	"database/sql"
	"gin/dbutils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect db
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Fatal("Driver creation failed")
	}

	dbutils.Initialize(DB)

	r := gin.Default()

	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)

	r.Run(":8080")
}

// GET station
func GetStation(c *gin.Context) {
	var station Station
	
	id := c.Param("station_id")

	err := DB.QueryRow(`SELECT ID, NAME, CAST(OPENING_TIME AS CHAR),
						CAST(CLOSING_TIME AS CHAR) FROM station WHERE ID=?`, id).
			Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H {
			"result": station,
		})
	}
}

// POST create station
func CreateStation(c *gin.Context) {
	var station Station

	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare(`INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME)
					VALUES (?,?,?)`)
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)

		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			newID, _ := result.LastInsertId()

			station.ID = int(newID)

			c.JSON(http.StatusCreated, gin.H {
				"result": station,
			})
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())	
	}
}


// DELETE /v1/station/1
func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, _ := DB.Prepare(`DELETE FROM station WHERE ID=?`)
	_, err := statement.Exec(id)

	if err != nil {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}
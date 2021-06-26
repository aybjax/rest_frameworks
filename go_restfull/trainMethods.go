package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

// GET /v1/trains/1
func (t *Train) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")

	err := DB.QueryRow(`SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE ID=?`, id).
			Scan(&t.ID, &t.DriverName, &t.OperatingStatus)

	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		response.WriteEntity(t)
	}
}

// POST /v1/trains
func (t *Train) createTrain(request *restful.Request, response *restful.Response) {
	log.Printf("%#v\n", request.Request.Body)

	decoder := json.NewDecoder(request.Request.Body)

	var train Train

	err := decoder.Decode(&train)

	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

	statement, err := DB.Prepare(`INSERT INTO train (DRIVER_NAME, OPERATING_STATUS)
					VALUES (?,?)`)

	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

	result, err := statement.Exec(train.DriverName, train.OperatingStatus)


	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		newID, _ := result.LastInsertId()

		train.ID = int(newID)

		response.WriteHeaderAndEntity(http.StatusCreated, train)
	}
}

// DELETE /v1/trains/1
func (t *Train) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare(`DELETE FROM train WHERE ID=?`)
	_, err := statement.Exec(id)

	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeader(http.StatusOK)
	}
}
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func main() {
	webservice := new(restful.WebService)

	webservice.Route(webservice.GET("/ping").To(pingTime))

	restful.Add(webservice)

	http.ListenAndServe(":8080", nil)
}

func pingTime(req *restful.Request, res *restful.Response) {
	io.WriteString(res, fmt.Sprintf("%s", time.Now()))
}
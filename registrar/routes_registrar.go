package registrar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Route struct {
	Name       string                 `json:"name"`
	Path       string                 `json:"path"`
	Method     string                 `json:"method"`
	IsPublic   bool                   `json:"isPublic"`
	HandleFunc func(ctx *gin.Context) `json:"-"`
}

type Routes struct {
	GroupPath string
	Routes    []Route
}

type RegisterRequest struct {
	Service string  `json:"service"`
	Routes  []Route `json:"routes"`
}

func (r *Routes) ToGin(router *gin.Engine) *gin.Engine {
	group := router.Group(r.GroupPath)
	{
		for _, route := range r.Routes {
			group.Handle(route.Method, route.Path, route.HandleFunc)
		}
	}
	//send to register service
	return router
}

func (r *Routes) SendToService() {
	url := os.Getenv("AUTH_SERVICE_URL")
	service := os.Getenv("SERVICE_NAME")
	register, err := strconv.Atoi(os.Getenv("REGISTER_AUTH_SERVICE"))
	if err != nil {
		log.Println(err)
		return
	}
	if register == 0 {
		log.Println("no registration required")
		return
	}
	//jsonData, err := json.Marshal(r.Routes)
	registerRequest := RegisterRequest{
		Service: service,
		Routes:  r.Routes,
	}
	jsonData, err := json.Marshal(registerRequest)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(jsonData))
	request, err := http.NewRequest("POST", fmt.Sprintf("%s", url), bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
}

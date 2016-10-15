package main

import (
	"net/http"
	"getsocial/router"
	ctrl "getsocial/controller"
	"github.com/go-playground/lars"
)

func main() {
	router := InitApp()
	http.ListenAndServe(":3030", router.Serve())
}

func InitApp() *lars.LARS {
	router := router.GetRouter()
	ec := ctrl.NewEventController()
	router.Post("/write", ec.PushData)
	router.Get("/read", ec.GetData)

	// By the REST ideology here should be GET method with route /readbytype/:dataType/:start/:end
	// but for brevity let's reduce by POST method and json data as the input parameters
	router.Post("/readbytype", ec.GetDataByType)

	return router
}



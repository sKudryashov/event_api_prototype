package app

import (
	"github.com/sKudryashov/social_event_api_prototype/router"
	ctrl "github.com/sKudryashov/social_event_api_prototype/controller"
	"github.com/go-playground/lars"
)

// InitApp Initialization of the app and router
//
// As router is used lars router as you can see
func InitApp() *lars.LARS {
	router := router.GetRouter()
	ec := ctrl.NewEventController()
	router.Post("/add", ec.PushData)
	router.Get("/read", ec.GetData)

	// By the REST ideology here should be the GET method with route /readbytype/:dataType/:start/:end
	// but for brevity let's reduce by POST method and json data as the input parameters
	router.Post("/readbytype", ec.GetDataByType)
	router.Get("/readbytimerange/:start/:end", ec.GetDataByRange)

	return router
}

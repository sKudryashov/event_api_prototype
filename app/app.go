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
	router.Post("/add", appMiddleware(ec.PushData))
	router.Get("/read", appMiddleware(ec.GetData))

	// By the REST ideology here should be the GET method with route /readbytype/:dataType/:start/:end
	// but for brevity let's reduce by POST method and json data as the input parameters
	router.Post("/readbytype", appMiddleware(ec.GetDataByType))
	router.Get("/readbytimerange/:start/:end", appMiddleware(ec.GetDataByRange))

	return router
}
// appMiddleware implements any app middleware calls,
// for example net/context or any other logic you can wrap controllers call in
// In our case we handle there error logic
func appMiddleware(f func (c *router.MyContext)) func (c *router.MyContext) {
	appError := new(ctrl.EventError)
	return func (c *router.MyContext) {
		err := f(c)
		if err != nil {
			appError["error"] = err
		}
	}
}

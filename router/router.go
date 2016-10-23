package router

import (
	"github.com/go-playground/lars"
	"time"
	"log"
	"net/http"
	"os"
	"github.com/sKudryashov/social_event_api_prototype/model"
)

func GetRouter() *lars.LARS {
	router := lars.New()
	router.RegisterContext(newContext)
	router.RegisterCustomHandler(func(*MyContext) {}, castCustomContext)
	router.Use(Logger)

	return router
}

func Logger(c lars.Context) {
	start := time.Now()
	c.Next()
	stop := time.Now()
	path := c.Request().URL.Path
	if path == "" {
		path = "/"
	}
	log.Printf("%s %d %s %s", c.Request().Method, c.Response().Status(), path, stop.Sub(start))
}

func castCustomContext(c lars.Context, handler lars.Handler) {
	h := handler.(func(*MyContext))
	ctx := c.(*MyContext)
	h(ctx)
}

type ApplicationGlobals struct {
	Log *log.Logger
	Storage *model.Storage
}

func newGlobals() *ApplicationGlobals {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	return &ApplicationGlobals{
		Log: logger,
		Storage: model.Init(),
	}
}

type MyContext struct {
	*lars.Ctx
	AppContext *ApplicationGlobals
}

func newContext(l *lars.LARS) lars.Context {
	return &MyContext{
		Ctx:        lars.NewContext(l),
		AppContext: newGlobals(),
	}
}

func (mc *MyContext) RequestStart(w http.ResponseWriter, r *http.Request) {
	mc.Ctx.RequestStart(w, r)
}

func (mc *MyContext) RequestEnd() {
	mc.Ctx.RequestEnd()
}

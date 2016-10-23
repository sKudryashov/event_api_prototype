package controller

import (
	"testing"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"github.com/sKudryashov/go-playground/lars"
	"github.com/sKudryashov/social_event_api_prototype/model"
)

func TestMain(m *testing.M) {
	m.Run()
}

type ApplicationGlobals struct {
	Storage *model.Storage
}

func initContext() *router.MyContext {
	ctx := lars.New()
	context := &router.MyContext{
		Ctx:        lars.NewContext(ctx),
		AppContext: newGlobals(),
	}

	return context
}

func newGlobals() *ApplicationGlobals {

	return &ApplicationGlobals{
		Storage: model.Init(),
	}
}

func TestEventController_PushData(t *testing.T) {
	ec := new(EventController)
	context := initContext()
	err := ec.PushData(context)
	if err!=nil {
		t.Fatal("Push data controller error")
	}
}
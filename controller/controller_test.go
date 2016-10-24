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

type Storage struct {}

// GetAllEvents mock
func (s* Storage) GetAllEvents() (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// AddEvent mock
func (s* Storage) AddEvent (ed *model.EventData) error {
	return nil
}

// GetEvents mock
func (s* Storage) GetEvents(eventType string)  (rm *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// GetEventsByRange mock
func (s* Storage) GetEventsByRange (start, end int) (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{
	}
	return &responseModel, nil
}

// initTestModel initializes a fake model
func initTestModel () *Storage {
	return &Storage{}
}

// initContext initializes context mock
func initContext() *router.MyContext {
	ctx := lars.New()
	context := &router.MyContext{
		Ctx:        lars.NewContext(ctx),
		AppContext: newGlobals(),
	}

	return context
}
// newGlobals initializes globals for our controller
func newGlobals() *ApplicationGlobals {
	return &ApplicationGlobals{
		Storage: initTestModel(),
	}
}

func TestEventController_PushData(t *testing.T) {
	ec := new(EventController)
	context := initContext()
	context.Ctx.Request().Body = "something"
	err := ec.PushData(context)
	if err!=nil {
		t.Fatal("Push data controller error")
	}
}
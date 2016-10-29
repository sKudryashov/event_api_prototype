package controller

import (
	"testing"
	//"github.com/sKudryashov/social_event_api_prototype/router"
	"fmt"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/sKudryashov/social_event_api_prototype/model"
)

var (
	context *httptest.Server
	controller *EventController
)

func init()  {
	controller = new(EventController)
	context = initContext()
}

// ApplicationGlobals fake type
type ApplicationGlobals struct {
	Storage model.EventStorage
	Fetcher RequestFetcher
}

type StubReader struct {}
type Storage struct {}
type FetcherTest struct {}

// initTestFetcher initializes a test fetcher
func newTestFetcher() *Fetcher {
	return &FetcherTest{}
}

// initContext initializes context mock
func initContext() *router.MyContext {
	return &router.MyContext {
		AppContext: newGlobals(),
	}
}

// newGlobals initializes globals for our controller
func newGlobals() *ApplicationGlobals {
	return &ApplicationGlobals{
		Storage: newTestModel(),
		Fetcher: newTestFetcher(),
	}
}

// GetRequestBody fetcher stub
func (f FetcherTest) GetRequestBody(c *router.MyContext) ([]byte, error) {
	return []byte("some string"), error("some")
}

// GetStartStopRange fetcher stub
func (f FetcherTest) GetStartStopRange (c *router.MyContext) (string, string, error) {
	return "tik", "tok", error("some")
}

// GetAllEvents storage stub
func (s* Storage) GetAllEvents() (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// AddEvent storage stub
func (s* Storage) AddEvent (ed *model.EventData) error {
	return nil
}

// GetEvents storage stub
func (s* Storage) GetEvents(eventType string)  (rm *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// GetEventsByRange storage stub
func (s* Storage) GetEventsByRange (start, end int) (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// initTestModel initializes a fake storage
func newTestModel() *Storage {
	return &Storage{}
}

// Read is a mock for reader
func (r *StubReader) Read(p []byte) (n int, err error) {
	fmt.Println("A reader has been called")
	return 22, nil
}

/**
 * ==[ Tests ]==
 *
 * go test -v -run=EventController_PushData
 */
func TestEventController_PushData(t *testing.T) {
	assert.New(t)
	fmt.Println("Test is run")
	err := controller.PushData(context)
	if err != nil {
		t.Fatalf("TestEventController_PushData failed %s", err.Error())
	}
	//recorder := httptest.NewRecorder()
	//fake := gofake.New()
	//fake.Stub("getRequestBody").Returning()
	//reader := new(StubReader)
	//request := httptest.NewRequest("POST", "/add", reader)
	//request.Body
	//recorder.Body
	//ec := new(EventController)
	//context := initContext()
	//err := ec.PushData(context)
	//if err != nil {
	//	t.Fatal("Push data controller error")
	//}
}

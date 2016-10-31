package controller

import (
	"testing"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"github.com/stretchr/testify/assert"
	"github.com/sKudryashov/social_event_api_prototype/model"
	"net/http/httptest"
)

var (
	controller *EventController
	request []byte
	start int
	stop int
)

func init() {
	controller = new(EventController)
}

type Storage struct {
}
type TestResponseWriter struct {
}

func (tw TestResponseWriter) WriteSuccess(_ *router.MyContext, _[]byte) (int, error) {
	var err error
	return 23, err
}

func (tw TestResponseWriter) WriteNotFound(_ *router.MyContext, _ string) (int, error) {
	var err error
	return 23, err
}

func (tw TestResponseWriter) WriteForbidden(_ *router.MyContext, _ string) (int, error) {
	var err error
	return 23, err
}

type TestContext interface {
	Response() *httptest.ResponseRecorder
}

// initContext initializes context mock
func initContext() *router.MyContext {
	return &router.MyContext{
		AppContext: newTestGlobals(),
	}
}

func newResponseWriter() *TestResponseWriter {
	return new(TestResponseWriter)
}

// newGlobals initializes globals for our controller
func newTestGlobals() *router.ApplicationGlobals {
	return &router.ApplicationGlobals{
		Storage: newTestModel(),
		Fetcher: newTestFetcher(),
		Writer: newResponseWriter(),
	}
}

type FetcherTest struct{}

// initTestFetcher initializes a test fetcher
func newTestFetcher() *FetcherTest {
	return &FetcherTest{}
}

// GetRequestBody fetcher stub
func (f FetcherTest) GetRequestBody(c router.MyContext) ([]byte, error) {
	var err error
	return request, err
}

// GetStartStopRange fetcher stub
func (f FetcherTest) GetStartStopRange(c router.MyContext) (int, int, error) {
	var err error
	return start, stop, err
}

// GetAllEvents storage stub
func (s*Storage) GetAllEvents() (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// AddEvent storage stub
func (s*Storage) AddEvent(ed *model.EventData) error {
	return nil
}

// GetEvents storage stub
func (s*Storage) GetEvents(eventType string) (rm *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// GetEventsByRange storage stub
func (s*Storage) GetEventsByRange(start, end int) (ed *[]model.EventData, err error) {
	responseModel := []model.EventData{}
	return &responseModel, nil
}

// initTestModel initializes a fake storage
func newTestModel() *Storage {
	return &Storage{}
}

/**
 * ==[ Tests ]==
 *
 * go test -v -run=EventController_PushData
 */
func TestEventController_PushData(t *testing.T) {
	assert.New(t)
	request = []byte(`{"eventType":"Usual","sessionStart":1476628565,"sessionEnd":1476628965,"linkClicked":"https://blog.golang.org/c-go-cgo","timestamp":12039109203,"params":{"C":"c++","D":"D++","R":"R is not a real language"}}`)
	err := controller.PushData(initContext())
	if err != nil {
		t.Error("TestEventController_PushData failed -> ", err.Error())
	}
}

func TestEventController_GetData(t *testing.T) {
	assert.New(t)
	err := controller.GetData(initContext())
	if err != nil {
		t.Error("TestEventController_GetData failed -> ", err.Error())
	}
}

func TestEventController_GetDataByType(t *testing.T) {
	assert.New(t)
	request = []byte(`{"eventType":"Usual"}`)
	err := controller.GetDataByType(initContext())
	if err != nil {
		t.Error("TestEventController_GetDataByType failed -> ", err.Error())
	}
}

func TestEventController_GetDataByRange(t *testing.T) {
	assert.New(t)
	err := controller.GetDataByRange(initContext())
	if err != nil {
		t.Error("TestEventController_GetDataByRange failed -> ", err.Error())
	}
}
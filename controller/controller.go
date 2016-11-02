package controller

import (
	"gopkg.in/go-playground/validator.v8"
	"github.com/sKudryashov/social_event_api_prototype/model"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"encoding/json"
	"github.com/pkg/errors"
)

// EventController serves events logic of the application
type EventController struct {
	error EventError
}
// EventError this is the custom error type of event handling
type EventError struct {
	Message string
	Err error
}

// NewEventController initialization of the controller
func NewEventController() *EventController {
	return &EventController{}
}

// PushData adding data to a storage (whatever it is)
func (ec *EventController) PushData (c *router.MyContext) error {
	data, _ := c.AppContext.Fetcher.GetRequestBody(*c)
	request := model.EventData{}
	validate := ec.getValidator()

	if err := json.Unmarshal(data, &request); err != nil {
		return errors.Wrap( err, "Unmarshalling error")
	}

	if err := validate.Struct(request); err != nil {
		return errors.Wrap(err, "Validation error")
	}

	if err := c.AppContext.Storage.AddEvent(&request); err != nil {
		return errors.Wrap(err, "Data recording error")
	}

	_, err := c.AppContext.Writer.WriteSuccess(c, []byte("Data has been written successfully"))

	if err != nil {
		return errors.Wrap(err, "Data writer error")
	}

	return nil
}

// GetData returns the whole dataset
func (ec *EventController) GetData (c *router.MyContext) error {
	responseModel, err := c.AppContext.Storage.GetAllEvents()

	if err != nil {
		return errors.Wrap(err, "Db fetching error")
	}

	dataFoundJson, err := json.Marshal(responseModel)

	if err != nil {
		return errors.Wrap( err, "Unmarshalling error")
	}

	_, err = c.AppContext.Writer.WriteSuccess(c, dataFoundJson)

	if err != nil {
		//todo: move this out to constants
		return errors.Wrap(err, "Response error")
	}

	return nil
}

// GetDataByType Fetching data by event type from storage
func (ec *EventController) GetDataByType(c *router.MyContext) error {
	data, _ := c.AppContext.Fetcher.GetRequestBody(*c)
	request := model.FetchBy{}

	if err := json.Unmarshal(data, &request); err != nil {
		return errors.Wrap(err, "Unmarshalling error")
	}

	validate := ec.getValidator()

	if err := validate.Struct(request); err != nil {
		return errors.Wrap(err, "Validation error")
	}

	events, err := c.AppContext.Storage.GetEvents(request.Type)

	if err != nil {
		return errors.Wrap(err, "Data fetching error")
	}

	dataFoundJson, err := json.Marshal(events)

	if err != nil {
		return errors.Wrap(err, "Unmarshalling error")
	}
	_, err = c.AppContext.Writer.WriteSuccess(c, dataFoundJson)
	if err != nil {
		//todo: move this out to constants
		return errors.Wrap(err, "Response error")
	}
	return nil
}

// GetDataByRange returns data in a given time range
func (ec *EventController) GetDataByRange(c *router.MyContext) error {
	var start, end int
	var err error

	start, end, errorFetch := c.AppContext.Fetcher.GetStartStopRange(*c)
	if errorFetch != nil {
		return errors.Wrap(errorFetch, "Wrong URL")
	}

	responseModel, err := c.AppContext.Storage.GetEventsByRange(start, end)
	if err != nil {
		return errors.Wrap(err, "Storage error")
	}

	dataFoundJson, err := json.Marshal(responseModel)
	_, err = c.AppContext.Writer.WriteSuccess(c, dataFoundJson)

	if err != nil {
		//todo: move this out to constants
		return errors.Wrap(err, "Response error")
	}

	return nil
}

func (ec *EventController) getValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")

	return validate
}

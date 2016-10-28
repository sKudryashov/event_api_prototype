package controller

import (
	"io/ioutil"
	"gopkg.in/go-playground/validator.v8"
	"github.com/sKudryashov/social_event_api_prototype/model"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"encoding/json"
	"github.com/go-playground/lars"
	"strconv"
	"net/http"
	"github.com/pkg/errors"
)

// EventController serves events logic of the application
type EventController struct {
	error EventError
}
// EventError this is the custom error type of event handling
type EventError struct {
<<<<<<< HEAD
	message string
	err error
=======
	Message string
	Err error
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
}

// NewEventController initialization of the controller
func NewEventController() *EventController {
	return &EventController{}
}

// PushData adding data to a storage (whatever it is)
func (ec *EventController) PushData (c *router.MyContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.EventData{}
	validate := ec.getValidator()

	if err := json.Unmarshal(data, &request); err != nil {
<<<<<<< HEAD
		return errors.Wrap(err, "Unmarshalling error")
=======
		return errors.Wrap( err, "Unmarshalling error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	if err := validate.Struct(request); err != nil {
		return errors.Wrap(err, "Validation error")
	}

<<<<<<< HEAD
	if err := c.AppContext.Storage.AddEvent(request); err != nil {
=======
	if err := c.AppContext.Storage.AddEvent(&request); err != nil {
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
		return errors.Wrap(err, "Data recording error")
	}

	rsp := ec.getSuccessWriter(c)
	rsp.Write([]byte("Data has been written successfully"))
	c.AppContext.Log.Println("Data has been written successfully")

	return nil
}

// GetData returns the whole dataset
func (ec *EventController) GetData (c *router.MyContext) error {
	responseModel, err := c.AppContext.Storage.GetAllEvents()

	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Error with db fetching: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Db fetching error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)

	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Error with unmarshalling: " + err.Error())
		return err
=======
		return errors.Wrap( err, "Unmarshalling error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}
	rsp.Write([]byte(dataFoundJson))

	return nil
}

// GetDataByType Fetching data by event type from storage
func (ec *EventController) GetDataByType(c *router.MyContext) error {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.FetchBy{}

	if err := json.Unmarshal(data, &request); err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Error with json unmarshalling: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Unmarshalling error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	validate := ec.getValidator()

	if err := validate.Struct(request); err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Error with validation: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Validation error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	events, err := c.AppContext.Storage.GetEvents(request.Type)

	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Data fetching error:" + err.Error())
		return err
=======
		return errors.Wrap(err, "Data fetching error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(events)

	if err != nil {
<<<<<<< HEAD
		return err
		//c.AppContext.Log.Println("Unmarshalling error: " + err.Error())
=======
		return errors.Wrap(err, "Unmarshalling error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	rsp.Write([]byte(dataFoundJson))

	return nil
}

// GetDataByRange returns data in a given time range
func (ec *EventController) GetDataByRange(c *router.MyContext) error {
	var start, end int
	var err error

	start, err = strconv.Atoi(c.Ctx.Param("start"))
	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Wrong URL start parameter: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Wrong URL")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	end, err = strconv.Atoi(c.Ctx.Param("end"))
	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Wrong URL end parameter: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Wrong URL")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	responseModel, err := c.AppContext.Storage.GetEventsByRange(start, end)

	if err != nil {
<<<<<<< HEAD
		//c.AppContext.Log.Println("Error with db fetching: " + err.Error())
		return err
=======
		return errors.Wrap(err, "Storage error")
>>>>>>> 8a2db1083740273a2ca834cba8f4b9f8b1b8acbc
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)
	rsp.Write([]byte(dataFoundJson))

	return nil
}

func (ec *EventController) getValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")

	return validate
}

func (ec *EventController) getSuccessWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusOK)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

func (ec *EventController) getErrorNotFoundWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusNotFound)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

// Returns writer for HTTP forbidden
func (ec *EventController) getErrorForbiddenWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusForbidden)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

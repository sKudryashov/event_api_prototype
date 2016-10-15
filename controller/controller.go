package controller

import (
	"io/ioutil"
	"gopkg.in/go-playground/validator.v8"
	"getsocial/model"
	"getsocial/router"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/go-playground/lars"
)

type EventController struct {

}

func NewEventController() *EventController {
	return &EventController{}
}

func (ec *EventController) PushData (c *router.MyContext) {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.EventData{}
	validate := ec.getValidator()

	if err := json.Unmarshal(data, &request); err != nil {
		c.AppContext.Log.Println("Error with json unmarshalling: " + err.Error())
		panic(err)
	}

	if err := validate.Struct(request); err != nil {
		c.AppContext.Log.Println("Error with validation: " + err.Error())
		panic(err)
	}

	request.EventId = bson.NewObjectId()
	c.AppContext.Storage.DB("event_model").C("events").Insert(request)
	rsp := ec.getSuccessWriter(c)
	rsp.Write([]byte("Data has been written successfully"))
	c.AppContext.Log.Println("Successfull response")
}

func (ec *EventController) GetData (c *router.MyContext)  {
	responseModel := make([]model.EventData, 0, 3)
	err := c.AppContext.Storage.DB("event_model").C("events").Find(nil).All(&responseModel)

	if err != nil {
		c.AppContext.Log.Println("Error with db fetching: " + err.Error())
		panic(err)
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)

	if err != nil {
		c.AppContext.Log.Println("Error with unmarshalling: " + err.Error())
	}
	rsp.Write([]byte(dataFoundJson))
}

func (ec *EventController) GetDataByType(c *router.MyContext) {
	data, _ := ioutil.ReadAll(c.Request().Body)
	request := model.FetchBy{}

	if err := json.Unmarshal(data, &request); err != nil {
		c.AppContext.Log.Println("Error with json unmarshalling: " + err.Error())
		panic(err)
	}

	validate := ec.getValidator()

	if err := validate.Struct(request); err != nil {
		c.AppContext.Log.Println("Error with validation: " + err.Error())
		panic(err)
	}

	responseModel := make([]model.EventData, 0, 10)

	err := c.AppContext.Storage.DB("event_model").C("events").Find(bson.M{"eventType": "Usual"}).All(&responseModel)
	if err != nil {
		c.AppContext.Log.Println("Error with db fetching: " + err.Error())
		panic(err)
	}

	rsp := ec.getSuccessWriter(c)
	dataFoundJson, err := json.Marshal(responseModel)

	if err != nil {
		c.AppContext.Log.Println("Error with unmarshalling: " + err.Error())
	}
	rsp.Write([]byte(dataFoundJson))
}

func (ec *EventController) getValidator() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("validate")

	return validate
}

func (ec *EventController) getSuccessWriter(c *router.MyContext) *lars.Response {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(200)
	rsp.Header().Set("Content-Type", "application/json")

	return rsp
}

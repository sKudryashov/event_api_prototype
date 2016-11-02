package model

import (
	"gopkg.in/mgo.v2/bson"
	mongo "gopkg.in/mgo.v2"
)

type EventData struct{
	EventId bson.ObjectId `json:"id" bson:"_id"`
	EventType string `json:"eventType" validate:"required,alphanum" bson:"eventType"`
	SessionStart int `json:"sessionStart" validate:"required" bson:"SessionStart"`
	SessionEnd int `json:"sessionEnd" validate:"required" bson:"SessionEnd"`
	LinkClicked string `json:"linkClicked" validate:"required,url" bson:"LinkClicked"`
	Ts int `json:"timestamp" validate:"required" bson:"timestamp"`
	Params map[string]interface{} `json:"params" bson:"fulltext"`
}

type FetchBy struct {
	Start string `json:"sessionStart"`
	End string `json:"sessionEnd"`
	Type string `json:"eventType" validate:"required"`
}

type Storage struct {
	session *mongo.Session
	db *mongo.Database
}

type EventStorage interface {
	GetAllEvents() (ed *[]EventData, err error)
	AddEvent (ed *EventData) error
	GetEvents (eventType string) (rm *[]EventData, err error)
	GetEventsByRange (start, end int) (ed *[]EventData, err error)
}

// Init initializes application storage
func Init() *Storage {
	storage := new(Storage)
	storage.initSession()

	return storage
}

// InitSession initialization of storage connection
func (s* Storage) initSession() {
	var err error
	s.session, err = mongo.Dial("mongodb://localhost")
	s.db = s.session.DB("event_model")

	if err != nil {
		panic(err)
	}
}

// GetAllEvents returns all dataset from storage
func (s* Storage) GetAllEvents() (ed *[]EventData, err error) {
	query := s.db.C("events").Find(nil)
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	responseModel := make([]EventData, 0, count)
	query.All(&responseModel)

	return &responseModel, err
}

// AddEvent adds an event into the storage
func (s* Storage) AddEvent (ed *EventData) error {
	ed.EventId = bson.NewObjectId()
	return s.db.C("events").Insert(ed)
}

// GetEvents returns an events slice by event type
func (s* Storage) GetEvents(eventType string) (rm *[]EventData, err error) {
	var responseModel []EventData

	findBy := bson.M{"eventType": eventType}
	query := s.db.C("events").Find(findBy);
	count, err := query.Count()

	if err == nil {
		responseModel = make([]EventData, 0, count)
		err = query.All(&responseModel)
	}

	return &responseModel, err
}

// GetEventsByRange returns event fetched by given range
func (s* Storage) GetEventsByRange (start, end int) (ed *[]EventData, err error) {
	findBy := bson.M{"sessionStart": bson.M{ "$gte": start}, "sessionEnd":bson.M{"$lte":end} }
	query := s.db.C("events").Find(findBy)
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	responseModel := make([]EventData, 0, count)
	query.All(&responseModel)

	return &responseModel, err
}


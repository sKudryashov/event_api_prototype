package model

import (
	"gopkg.in/mgo.v2/bson"
)

type EventData struct{
	EventId bson.ObjectId `json:"id" bson:"_id"`
	EventType string `json:"eventType" validate:"required,alphanum" bson:"eventType"`
	SessionStart string `json:"sessionStart" validate:"required" bson:"SessionStart"`
	SessionEnd string `json:"sessionEnd" validate:"required" bson:"SessionEnd"`
	LinkClicked string `json:"linkClicked" validate:"required" bson:"LinkClicked"`
	Ts int `json:"timestamp" validate:"required" bson:"timestamp"`
	Params map[string]interface{} `json:"params" bson:"fulltext"`
}

type FetchBy struct {
	Start string `json:"sessionStart"`
	End string `json:"sessionEnd"`
	Type string `json:"eventType" validate:"required"`
}


package controller

import (
	"io/ioutil"
	"github.com/sKudryashov/social_event_api_prototype/router"
	"strconv"
)

type Fetcher struct{}

// RequestFetcher is an interface which is used in context to
//
type RequestFetcher interface {
	GetRequestBody(c *router.MyContext) ([]byte, error)
	GetStartStopRange (c *router.MyContext) (string, string, error)
}

// NewFetcher returns an instance of the standard fetcher
func NewFetcher() *Fetcher {
	return new(Fetcher)
}

// GetRequestBody returns a reference to a byte slice request body
func (f Fetcher) GetRequestBody(c *router.MyContext) ([]byte, error)  {
	data, err := ioutil.ReadAll(c.Request().Body)
	return data, err
}

// GetStartStopRange return start and stop range form URL request
func (f Fetcher) GetStartStopRange (c *router.MyContext) (string, string, error) {
	var start, end string
	var err error

	start, err = strconv.Atoi(c.Ctx.Param("start"))
	if err != nil {
		return nil, nil, err
	}
	end, err = strconv.Atoi(c.Ctx.Param("end"))
	if err != nil {
		return nil, nil, err
	}

	return start, end, nil
}
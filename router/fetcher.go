package router

import (
	"io/ioutil"
	"strconv"
)

type Fetcher struct{}

// RequestFetcher is an interface which is used in context to
type RequestFetcher interface {
	GetRequestBody(c MyContext) ([]byte, error)
	GetStartStopRange (c MyContext) (int, int, error)
}

// NewFetcher returns an instance of the standard fetcher
func NewFetcher() *Fetcher {
	return new(Fetcher)
}

// GetRequestBody returns a reference to a byte slice request body
func (f Fetcher) GetRequestBody(c MyContext) ([]byte, error)  {
	data, err := ioutil.ReadAll(c.Request().Body)
	return data, err
}

// GetStartStopRange return start and stop range form URL request
func (f Fetcher) GetStartStopRange (c MyContext) (int, int, error) {
	var start, end int
	var err error

	start, err = strconv.Atoi(c.Ctx.Param("start"))
	if err != nil {
		return start, end, err
	}
	end, err = strconv.Atoi(c.Ctx.Param("end"))
	if err != nil {
		return start, end, err
	}

	return start, end, err
}
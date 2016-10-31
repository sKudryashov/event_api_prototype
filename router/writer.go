package router

import (
	"github.com/sKudryashov/switch/src/router"
	"net/http"
)

type ResponseWriter interface {
	WriteSuccess(*router.MyContext, string) (int, error)
	WriteNotFound(*router.MyContext, string) (int, error)
	WriteForbidden(*router.MyContext, string) (int, error)
}

// NewAppResponseWriter returns the app writer
func NewAppResponseWriter() *AppResponseWriter {
	return new(AppResponseWriter)
}

// AppResponseWriter manages by all writer operations in the app
type AppResponseWriter struct {
}

func (wr AppResponseWriter) WriteSuccess(c *router.MyContext, response string) (int, error) {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusOK)
	rsp.Header().Set("Content-Type", "application/json")
	return rsp.Write([]byte(response))
}

func (wr AppResponseWriter) WriteNotFound(c *router.MyContext, response string) (int, error)  {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusNotFound)
	rsp.Header().Set("Content-Type", "text")
	return rsp.Write([]byte(response))
}

func (wr AppResponseWriter)  WriteForbidden(c *router.MyContext, response string) (int, error) {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusForbidden)
	rsp.Header().Set("Content-Type", "text")
	return rsp.Write([]byte(response))
}

//func (ec *EventController) GetSuccessWriter(c *router.MyContext) *lars.Response {
//	rsp := c.Ctx.Response()
//	rsp.WriteHeader(http.StatusOK)
//	rsp.Header().Set("Content-Type", "application/json")
//
//	return rsp
//}
//
//func (ec *EventController) GetErrorNotFoundWriter(c *router.MyContext) *lars.Response {
//	rsp := c.Ctx.Response()
//	rsp.WriteHeader(http.StatusNotFound)
//	rsp.Header().Set("Content-Type", "application/json")
//
//	return rsp
//}
//
//// Returns writer for HTTP forbidden
//func (ec *EventController) GetErrorForbiddenWriter(c *router.MyContext) *lars.Response {
//	rsp := c.Ctx.Response()
//	rsp.WriteHeader(http.StatusForbidden)
//	rsp.Header().Set("Content-Type", "application/json")
//
//	return rsp
//}
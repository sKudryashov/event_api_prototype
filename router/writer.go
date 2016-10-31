package router

import (
	"net/http"
)

type ResponseWriter interface {
	WriteSuccess(*MyContext, []byte) (int, error)
	WriteNotFound(*MyContext, string) (int, error)
	WriteForbidden(*MyContext, string) (int, error)
}

// NewAppResponseWriter returns the app writer
func NewAppResponseWriter() *AppResponseWriter {
	return new(AppResponseWriter)
}

// AppResponseWriter manages by all writer operations in the app
type AppResponseWriter struct {
}

func (wr AppResponseWriter) WriteSuccess(c *MyContext, response []byte) (int, error) {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusOK)
	rsp.Header().Set("Content-Type", "application/json")
	return rsp.Write(response)
}

func (wr AppResponseWriter) WriteNotFound(c *MyContext, response string) (int, error)  {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusNotFound)
	rsp.Header().Set("Content-Type", "text")
	return rsp.Write([]byte(response))
}

func (wr AppResponseWriter)  WriteForbidden(c *MyContext, response string) (int, error) {
	rsp := c.Ctx.Response()
	rsp.WriteHeader(http.StatusForbidden)
	rsp.Header().Set("Content-Type", "text")
	return rsp.Write([]byte(response))
}
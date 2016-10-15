package lars

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router.
// The slice is ordered, the first URL parameter is also the first slice value.
// It is therefore safe to read values by the index.
type Params []Param

// NewContext returns a new default lars Context object.
func NewContext(l *LARS) *Ctx {

	c := &Ctx{
		params: make(Params, l.mostParams),
	}

	c.response = newResponse(nil, c)

	return c
}

// BaseContext returns the underlying context object LARS uses internally.
// used when overriding the context object
func (c *Ctx) BaseContext() *Ctx {
	return c
}

// Request returns context assotiated *http.Request.
func (c *Ctx) Request() *http.Request {
	return c.request
}

// Response returns http.ResponseWriter.
func (c *Ctx) Response() *Response {
	return c.response
}

// WebSocket returns context's assotiated *websocket.Conn.
func (c *Ctx) WebSocket() *websocket.Conn {
	return c.websocket
}

// RequestEnd fires after request completes and just before
// the *Ctx object gets put back into the pool.
// Used to close DB connections and such on a custom context
func (c *Ctx) RequestEnd() {
}

// Param returns the value of the first Param which key matches the given name.
// If no matching Param is found, an empty string is returned.
func (c *Ctx) Param(name string) string {

	for _, entry := range c.params {
		if entry.Key == name {
			return entry.Value
		}
	}

	return blank
}

// QueryParams returns the http.Request.URL.Query() values
// this function is not for convenience, but rather performance
// URL.Query() reparses the RawQuery every time it's called, but this
// function will cache the initial parsing so it doesn't have to reparse;
// which is useful if when accessing these Params from multiple middleware.
func (c *Ctx) QueryParams() url.Values {

	if c.queryParams != nil {
		return c.queryParams
	}

	c.queryParams = c.request.URL.Query()

	return c.queryParams
}

// ParseForm calls the underlying http.Request ParseForm
// but also adds the URL params to the request Form as if
// they were defined as query params i.e. ?id=13&ok=true but
// does not add the params to the http.Request.URL.RawQuery
// for SEO purposes
func (c *Ctx) ParseForm() error {

	if c.formParsed {
		return nil
	}

	if err := c.request.ParseForm(); err != nil {
		return err
	}

	for _, entry := range c.params {
		c.request.Form.Add(entry.Key, entry.Value)
	}

	c.formParsed = true

	return nil
}

// ParseMultipartForm calls the underlying http.Request ParseMultipartForm
// but also adds the URL params to the request Form as if they were defined
// as query params i.e. ?id=13&ok=true but does not add the params to the
// http.Request.URL.RawQuery for SEO purposes
func (c *Ctx) ParseMultipartForm(maxMemory int64) error {

	if c.multipartFormParsed {
		return nil
	}

	if err := c.request.ParseMultipartForm(maxMemory); err != nil {
		return err
	}

	for _, entry := range c.params {
		c.request.Form.Add(entry.Key, entry.Value)
	}

	c.multipartFormParsed = true

	return nil
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in github.
func (c *Ctx) Next() {
	c.index++
	c.handlers[c.index](c.parent)
}

// http response helpers

// JSON marshals provided interface + returns JSON + status code
func (c *Ctx) JSON(code int, i interface{}) (err error) {

	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return c.JSONBytes(code, b)
}

// JSONBytes returns provided JSON response with status code
func (c *Ctx) JSONBytes(code int, b []byte) (err error) {

	c.response.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	c.response.WriteHeader(code)
	_, err = c.response.Write(b)
	return
}

// JSONP sends a JSONP response with status code and uses `callback` to construct
// the JSONP payload.
func (c *Ctx) JSONP(code int, i interface{}, callback string) (err error) {

	b, e := json.Marshal(i)
	if e != nil {
		err = e
		return
	}

	c.response.Header().Set(ContentType, ApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeader(code)

	if _, err = c.response.Write([]byte(callback + "(")); err == nil {

		if _, err = c.response.Write(b); err == nil {
			_, err = c.response.Write([]byte(");"))
		}
	}

	return
}

// XML marshals provided interface + returns XML + status code
func (c *Ctx) XML(code int, i interface{}) error {

	b, err := xml.Marshal(i)
	if err != nil {
		return err
	}

	return c.XMLBytes(code, b)
}

// XMLBytes returns provided XML response with status code
func (c *Ctx) XMLBytes(code int, b []byte) (err error) {

	c.response.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	c.response.WriteHeader(code)

	if _, err = c.response.Write([]byte(xml.Header)); err == nil {
		_, err = c.response.Write(b)
	}

	return
}

// Text returns the provided string with status code
func (c *Ctx) Text(code int, s string) error {
	return c.TextBytes(code, []byte(s))
}

// TextBytes returns the provided response with status code
func (c *Ctx) TextBytes(code int, b []byte) (err error) {

	c.response.Header().Set(ContentType, TextPlainCharsetUTF8)
	c.response.WriteHeader(code)
	_, err = c.response.Write(b)
	return
}

// http request helpers

// ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
func (c *Ctx) ClientIP() (clientIP string) {

	var values []string

	if values, _ = c.request.Header[XRealIP]; len(values) > 0 {

		clientIP = strings.TrimSpace(values[0])
		if clientIP != blank {
			return
		}
	}

	if values, _ = c.request.Header[XForwardedFor]; len(values) > 0 {
		clientIP = values[0]

		if index := strings.IndexByte(clientIP, ','); index >= 0 {
			clientIP = clientIP[0:index]
		}

		clientIP = strings.TrimSpace(clientIP)
		if clientIP != blank {
			return
		}
	}

	clientIP, _, _ = net.SplitHostPort(strings.TrimSpace(c.request.RemoteAddr))

	return
}

// AcceptedLanguages returns an array of accepted languages denoted by
// the Accept-Language header sent by the browser
// NOTE: some stupid browsers send in locales lowercase when all the rest send it properly
func (c *Ctx) AcceptedLanguages(lowercase bool) []string {

	var accepted string

	if accepted = c.request.Header.Get(AcceptedLanguage); accepted == blank {
		return []string{}
	}

	options := strings.Split(accepted, ",")
	l := len(options)

	language := make([]string, l)

	if lowercase {

		for i := 0; i < l; i++ {
			locale := strings.SplitN(options[i], ";", 2)
			language[i] = strings.ToLower(strings.Trim(locale[0], " "))
		}
	} else {

		for i := 0; i < l; i++ {
			locale := strings.SplitN(options[i], ";", 2)
			language[i] = strings.Trim(locale[0], " ")
		}
	}

	return language
}

// HandlerName returns the current Contexts final handler's name
func (c *Ctx) HandlerName() string {
	return c.handlerName
}

// Stream provides HTTP Streaming
func (c *Ctx) Stream(step func(w io.Writer) bool) {
	w := c.response
	clientGone := w.CloseNotify()

	for {
		select {
		case <-clientGone:
			return
		default:
			keepOpen := step(w)
			w.Flush()
			if !keepOpen {
				return
			}
		}
	}
}

// Attachment is a helper method for returning an attachement file
// to be downloaded, if you with to open inline see function
func (c *Ctx) Attachment(r io.Reader, filename string) (err error) {

	c.response.Header().Set(ContentDisposition, "attachment;filename="+filename)
	c.response.Header().Set(ContentType, detectContentType(filename))
	c.response.WriteHeader(http.StatusOK)

	_, err = io.Copy(c.response, r)

	return
}

// Inline is a helper method for returning a file inline to
// be rendered/opened by the browser
func (c *Ctx) Inline(r io.Reader, filename string) (err error) {

	c.response.Header().Set(ContentDisposition, "inline;filename="+filename)
	c.response.Header().Set(ContentType, detectContentType(filename))
	c.response.WriteHeader(http.StatusOK)

	_, err = io.Copy(c.response, r)

	return
}

// Decode takes the request and attempts to discover it's content type via
// the http headers and then decode the request body into the provided struct.
// Example if header was "application/json" would decode using
// json.NewDecoder(io.LimitReader(c.request.Body, maxMemory)).Decode(v).
func (c *Ctx) Decode(includeFormQueryParams bool, maxMemory int64, v interface{}) (err error) {

	initFormDecoder()

	typ := c.request.Header.Get(ContentType)

	if idx := strings.Index(typ, ";"); idx != -1 {
		typ = typ[:idx]
	}

	switch typ {

	case ApplicationJSON:
		err = json.NewDecoder(io.LimitReader(c.request.Body, maxMemory)).Decode(v)

	case ApplicationXML:
		err = xml.NewDecoder(io.LimitReader(c.request.Body, maxMemory)).Decode(v)

	case ApplicationForm:

		if err = c.ParseForm(); err == nil {
			if includeFormQueryParams {
				err = formDecoder.Decode(v, c.request.Form)
			} else {
				err = formDecoder.Decode(v, c.request.PostForm)
			}
		}

	case MultipartForm:

		if err = c.ParseMultipartForm(maxMemory); err == nil {
			if includeFormQueryParams {
				err = formDecoder.Decode(v, c.request.Form)
			} else {
				err = formDecoder.Decode(v, c.request.MultipartForm.Value)
			}
		}
	}
	return
}

// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"strings"
)

func newSQLTranslateFunc(t Transport) SQLTranslate {
	return func(body io.Reader, o ...func(*SQLTranslateRequest)) (*Response, error) {
		var r = SQLTranslateRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at Translate SQL into Elasticsearch queries.
//
type SQLTranslate func(body io.Reader, o ...func(*SQLTranslateRequest)) (*Response, error)

// SQLTranslateRequest configures the Sql Translate API request.
//
type SQLTranslateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SQLTranslateRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_sql/translate"))
	path.WriteString("/_sql/translate")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f SQLTranslate) WithContext(v context.Context) func(*SQLTranslateRequest) {
	return func(r *SQLTranslateRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SQLTranslate) WithPretty() func(*SQLTranslateRequest) {
	return func(r *SQLTranslateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SQLTranslate) WithHuman() func(*SQLTranslateRequest) {
	return func(r *SQLTranslateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SQLTranslate) WithErrorTrace() func(*SQLTranslateRequest) {
	return func(r *SQLTranslateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SQLTranslate) WithFilterPath(v ...string) func(*SQLTranslateRequest) {
	return func(r *SQLTranslateRequest) {
		r.FilterPath = v
	}
}
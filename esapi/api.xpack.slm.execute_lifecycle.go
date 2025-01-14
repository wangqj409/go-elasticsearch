// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newSlmExecuteLifecycleFunc(t Transport) SlmExecuteLifecycle {
	return func(o ...func(*SlmExecuteLifecycleRequest)) (*Response, error) {
		var r = SlmExecuteLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SlmExecuteLifecycle - https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-execute.html
//
type SlmExecuteLifecycle func(o ...func(*SlmExecuteLifecycleRequest)) (*Response, error)

// SlmExecuteLifecycleRequest configures the Slm Execute Lifecycle API request.
//
type SlmExecuteLifecycleRequest struct {
	PolicyID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SlmExecuteLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(1 + len("_slm") + 1 + len("policy") + 1 + len(r.PolicyID) + 1 + len("_execute"))
	path.WriteString("/")
	path.WriteString("_slm")
	path.WriteString("/")
	path.WriteString("policy")
	if r.PolicyID != "" {
		path.WriteString("/")
		path.WriteString(r.PolicyID)
	}
	path.WriteString("/")
	path.WriteString("_execute")

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

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
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
func (f SlmExecuteLifecycle) WithContext(v context.Context) func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicyID - the ID of the snapshot lifecycle policy to be executed.
//
func (f SlmExecuteLifecycle) WithPolicyID(v string) func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.PolicyID = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SlmExecuteLifecycle) WithPretty() func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SlmExecuteLifecycle) WithHuman() func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SlmExecuteLifecycle) WithErrorTrace() func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SlmExecuteLifecycle) WithFilterPath(v ...string) func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f SlmExecuteLifecycle) WithHeader(h map[string]string) func(*SlmExecuteLifecycleRequest) {
	return func(r *SlmExecuteLifecycleRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

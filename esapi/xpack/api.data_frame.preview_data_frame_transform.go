// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newDataFramePreviewDataFrameTransformFunc(t Transport) DataFramePreviewDataFrameTransform {
	return func(body io.Reader, o ...func(*DataFramePreviewDataFrameTransformRequest)) (*Response, error) {
		var r = DataFramePreviewDataFrameTransformRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-data-frame-transform.html.
//
type DataFramePreviewDataFrameTransform func(body io.Reader, o ...func(*DataFramePreviewDataFrameTransformRequest)) (*Response, error)

// DataFramePreviewDataFrameTransformRequest configures the Data Frame    Preview Data Frame Transform API request.
//
type DataFramePreviewDataFrameTransformRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r DataFramePreviewDataFrameTransformRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_data_frame/transforms/_preview"))
	path.WriteString("/_data_frame/transforms/_preview")

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
func (f DataFramePreviewDataFrameTransform) WithContext(v context.Context) func(*DataFramePreviewDataFrameTransformRequest) {
	return func(r *DataFramePreviewDataFrameTransformRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f DataFramePreviewDataFrameTransform) WithPretty() func(*DataFramePreviewDataFrameTransformRequest) {
	return func(r *DataFramePreviewDataFrameTransformRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f DataFramePreviewDataFrameTransform) WithHuman() func(*DataFramePreviewDataFrameTransformRequest) {
	return func(r *DataFramePreviewDataFrameTransformRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f DataFramePreviewDataFrameTransform) WithErrorTrace() func(*DataFramePreviewDataFrameTransformRequest) {
	return func(r *DataFramePreviewDataFrameTransformRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f DataFramePreviewDataFrameTransform) WithFilterPath(v ...string) func(*DataFramePreviewDataFrameTransformRequest) {
	return func(r *DataFramePreviewDataFrameTransformRequest) {
		r.FilterPath = v
	}
}
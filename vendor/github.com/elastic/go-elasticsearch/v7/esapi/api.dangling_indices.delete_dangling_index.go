// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.9.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newDanglingIndicesDeleteDanglingIndexFunc(t Transport) DanglingIndicesDeleteDanglingIndex {
	return func(index_uuid string, o ...func(*DanglingIndicesDeleteDanglingIndexRequest)) (*Response, error) {
		var r = DanglingIndicesDeleteDanglingIndexRequest{IndexUuid: index_uuid}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// DanglingIndicesDeleteDanglingIndex deletes the specified dangling index
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-gateway-dangling-indices.html.
//
type DanglingIndicesDeleteDanglingIndex func(index_uuid string, o ...func(*DanglingIndicesDeleteDanglingIndexRequest)) (*Response, error)

// DanglingIndicesDeleteDanglingIndexRequest configures the Dangling Indices Delete Dangling Index API request.
//
type DanglingIndicesDeleteDanglingIndexRequest struct {
	IndexUuid string

	AcceptDataLoss *bool
	MasterTimeout  time.Duration
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r DanglingIndicesDeleteDanglingIndexRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_dangling") + 1 + len(r.IndexUuid))
	path.WriteString("/")
	path.WriteString("_dangling")
	path.WriteString("/")
	path.WriteString(r.IndexUuid)

	params = make(map[string]string)

	if r.AcceptDataLoss != nil {
		params["accept_data_loss"] = strconv.FormatBool(*r.AcceptDataLoss)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

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
func (f DanglingIndicesDeleteDanglingIndex) WithContext(v context.Context) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.ctx = v
	}
}

// WithAcceptDataLoss - must be set to true in order to delete the dangling index.
//
func (f DanglingIndicesDeleteDanglingIndex) WithAcceptDataLoss(v bool) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.AcceptDataLoss = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f DanglingIndicesDeleteDanglingIndex) WithMasterTimeout(v time.Duration) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f DanglingIndicesDeleteDanglingIndex) WithTimeout(v time.Duration) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f DanglingIndicesDeleteDanglingIndex) WithPretty() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f DanglingIndicesDeleteDanglingIndex) WithHuman() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f DanglingIndicesDeleteDanglingIndex) WithErrorTrace() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f DanglingIndicesDeleteDanglingIndex) WithFilterPath(v ...string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f DanglingIndicesDeleteDanglingIndex) WithHeader(h map[string]string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f DanglingIndicesDeleteDanglingIndex) WithOpaqueID(s string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	return func(r *DanglingIndicesDeleteDanglingIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}

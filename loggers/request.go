package loggers

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// RecordThridParty ...
func RecordThridParty(ctx context.Context, req *http.Request, start time.Time, service string, status int, body io.Reader, response []byte) context.Context {
	var (
		payload string
	)
	t := time.Since(start)
	if body != nil {

		bd, _ := req.GetBody()
		reqBody, err := ioutil.ReadAll(bd)
		if err != nil {
			payload = ""
		} else {

			payload = string(reqBody)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		third := ThirdParty{}

		third.Service = service
		third.URL = req.Host + req.URL.Path
		third.Response = string(response)
		third.StatusCode = status
		third.RequestHeader = DumpRequest(req)
		third.RequestBody = payload
		third.ExecTime = t.Seconds()
		third.RequestID = v.RequestID
		v.ThirdParty = append(v.ThirdParty, third)
		ctx = context.WithValue(ctx, logKey, v)

		return ctx
	}
	return ctx
}

// RecordThridPartyFailed ...
func RecordThridPartyFailed(ctx context.Context, req *http.Request, start time.Time, service string, status int, body io.Reader, messages string) context.Context {
	var (
		url     = req.Host + req.URL.Path
		payload string
	)

	t := time.Since(start)
	if req == nil {
		url = ""
	}

	if body != nil {
		bd, _ := req.GetBody()
		reqBody, err := ioutil.ReadAll(bd)
		if err != nil {
			payload = ""
		} else {
			payload = string(reqBody)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		third := ThirdParty{}

		third.Service = service
		third.URL = url
		third.Response = messages
		third.StatusCode = status
		third.RequestHeader = DumpRequest(req)
		third.RequestBody = payload
		third.ExecTime = t.Seconds()

		third.RequestID = v.RequestID
		v.ThirdParty = append(v.ThirdParty, third)

		ctx = context.WithValue(ctx, logKey, v)

		return ctx
	}

	return ctx
}

package loggers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// StartRecord for initialize context first time
func StartRecord(req *http.Request) context.Context {
	ctx := req.Context()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	start := time.Now().In(loc)

	v := new(Data)
	v.RequestID = uuid.New().String()
	v.Service = filepath.Base(os.Args[0])
	v.Host = req.Host
	v.Endpoint = req.URL.Path
	v.TimeStart = start
	v.RequestMethod = req.Method
	v.RequestHeader = DumpRequest(req)

	ctx = context.WithValue(ctx, logKey, v)
	return ctx
}

// DumpRequest is for get all data request header
func DumpRequest(req *http.Request) string {
	header, err := httputil.DumpRequest(req, true)
	if err != nil {
		return "cannot dump request"
	}

	trim := bytes.ReplaceAll(header, []byte("\r\n"), []byte("   "))
	return string(trim)
}

//GetName is return binary name
func GetName() string {
	return filepath.Base(os.Args[0])
}

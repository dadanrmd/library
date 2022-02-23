package loggers

import "time"

const (
	logKey = "data"
)

// Data is data standard output
type Data struct {
	RequestID     string       `json:"RequestID"`
	TimeStart     time.Time    `json:"TimeStart"`
	Service       string       `json:"Service"`
	Host          string       `json:"Host"`
	Endpoint      string       `json:"Endpoint"`
	RequestMethod string       `json:"RequestMethod"`
	RequestHeader string       `json:"RequestHeader"`
	StatusCode    int          `json:"StatusCode"`
	Response      string       `json:"Response"`
	ExecTime      float64      `json:"ExecutionTime"`
	Messages      []string     `json:"Messages"`
	ThirdParty    []ThirdParty `json:"3rdParty"`
}

// ThirdParty is data logging for any request to outside
type ThirdParty struct {
	RequestID     string  `json:"request_id,omitempty"`
	Service       string  `json:"service"`
	URL           string  `json:"url"`
	RequestHeader string  `json:"request_header"`
	RequestBody   string  `json:"request_body"`
	Response      string  `json:"response"`
	StatusCode    int     `json:"status_code"`
	ExecTime      float64 `json:"exec_time"`
}

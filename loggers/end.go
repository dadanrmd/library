package loggers

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

// EndRecord for initialize context first time
func EndRecord(ctx context.Context, response string, statuscode int) {
	var level string

	v, ok := ctx.Value(logKey).(*Data)
	if ok {
		t := time.Since(v.TimeStart)

		if statuscode >= 200 && statuscode < 400 {
			level = "INFO"
		} else if statuscode >= 400 && statuscode < 500 {
			level = "WARN"
		} else {
			level = "ERROR"
		}

		v.StatusCode = statuscode
		v.Response = response
		v.ExecTime = t.Seconds()

		if statuscode == 0 {
			v.StatusCode = 200
		}

		// Getprometheus().MetricRecord(strconv.Itoa(v.StatusCode), v.RequestMethod, v.Endpoint, GetName(), t)

		Output(v, level)
	}
}

// UTCFormatter ...
type UTCFormatter struct {
	logrus.Formatter
}

// Format ...
func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	e.Time = e.Time.In(loc)
	return u.Formatter.Format(e)
}

// Output for output to terminal
func Output(out *Data, level string) {
	logrus.SetFormatter(UTCFormatter{&logrus.JSONFormatter{}})

	if level == "ERROR" {
		logrus.WithField("data", out).Error("apps")
	} else if level == "INFO" {
		logrus.WithField("data", out).Info("apps")
	} else if level == "WARN" {
		logrus.WithField("data", out).Warn("apps")
	}
}

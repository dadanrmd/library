package utils

import "github.com/sirupsen/logrus"

//dataUser is main from device
type dataUser struct {
	SubscriberID string `json:"subscriber_id"`
	Ppsoccd      string `json:"ppsoccd"`
	Msisdn       string `json:"msisdn"`
	DeviceType   string `json:"device_type"`
	DeviceOs     string `json:"device_os"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
}

//Device for logging device
type Device struct {
	Msisdn     string `json:"msisdn"`
	DeviceType string `json:"device_type"`
	DeviceOS   string `json:"device_os"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
}

//CustomFormatter is struct
type CustomFormatter struct {
	logrus.Formatter
}

//Loggers is field need determine
type Loggers struct {
	ReqID     string
	Method    string
	Endpoints string
	Device
}

// SetCache ..
type SetCache struct {
	Set bool
	Env string
}

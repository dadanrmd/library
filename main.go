package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/dadanrmd/library/loggers"
	"github.com/dadanrmd/library/utils"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func getAllEvents(c *gin.Context) {
	ctx := loggers.StartRecord(c.Request)
	ctx = loggers.Logf(ctx, "Error function RequestHandler() ")
	curl(ctx)
	curl(ctx)

	utils.BasicResponse(ctx, c.Writer, true, http.StatusOK, events)
}
func curl(ctx context.Context) (context.Context, []byte, int, error) {
	ctx = loggers.Logf(ctx, "hit dati ii ")
	var rs utils.Request
	urls, timer := "/data/api/v1/dati/3203112503770066", "30"
	baseurl := "http://35.219.77.34:31000"

	rs.Service = "soa_subcriber_atomic_info"
	rs.Method = http.MethodGet
	rs.URL = baseurl + urls
	rs.Header = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	t, _ := strconv.Atoi(timer)
	// rs.Payload = bytes.NewBuffer(payload)
	rs.Timeout = time.Duration(t) * time.Second
	return rs.DoRequest(ctx)
}

func main() {
	router := gin.Default()
	router.GET("/events", getAllEvents)

	router.Run("localhost:8080")
}

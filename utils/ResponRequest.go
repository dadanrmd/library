package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dadanrmd/library/loggers"

	_ "github.com/joho/godotenv/autoload" //buat jaga2
	"github.com/spf13/cast"
)

//SecureResponse is struct
type SecureResponse struct {
	Status       bool        `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
	Title        string      `json:"title"`
	Data         interface{} `json:"data"`
}

//recordCode is func record status code
func generateResponse(ctx context.Context, w http.ResponseWriter, code int, res *SecureResponse) {
	response, err := json.Marshal(res)
	if err != nil {
		loggers.Logf(ctx, "Error marshal on line 86 ResponseRequest.go => %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, string(response))
}

/*BasicResponse parts
 * @updated: Wednesday, February 5th, 2020.
 * --
 * @param	w    	io.Writer
 * @param	mixed	msg
 * @param	mixed	code
 * @param	data 	string
 * @return	void
 */
func BasicResponse(ctx context.Context, w http.ResponseWriter, status bool, code int, rs interface{}, message string) {
	var (
		response SecureResponse
		result   string
	)

	response.Status = status
	response.ResponseCode = code

	if status {
		response.Message = message
		response.Data = rs
		input, _ := JSONMarshal(rs)
		result = cast.ToString(input)
	} else {
		response.Title = rs.(string)
		response.Message = message
		result = response.Title
	}
	// data.Response = result
	loggers.EndRecord(ctx, result, code)

	generateResponse(ctx, w, code, &response)
}

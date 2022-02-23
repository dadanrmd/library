package utils

import (
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload" // buat jaga2

	"github.com/sony/gobreaker"
)

var (
	servbreaker *gobreaker.CircuitBreaker
	transports  = http.DefaultTransport.(*http.Transport).Clone()
	defclient   = http.DefaultClient
)

/*LoadConfig parts
 * @updated: Tuesday, December 10th, 2019.
 * --
 * @return	void
 */

func loadClient() {
	transports.TLSClientConfig = &tls.Config{InsecureSkipVerify: checkTLS()}
	defclient.Transport = transports
}

func checkTLS() bool {
	skip := strings.ToLower(os.Getenv("INSECURE_SKIP_VERIFY"))
	if skip == "false" {
		return false
	}
	return true
}

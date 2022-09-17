package learngoweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	requestHeader := request.Header

	fmt.Println("REQUEST HEADER:")
	for key, value := range requestHeader {
		fmt.Fprint(writer, key, ":", value)
		fmt.Println(key, ":", value)
	}

	writer.Header().Add("pesan", "Ini response header dari server")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/", nil)

	// menambahkan request header
	request.Header.Add("Browser", "Opera Browser")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	responseHeader := recorder.Header()

	fmt.Println("RESPONSE HEADER:")
	fmt.Println(responseHeader)
}

/*
=== RUN   TestRequestHeader
REQUEST HEADER:
Browser : [Opera Browser]
RESPONSE HEADER:
map[Content-Type:[text/plain; charset=utf-8] Pesan:[Ini response header dari server]]
--- PASS: TestRequestHeader (0.00s)
PASS
ok      learn-go-web    0.491s
*/

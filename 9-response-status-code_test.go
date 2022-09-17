package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseStatusCodeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "ini adalah materi response status code")
}

func GetNameHandler(writer http.ResponseWriter, request *http.Request) {
	fullName := request.URL.Query().Get("fullName")

	if fullName == "" {
		writer.WriteHeader(http.StatusBadRequest) // bad request: 400
		fmt.Fprint(writer, "Tidak ada nama yang dikirim")
	} else {
		writer.WriteHeader(http.StatusOK) // Status OK : 200
		fmt.Fprint(writer, "Hello "+fullName)
	}
}

func TestResponseStatusCodeHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	ResponseStatusCodeHandler(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	fmt.Println(string(body))

	fmt.Println("Response Status Code: ", statusCode)
}

/*
=== RUN   TestResponseStatusCodeHandler
ini adalah materi response status code 200
--- PASS: TestResponseStatusCodeHandler (0.00s)
PASS
ok      learn-go-web    0.449s
*/

func TestGetNameValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/?fullName=RasyidevPro", nil)
	recorder := httptest.NewRecorder()

	GetNameHandler(recorder, request)

	response := recorder.Result()
	statusCode := response.StatusCode
	status := response.Status
	body, _ := io.ReadAll(response.Body)
	fmt.Println(statusCode)
	fmt.Println(status)
	fmt.Println(string(body))
}

/*
=== RUN   TestGetNameValid
200
200 OK
Hello RasyidevPro
--- PASS: TestGetNameValid (0.00s)
PASS
ok      learn-go-web    0.530s
*/

func TestGetNameInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090/", nil)
	recorder := httptest.NewRecorder()

	GetNameHandler(recorder, request)

	response := recorder.Result()
	statusCode := response.StatusCode
	status := response.Status
	body, _ := io.ReadAll(response.Body)
	fmt.Println(statusCode)
	fmt.Println(status)
	fmt.Println(string(body))
}

/*
=== RUN   TestGetNameInvalid
400
400 Bad Request
Tidak ada nama yang dikirim
--- PASS: TestGetNameInvalid (0.00s)
PASS
ok      learn-go-web    0.483s
*/

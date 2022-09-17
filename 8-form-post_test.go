package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func GetPostFormHandler(writer http.ResponseWriter, request *http.Request) {
	// parsing form dulu sebelum ambil value
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	fmt.Println("REQUEST BODY")
	fmt.Println(request.PostForm)
	firstName := request.Form.Get("firstName")
	lastName := request.Form.Get("lastName")

	fmt.Fprintln(writer, "firstName:", firstName)
	fmt.Fprintln(writer, "lastName:", lastName)

}

func TestGetPostForm(t *testing.T) {
	// request body: tempat meletakkan value post
	requestBody := strings.NewReader("firstName=Rasyidev&lastName=Pro")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:9090/", requestBody)

	// Wajib tambahkan application/x-www-form-urlencoded untuk kirim form di request body
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	GetPostFormHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("RESPONSE HEADER")
	fmt.Println(string(body))

}

/*
$ go test -v -run TestGetPostForm
=== RUN   TestGetPostForm
REQUEST BODY
map[firstName:[Rasyidev] lastName:[Pro]]
RESPONSE HEADER
firstName: Rasyidev
lastName: Pro

--- PASS: TestGetPostForm (0.00s)
PASS
ok      learn-go-web    0.504s
*/

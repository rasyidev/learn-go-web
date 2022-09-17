package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func SetCookieHandler(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-RASYIDEV-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Berhasil menambahkan cookie")
}

func GetCookieHandler(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-RASYIDEV-Name")
	if err != nil {
		panic(err)
	}

	if cookie.Name == "" {
		fmt.Fprint(writer, "Cookie Kosong")
	} else {
		fmt.Fprintln(writer, "Nama Cookie\t:"+cookie.Name)
		fmt.Fprintln(writer, "Nilai Cookie\t:"+cookie.Value)
		fmt.Fprintln(writer, "Lokasi Cookie\t:"+cookie.Path)
	}

}

func DeleteCookieHandler(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:    "X-RASYIDEV-Name",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(writer, &cookie)
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookieHandler)
	mux.HandleFunc("/get-cookie", GetCookieHandler)
	mux.HandleFunc("/delete-cookie", DeleteCookieHandler)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Test di Browser sambil pantau Storage>Cookies (CTRL+SHIFT+I):
localhost:9090/set-cookie?name=RasyidevPro
localhost:9090/get-cookie
localhost:9090/delete-cookie

*/

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:9090/set-cookie?name=RasyidevPro", nil)
	recorder := httptest.NewRecorder()

	SetCookieHandler(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	cookie := response.Cookies()[0]
	assert.Equal(t, "X-RASYIDEV-Name", cookie.Name)
	assert.Equal(t, "RasyidevPro", cookie.Value)
	assert.Equal(t, "/", cookie.Path)
	assert.Equal(t, "Berhasil menambahkan cookie", string(body))
}

/*
=== RUN   TestSetCookie
--- PASS: TestSetCookie (0.00s)
PASS
ok      learn-go-web    0.501s
*/

package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	templateName := "SIMPLE"

	// membuat template Cara 1
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	// membuat template Cara 2
	t := template.Must(template.New(templateName).Parse(templateText))
	// Ekeskusi template templateName. Bisa dimasukkan banyak template
	t.ExecuteTemplate(writer, templateName, "<h1>Hello HTML Template<h1>")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

/*
$ go test -v -run TestSimpleHTML
=== RUN   TestSimpleHTML
<html><body><h1>Hello HTML Template<h1></body></html>
--- PASS: TestSimpleHTML (0.00s)
PASS
ok      learn-go-web    0.309s
*/

package learngoweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates2 embed.FS

// dibuat di global variabel supaya hanya diakses sekali aja
var myTemplates = template.Must(template.ParseFS(templates2, "templates/*.gohtml"))

func TemplateChaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateChaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateCaching
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Hello Template Caching</title>
</head>
<body>
  <h1>Hello Template Caching</h1>
</body>
</html>
--- PASS: TestTemplateCaching (0.00s)
PASS
ok      learn-go-web    0.301s
*/

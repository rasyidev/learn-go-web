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
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template Embed")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

/*
=== RUN   TestTemplateEmbed
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Hello HTML Template Embed</title>
</head>
<body>
  <h1>Hello HTML Template Embed</h1>
</body>
</html>
--- PASS: TestTemplateEmbed (0.00s)
PASS
ok      learn-go-web    0.257s
*/

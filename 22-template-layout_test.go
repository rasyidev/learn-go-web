package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"templates/header.gohtml",
		"templates/footer.gohtml",
		"templates/layout.gohtml",
	))

	t.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Name":  "Rasyidev Pro",
		"Title": "Layout Template",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateLayout
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title><no value></title>
</head>
<body>
<h1>Hello Rasyidev Pro</h1>
</body>
</html>

--- PASS: TestTemplateLayout (0.00s)
PASS
ok      learn-go-web    0.357s
*/

package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateIfNameExist(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/template-if.gohtml"))
	t.ExecuteTemplate(writer, "template-if.gohtml", map[string]string{
		"Name": "Rasyidev Pro",
	})
}

func TestTemplateIfNameExist(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateIfNameExist(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateIf
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template If</title>
	</head>
	<body>

	<h1>Hello Rasyidev Pro</h1>

	</body>
</html>
--- PASS: TestTemplateIf (0.00s)
PASS
ok      learn-go-web    0.317s
*/

func TemplateIfNameNotExist(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/template-if.gohtml"))
	t.ExecuteTemplate(writer, "template-if.gohtml", nil)
}

func TestTemplateIfNameNotExist(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateIfNameNotExist(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateIfNameNotExist
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template If</title>
</head>
<body>

    <h1>Hello Stranger</h1>

</body>
</html>
--- PASS: TestTemplateIfNameNotExist (0.00s)
PASS
ok      learn-go-web    (cached)
*/

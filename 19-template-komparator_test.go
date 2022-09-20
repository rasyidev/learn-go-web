package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Page struct {
	Title      string
	FinalScore int
}

func TemplateKomparator(writer http.ResponseWriter, request *http.Request) {
	page := Page{
		Title:      "Template Komparator",
		FinalScore: 89,
	}

	t := template.Must(template.ParseFiles("templates/template-komparator.gohtml"))
	t.ExecuteTemplate(writer, "template-komparator.gohtml", page)
}

func TestTemplateKomparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateKomparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateKomparator
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template Komparator</title>
</head>
<body>

    <h1>Excelent</h1>

</body>
</html>
--- PASS: TestTemplateKomparator (0.00s)
PASS
ok      learn-go-web    0.301s
*/

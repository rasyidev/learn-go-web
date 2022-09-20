package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateNestedValue(writer http.ResponseWriter, request *http.Request) {
	activities := map[string]interface{}{
		"Title": "Template Range",
		"Name":  "Rasyidev Pro",
		"Activities": []interface{}{
			"Eat", "Code", "Pray", "Repeat",
		},
		"Alamat": map[string]interface{}{
			"Kota":     "Baturaja",
			"Provinsi": "Sumatera Selatan",
		},
	}

	t := template.Must(template.ParseFiles("templates/template-nested-value.gohtml"))
	t.ExecuteTemplate(writer, "template-nested-value.gohtml", activities)
}

func TestTemplateNestedValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateNestedValue(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateNestedValue
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template Range</title>
</head>
<body>
  <h1>Rasyidev Pro</h1>
  <h2>Activities</h2>

    <p>0. Eat</p>

    <p>1. Code</p>

    <p>2. Pray</p>

    <p>3. Repeat</p>



    <p>Kota: Baturaja</p>
    <p>Provinsi: Sumatera Selatan</p>

</body>
</html>
--- PASS: TestTemplateNestedValue (0.00s)
PASS
ok      learn-go-web    0.295s
*/

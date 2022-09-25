package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func XSSHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "xss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTMLEscapeString("<p>Ini adalah body</p> <script>alert('Hayoloo, kena hack!')</script>"),
	})
}

func TestXSSAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:9090", nil)
	recorder := httptest.NewRecorder()

	XSSHandler(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestXSSAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: http.HandlerFunc(XSSHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

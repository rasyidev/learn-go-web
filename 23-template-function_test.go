package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + " my name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("MyFunction").Parse(`{{.SayHello "Rasyidev Pro"}}`))
	t.ExecuteTemplate(writer, "MyFunction", MyPage{
		Name: "Habib",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*
=== RUN   TestTemplateFunction
Hello Rasyidev Pro my name is Habib
--- PASS: TestTemplateFunction (0.00s)
PASS
ok      learn-go-web    0.246s
*/

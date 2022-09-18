package learngoweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") == "" {
		http.ServeFile(writer, request, "resources/noname.html")
	} else {
		http.ServeFile(writer, request, "resources/ok.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}

/*
Coba:
localhost:9090
localhost:9090?name=RasyidevPro
*/

//go:embed resources/noname.html
var resourceNoName string

//go:embed resources/ok.html
var resourceOk string

func EmbedServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") == "" {
		fmt.Fprint(writer, resourceNoName)
	} else {
		fmt.Fprint(writer, resourceOk)
	}
	// embed
}

func TestEmbedServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: http.HandlerFunc(EmbedServeFile),
	}

	server.ListenAndServe()
}

package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTarget(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Ini adalah target redirect")
}

func RedirectSource(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-target", http.StatusTemporaryRedirect)
}

func TestRedirectServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-source", RedirectSource)
	mux.HandleFunc("/redirect-target", RedirectTarget)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

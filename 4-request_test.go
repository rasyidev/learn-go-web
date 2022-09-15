package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request URI\t:", r.RequestURI)
		fmt.Fprintln(w, "Request Method\t:", r.Method)
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// Mengirim binary ke writer
		fmt.Fprint(w, "Hello Rasyidev")
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

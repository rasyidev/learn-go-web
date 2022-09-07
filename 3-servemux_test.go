package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Home Page of Rasyidev website")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page of Rasyidev website")
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

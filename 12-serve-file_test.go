package learngoweb

import (
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

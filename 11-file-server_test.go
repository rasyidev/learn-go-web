package learngoweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	server.ListenAndServe()

}

//go:embed resources
var resources embed.FS

func TestEmbedFileServer(t *testing.T) {
	// menghilangkan path resources. Sehingga aksesnya hanya: /static/
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}

	// konversi golang embed menjadi http File Server
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	server.ListenAndServe()
}

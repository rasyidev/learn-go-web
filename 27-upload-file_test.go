package learngoweb

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "file-upload.gohtml", nil)
}

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err.Error())
	}

	// membuat fileDestination
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err.Error())
	}

	// simpan file uploadnya ke folder tujuan
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err.Error())
	}

	filename := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "file-upload-success.gohtml", map[string]interface{}{
		"FileName": filename,
		"Title":    "Success Upload Form Multipart",
		"File":     "/static/" + fileHeader.Filename,
	})

}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", UploadHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

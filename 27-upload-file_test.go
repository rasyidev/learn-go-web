package learngoweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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

//go:embed resources/rasyidev.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	// writer multipart
	writer := multipart.NewWriter(body)
	// field name
	writer.WriteField("name", "Rasyidev Pro")
	// field file
	file, _ := writer.CreateFormFile("file", "rasyidev-pro.png")
	// write dari file menggunakan embed
	file.Write(uploadFileTest)
	// tutup writer filenya
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "localhost:9090/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadHandler(recorder, request)

	responseBody, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(responseBody))
}

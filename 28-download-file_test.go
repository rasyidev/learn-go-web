package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFileInline(writer http.ResponseWriter, request *http.Request) {
	// menangkap fileName dari url query
	fileName := request.URL.Query().Get("file")
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "BAD REQUEST FROM INLINE HANDLER")
		return
	}
	http.ServeFile(writer, request, "./resources/"+fileName)

}

func DownloadFileAttachment(writer http.ResponseWriter, request *http.Request) {
	// menangkap fileName dari url query
	fileName := request.URL.Query().Get("file")
	fmt.Println(fileName)
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "BAD REQUEST FROM ATTACHMENT HANDLER")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
	http.ServeFile(writer, request, "./resources/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/view", DownloadFileInline)
	mux.HandleFunc("/download", DownloadFileAttachment)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

/*
Coba di browser:
http://localhost:9090/view?view=rasyidev.png
http://localhost:9090/view?file=rasyidev.png
*/

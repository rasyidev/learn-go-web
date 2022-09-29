package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

// Middleware untuk logging setiap ada request
type LogMiddleware struct {
	Handler http.Handler
}

// Middleware untuk proses error
type ErrorHandler struct {
	Handler http.Handler
}

// Function milik Struct LogMiddleware untuk logging setiap ada request
// di handler manapun selama melewati LogMiddleware
func (logMiddleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	// forward aja ke Handler.ServeHTTP()
	logMiddleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		// untuk menangkap error
		err := recover()
		if err != nil {
			fmt.Println("Error nih")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()

	// forward ke Handler.ServeHTTP()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Home Page")
	})

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("Yah, error. Hehehe")
	})

	// handler: mux
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// handler: logMiddleware
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	// handler: logMiddleware
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Server -> errorHandler -> logHandler -> mux

/*
Before Execute Handler
Handler Executed
After Execute Handler
*/

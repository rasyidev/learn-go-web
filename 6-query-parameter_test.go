package learngoweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	name := queries["name"]
	married := queries.Get("married")

	if len(name) > 1 {
		fmt.Println("Multiple Value")
		fmt.Println(len(name))
	}

	if married != "" && len(name) == 1 {
		m, err := strconv.ParseBool(married)
		if err != nil {
			panic(err)
		}

		if m == false {
			fmt.Fprintf(w, "Hallo mas %v", name)
			// fmt.Printf("Hallo pak %v\n", name)
		} else {
			fmt.Fprintf(w, "Halo pak %v", name)
			// fmt.Printf("Hallo mas %v\n", name)
		}
	}
}

func TestSearchHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/search?name=rasyidev&married=true", nil)
	response := httptest.NewRecorder()

	SearchHandler(response, request)

	fmt.Println(response.Body)

}

func TestSearchHandlerMultipleValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/search?name=rasyidev&name=pro&married=true", nil)
	response := httptest.NewRecorder()

	SearchHandler(response, request)

	fmt.Println("Response Body")
	fmt.Println(response.Body)

}

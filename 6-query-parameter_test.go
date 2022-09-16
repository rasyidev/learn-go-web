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
	name := queries.Get("name")
	married := queries.Get("married")

	if married != "" && name != "" {
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

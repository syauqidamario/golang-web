package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(400) // Bad Request
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(200) // OK
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestRespons(t *testing.T){
	req := httptest.NewRequest("GET", "http://localhost/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, req)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == " "{
		fmt.Fprint(writer, "Hello")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter (t *testing.T){

	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Syauqi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleParameter(writer http.ResponseWriter, request *http.Request){
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "%s %s", firstName, lastName)
}

func TestMultipleParameter (t*testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Eko&last_name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultiParameterVal(writer http.ResponseWriter, request *http.Request){
	var query url.Values = request.URL.Query()
	var names []string = query["name"]
	fmt.Fprintln(writer, strings.Join(names, ", "))

}

func TestMultiParameterVal (t*testing.T){

	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultiParameterVal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
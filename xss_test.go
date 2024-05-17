package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEsc(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		"Body":  "<p>Selamat belajar Golang web</p>",
	})
}

func TestTemplateAutoEsc(t *testing.T){

	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEsc),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TemplateAutoEscDisabled(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "post_disabled.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		"Body":  "<p>Selamat belajar Golang web</p>",
	})
}

func TestTemplateAutoEscDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateXSS(writer http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post_xss.gohtml", map[string]interface{}{
		"Title": "Golang Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
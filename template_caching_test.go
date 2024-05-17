package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Embed the templates directory
//go:embed templates/*.gohtml
var templateFS embed.FS

// Parse the templates once at program startup
var myTemplates = template.Must(template.ParseFS(templateFS, "templates/*.gohtml"))

// Handler function that executes the template
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	err := myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

// Test function to test the template caching
func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	// Read the response body
	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Print the response body
	fmt.Println(string(body))

	// Optional: Add assertions to verify the response content
	expected := `<p>Hello Template Caching</p>`
	if string(body) != expected {
		t.Errorf("Expected %q but got %q", expected, string(body))
	}
}

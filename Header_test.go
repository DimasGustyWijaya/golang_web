package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeaderWeb(w http.ResponseWriter, r *http.Request) {
	ContentType := r.Header.Get("content-type")
	fmt.Fprintln(w, ContentType)
}
func TestHeaderWeb(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/JSON")
	recorder := httptest.NewRecorder()

	HeaderWeb(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s", body)
}

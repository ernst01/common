package response_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ernst01/common/pkg/response"
)

func TestSendSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response.SendSuccess(w, http.StatusOK, "something_in_writing")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if http.StatusOK != resp.StatusCode {
		t.Errorf("TestSendSuccess() StatusCode error: expected %d received %d", http.StatusOK, resp.StatusCode)
	}
	if "something_in_writing" != string(body) {
		t.Errorf("TestSendSuccess() body error: expected %s received %s", "something_in_writing", string(body))
	}
}

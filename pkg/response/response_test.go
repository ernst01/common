package response_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/ernst01/common/pkg/response"
)

type User struct {
	Id uint64 `json:"user_id"`
}

var testUser = &User{
	Id: 444,
}

func TestSendSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response.SendJSONSuccess(w, http.StatusOK, testUser)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if http.StatusOK != resp.StatusCode {
		t.Errorf("TestSendSuccess() StatusCode error: expected %d received %d", http.StatusOK, resp.StatusCode)
	}

	rUser := &User{}
	err := json.Unmarshal(body, rUser)
	if err != nil {
		t.Error("json Unmarshalling error")
	}

	if !reflect.DeepEqual(testUser, rUser) {
		t.Errorf("TestSendSuccess() expected user %+v doesn't match received user %+v", testUser, rUser)
	}
}

func TestSendError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		response.SendJSONError(w, http.StatusNotFound, "http://example.com/help_url", "Unable to find a record matching ID: %d", 12)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if http.StatusNotFound != resp.StatusCode {
		t.Errorf("TestSendSuccess() StatusCode error: expected %d received %d", http.StatusNotFound, resp.StatusCode)
	}
	if !json.Valid(body) {
		t.Errorf("TestSendSuccess() body error: Invalid JSON")
	}
}

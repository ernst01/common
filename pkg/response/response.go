package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type appError struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	HelpURL string `json:"help_url,omitempty"`
}

// SendJSONSuccess sends a success response
func SendJSONSuccess(w http.ResponseWriter, httpStatus int, data interface{}) {
	SendJSONResponse(w, httpStatus, data)
}

// SendJSONError sends an error response
func SendJSONError(w http.ResponseWriter, httpStatus int, helpURL string, format string, a ...interface{}) {
	errorObj := &appError{
		Type:    "error",
		Status:  slugify(http.StatusText(httpStatus)),
		Message: fmt.Sprintf(format, a...),
		Code:    httpStatus,
		HelpURL: helpURL,
	}
	SendJSONResponse(w, httpStatus, errorObj)
}

// SendJSONResponse sends a response
func SendJSONResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	var buffer bytes.Buffer

	if err := json.NewEncoder(&buffer).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if _, err := io.Copy(w, &buffer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func slugify(message string) string {
	message = strings.ToLower(message)
	message = strings.Replace(message, " ", "_", -1)
	return message
}

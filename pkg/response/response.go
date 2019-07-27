package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// AppError represents an application error object
type AppError struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	HelpURL string `json:"help_url,omitempty"`
}

// SendSuccess sends a success response
func SendSuccess(w http.ResponseWriter, httpStatus int, data interface{}) {
	var buffer bytes.Buffer

	if data != nil {
		if err := json.NewEncoder(&buffer).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if _, err := io.Copy(w, &buffer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SendError sends an error
func SendError(w http.ResponseWriter, httpStatus int, format string, a ...interface{}) {
	errorObj := AppError{
		Type:    "error",
		Status:  slugify(http.StatusText(httpStatus)),
		Message: fmt.Sprintf(format, a...),
		Code:    httpStatus,
	}
	jsonError, err := json.Marshal(errorObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, string(jsonError), httpStatus)
}

func slugify(message string) string {
	message = strings.ToLower(message)
	message = strings.Replace(message, " ", "_", -1)
	return message
}

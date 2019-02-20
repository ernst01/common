package common

import (
	"net/http"
	"testing"
)

func TestSlugify(t *testing.T) {
	var slugifyTests = []struct {
		input          string
		expectedOutput string
	}{
		{"TeSt", "test"},
		{http.StatusText(http.StatusAccepted), "accepted"},
		{http.StatusText(http.StatusExpectationFailed), "expectation_failed"},
	}
	for _, tt := range slugifyTests {
		output := slugify(tt.input)
		if output != tt.expectedOutput {
			t.Errorf("slugify(%s): expected %s, actual %s", tt.input, tt.expectedOutput, output)
		}
	}
}

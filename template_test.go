package bat

import (
	"strings"
	"testing"
	"text/template"
)

func TestGetRawStringOfTemplate(t *testing.T) {
	tt := template.Must(template.ParseFiles("testdata/template.html"))
	raw, err := GetRawStringOfTemplate(tt, map[string]interface{}{"user": "lola"})
	raw = strings.Trim(raw, "\n ")
	expected := "Hello, lola!"
	if err != nil || raw != expected {
		t.Errorf("Expected %q, got %q with error %v", expected, raw, err)
	}
}

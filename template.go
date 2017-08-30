package bat

import (
	"bytes"
	"text/template"
)

// GetRawStringOfTemplate executes a template with given values and returns resulting string
func GetRawStringOfTemplate(t *template.Template, values map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, values); err != nil {
		return "", err
	}
	return buf.String(), nil
}

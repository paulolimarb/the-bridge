package util

import (
	"bytes"
	"text/template"
)

func ParseValues(tmpl string, data map[string]interface{}) (string, error) {

	buf := bytes.NewBuffer(nil)

	t := template.Must(template.New("query").Parse(tmpl))
	err := t.ExecuteTemplate(buf, "query", data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

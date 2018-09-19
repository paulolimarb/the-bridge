package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseValues(t *testing.T) {
	Convey("Deve-se enviar um template e retornar os valores substituidos", t, func() {
		var data = map[string]interface{}{"name": "Bob", "idade": 50}
		var tmpl = "{{.name}}-{{.idade}}"

		valueParsed, err := ParseValues(tmpl, data)

		So(valueParsed, ShouldEqual, "Bob-50")
		So(err, ShouldBeNil)

	})

}

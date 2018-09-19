package main

import (
	"fmt"
	"the-bridge/from"
	"the-bridge/models"
	"the-bridge/to"

	yaml "gopkg.in/yaml.v2"
)

/*var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
e: >
 select * from pessoa
`*/

var data = `
from: sqlserver
to: elasticsearch
config_Destination:
    index: pessoas
    document_Type: pessoa
    document_Id: "{{.email}}-{{.idade}}"
configuration:
    from_Conn: server=localhost;port=1433;user id=sa;password=sa;database=base;connection timeout=130
    to_Conn: http://localhost:9200/
    schedule: "*/3 * * * *"
query: >
 select * from pessoa
`

func check(err error) {

	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	t := new(models.Yaml)

	err := yaml.Unmarshal([]byte(data), &t)
	check(err)
	v, err := from.ExtractData(t.From, t.Configuration.From_Conn, t.Query)
	check(err)

	to.InsertElastic(t.Configuration.To_Conn, t.Config_Destination, v)

	fmt.Scanln()

}

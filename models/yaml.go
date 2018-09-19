package models

type Yaml struct {
	From               string                 `yaml:"from"` //sqlserver
	To                 string                 `yaml:"to"`   //elasticsearch, mongo
	Config_Destination map[string]interface{} `yaml:"config_Destination"`
	Configuration      struct {
		From_Conn string `yaml:"from_Conn"` //ConnectionString
		To_Conn   string `yaml:"to_Conn"`
		Schedule  string `yaml:"schedule"`
	}
	Query string
}

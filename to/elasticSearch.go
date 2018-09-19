package to

import (
	"encoding/json"
	"strings"
	"the-bridge/models"
	"the-bridge/util"
)

func InsertElastic(urlServer string, t map[string]interface{}, data map[string]interface{}) error {
	e := models.Elastic{}

	for k, v := range t {

		if k == "index" {
			e.Index = v.(string)
		}

		if k == "document_Type" {
			e.Document_type = v.(string)
		}

		if k == "document_Id" {
			ident, err := util.ParseValues(v.(string), data)

			if err != nil {
				return err
			}

			e.Document_id = ident
		}

	}

	b, _ := json.Marshal(data)
	url := mountURL(urlServer, e.Index, e.Document_type, e.Document_id)

	util.SendRequest(url, "application/json", strings.NewReader(string(b)))

	return nil

}

func mountURL(urlServer string, index string, docType string, docID string) string {

	return urlServer + index + "/" + docType + "/" + docID
}

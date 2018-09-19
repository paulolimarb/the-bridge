package from

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func ExtractData(sourceData string, connString string, query string) (map[string]interface{}, error) {

	db, err := sql.Open(sourceData, connString)

	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	columns := make([]interface{}, len(cols))
	vals := make([]interface{}, len(cols))

	for i, _ := range columns {
		vals[i] = &columns[i]
	}
	for rows.Next() {
		err = rows.Scan(vals...)

		if err != nil {
			return nil, err
		}

		for i, colName := range cols {
			val := vals[i].(*interface{})
			m[colName] = *val
		}

	}

	return m, nil
}

package convUtil

import "encoding/json"

func JsonToMap(str string, m map[string]interface{}) error {
	return json.Unmarshal([]byte(str), &m)
}

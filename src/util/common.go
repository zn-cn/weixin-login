package util

import (
	"encoding/json"
)

// JSONStructToMap convert struct to map by json comment
func JSONStructToMap(obj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(obj)
	var data map[string]interface{}
	json.Unmarshal(jsonBytes, &data)
	return data
}

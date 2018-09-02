package json

import "encoding/json"

func UnmarshallGenericJSON(iface interface{}) (map[string]interface{}, error) {
	var genericJSONStruct map[string]interface{}
	if iface != nil {
		raw, err := json.Marshal(iface)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(raw, &genericJSONStruct); err != nil {
			return nil, err
		}
	}
	return genericJSONStruct, nil
}

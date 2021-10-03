package cli

import json "github.com/json-iterator/go"

func JSONInput(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func JSONOutput(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

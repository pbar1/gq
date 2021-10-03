package cli

import "gopkg.in/yaml.v2"

func YAMLInput(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func YAMLOutput(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

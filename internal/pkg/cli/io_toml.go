package cli

import "github.com/pelletier/go-toml"

func TOMLInput(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func TOMLOutput(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

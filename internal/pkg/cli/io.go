package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl"
	json "github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

type (
	InputFuncMap  map[string]func([]byte, interface{}) error
	OutputFuncMap map[string]func(v interface{}) ([]byte, error)
)

var (
	inputFuncMap = InputFuncMap{
		"json": json.Unmarshal,
		"yaml": yaml.Unmarshal,
		"toml": toml.Unmarshal,
		"hcl":  hcl.Unmarshal,
		"hcl2": hcl2Unmarshal,
	}

	outputFuncMap = OutputFuncMap{
		"go-template": goTemplateMarshal,
		"jsonpath":    jsonpathMarshal,
		"json":        json.Marshal,
		"yaml":        yaml.Marshal,
		"toml":        toml.Marshal,
	}
)

// input unmarshals raw bytes into the specified format and returns an object.
func input(in []byte, format string) (interface{}, error) {
	unmarshal, found := inputFuncMap[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
	var v interface{}
	err := unmarshal(in, &v)
	return v, err
}

// output marshals an object into the specified format and returns raw bytes.
func output(obj interface{}, format string) ([]byte, error) {
	marshal, found := outputFuncMap[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
	return marshal(obj)
}

func (m *InputFuncMap) Options() string {
	keys := make(sort.StringSlice, 0)
	for k := range *m {
		keys = append(keys, k)
	}
	keys.Sort()
	return strings.Join(keys, "|")
}

func (m *OutputFuncMap) Options() string {
	keys := make(sort.StringSlice, 0)
	for k := range *m {
		keys = append(keys, k)
	}
	keys.Sort()
	return strings.Join(keys, "|")
}

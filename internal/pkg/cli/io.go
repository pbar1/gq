package cli

import (
	"fmt"
	"sort"
	"strings"
)

type (
	InputFuncMap  map[string]func([]byte, interface{}) error
	OutputFuncMap map[string]func(v interface{}) ([]byte, error)
)

var (
	InputFuncs = InputFuncMap{
		"json": JSONInput,
		"yaml": YAMLInput,
		"toml": TOMLInput,
		"hcl1": HCL1Input,
		"hcl2": HCL2Input,
		"hcl":  HCL2Input,
	}

	OutputFuncs = OutputFuncMap{
		"json":        JSONOutput,
		"yaml":        YAMLOutput,
		"toml":        TOMLOutput,
		"go-template": GoTemplateOutput,
		"jsonpath":    JSONPathOutput,
	}
)

// Input decodes raw bytes into the specified format and returns an object.
func Input(in []byte, format string) (interface{}, error) {
	input, found := InputFuncs[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
	var v interface{}
	err := input(in, &v)
	return v, err
}

// Output encodes an object into the specified format and returns raw bytes.
func Output(obj interface{}, format string) ([]byte, error) {
	output, found := OutputFuncs[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
	return output(obj)
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

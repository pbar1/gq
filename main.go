package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/hashicorp/hcl"
	"github.com/pelletier/go-toml"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var (
	version      = "unknown"
	printVersion = flag.Bool("version", false, "Print program version")
	inputFormat  = flag.StringP("input", "i", "json", "Input format. One of: json|yaml|toml|hcl")
	outputFormat = flag.StringP("output", "o", "go-template", "Output format. One of: go-template|json|yaml|toml")
	outTemplate  = flag.StringP("template", "t", `{{.}}`, "Go template string")
)

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var filename string
	var in []byte
	var err error

	if flag.Arg(0) == "" || flag.Arg(0) == "-" {
		filename = "stdin"
		in, err = ioutil.ReadAll(os.Stdin)
	} else {
		in, err = ioutil.ReadFile(filename)
	}
	check(err, "unable to read input from "+filename)

	intermediate, err := input(in, *inputFormat)
	check(err, "unable to input "+*inputFormat+" input from "+filename)

	out, err := output(intermediate, *outputFormat)
	check(err, "unable to render output")

	fmt.Println(string(out))
}

func output(v interface{}, format string) ([]byte, error) {
	var marshalFn func(v interface{}) ([]byte, error)
	switch strings.ToLower(*outputFormat) {
	case "go-template":
		tpl := template.Must(template.New("go-template").Funcs(sprig.TxtFuncMap()).Parse(*outTemplate))
		var buf bytes.Buffer
		if err := tpl.Execute(&buf, v); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	case "json":
		marshalFn = json.Marshal
	case "yaml":
		marshalFn = yaml.Marshal
	case "toml":
		marshalFn = toml.Marshal
	default:
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
	return marshalFn(v)
}

func input(in []byte, format string) (interface{}, error) {
	var unmarshalFn func(data []byte, v interface{}) error
	switch strings.ToLower(format) {
	case "json":
		unmarshalFn = json.Unmarshal
	case "yaml":
		unmarshalFn = yaml.Unmarshal
	case "toml":
		unmarshalFn = toml.Unmarshal
	case "hcl":
		unmarshalFn = hcl.Unmarshal
	default:
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
	var v interface{}
	err := unmarshalFn(in, &v)
	return v, err
}

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

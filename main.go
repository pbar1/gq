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
	printVersion = flag.BoolP("version", "v", false, "Print program version")
	inputFormat  = flag.StringP("input", "i", "json", "Input format. One of: json|yaml|toml|hcl")
	outputFormat = flag.StringP("output", "o", "go-template", "Output format. One of: go-template|json|yaml|toml")
	outTemplate  = flag.StringP("template", "t", `{{.}}`, "Go template string")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `Reads from stdin and writes to stdout. Can convert between input and output formats, including Go templates.

Examples:
  Feed Kubernetes YAML into gq and render it as a Go template
  $ kubectl get namespaces -o yaml | gq -i yaml -t '{{range .items}}{{.metadata.name}}{{printf "\n"}}{{end}}'

  Convert Terraform HCL into JSON (and feed that into jq for querying!)
  $ cat *.tf | gq -i hcl -o json | jq
`)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	filename := flag.Arg(0)
	if filename == "" || filename == "-" {
		filename = "stdin"
	}

	var in []byte
	var err error

	if filename == "stdin" {
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

func output(v interface{}, format string) ([]byte, error) {
	var marshalFn func(v interface{}) ([]byte, error)
	switch strings.ToLower(*outputFormat) {
	case "go-template":
		tpl, err := template.New("go-template").Funcs(sprig.TxtFuncMap()).Parse(*outTemplate)
		if err != nil {
			return nil, err
		}
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

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

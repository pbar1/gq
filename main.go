package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/hashicorp/hcl"
	json "github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const helpText = `Converts between input and output formats, including Go templates. Reads from stdin and writes to stdout.
Default Go template is "{{.}}"

Examples:
  Feed Kubernetes YAML into gq and render it as a Go template
  $ kubectl get namespaces -o yaml | gq -i yaml '{{range .items}}{{.metadata.name}}{{println}}{{end}}'

  You can omit the {{ }} if the template is simple enough. Sprig and more functions are in scope.
  $ kubectl get secret demo-tls -o json | gq '(index (index .data "tls.crt" | b64dec | x509decode) 0).NotBefore'

  Convert Terraform HCL into JSON (and feed that into jq for querying!)
  $ cat *.tf | gq -i hcl -o json | jq

Usage:
  gq [template string] [flags]

Flags:`

var (
	version      = "unknown"
	outputTpl    = `{{.}}`
	printVersion = flag.BoolP("version", "v", false, "Print program version")
	filename     = flag.StringP("file", "f", "-", "File to read input from. Defaults to stdin.")
	inputFmt     = flag.StringP("input", "i", "json", "Input format. One of: json|yaml|toml|hcl")
	outputFmt    = flag.StringP("output", "o", "go-template", "Output format. One of: go-template|json|yaml|toml")
	simple       = flag.BoolP("simple", "s", true, "Automatically wraps template in {{ }} if not already")
	inputFns     = map[string]func([]byte, interface{}) error{
		"json": json.Unmarshal,
		"yaml": yaml.Unmarshal,
		"toml": toml.Unmarshal,
		"hcl":  hcl.Unmarshal,
	}
	outputFns = map[string]func(v interface{}) ([]byte, error){
		"go-template": gotplMarshal,
		"json":        json.Marshal,
		"yaml":        yaml.Marshal,
		"toml":        toml.Marshal,
	}
	tplFns map[string]interface{}
)

func init() {
	flag.Usage = func() {
		_, err := fmt.Fprintln(os.Stderr, helpText)
		check(err, "unable to print help text to stderr")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if flag.Arg(0) != "" {
		outputTpl = flag.Arg(0)
	}

	var in []byte
	var err error
	f := *filename
	if f == "" || f == "-" {
		f = "stdin"
		in, err = ioutil.ReadAll(os.Stdin)
	} else {
		in, err = ioutil.ReadFile(f)
	}
	check(err, "unable to read input from "+f)

	intermediate, err := input(in, *inputFmt)
	check(err, "unable to parse input as "+*inputFmt)

	out, err := output(intermediate, *outputFmt)
	check(err, "unable to render output as "+*outputFmt)

	fmt.Println(string(out))
}

func input(in []byte, format string) (interface{}, error) {
	fn, found := inputFns[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported input format: %s", format)
	}
	var v interface{}
	err := fn(in, &v)
	return v, err
}

func output(v interface{}, format string) ([]byte, error) {
	fn, found := outputFns[strings.ToLower(format)]
	if !found {
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
	return fn(v)
}

func gotplMarshal(v interface{}) ([]byte, error) {
	tplFns = sprig.TxtFuncMap()
	tplFns["x509decode"] = x509decode
	tplFns["access"] = access
	tplFns["fnptr"] = fnptr

	// TODO: think about extracting this logic
	tplStr := outputTpl
	if *simple {
		if !strings.Contains(tplStr, "{{") {
			tplStr = "{{" + tplStr + "}}"
		}
		tplStr = strings.ReplaceAll(tplStr, "[", " | access ")
		tplStr = strings.ReplaceAll(tplStr, "]", " ")
	}

	tpl, err := template.New("base").Funcs(tplFns).Parse(tplStr)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

func x509decode(pemData string) ([]x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("pem decoded block empty")
	}
	crts, err := x509.ParseCertificates(block.Bytes)
	if err != nil {
		return nil, err
	}
	crtsCopy := make([]x509.Certificate, 0, 0)
	for _, c := range crts {
		crtsCopy = append(crtsCopy, *c)
	}
	return crtsCopy, nil
}

func access(property interface{}, obj interface{}) (interface{}, error) {
	if x, ok := obj.(map[interface{}]interface{}); ok {
		return x[property], nil
	}
	if x, ok := obj.([]interface{}); ok {
		i, iok := property.(int)
		if !iok {
			return nil, fmt.Errorf("%v not an int for list accessor", property)
		}
		return x[i], nil
	}
	return nil, nil
}

func fnptr(f string) interface{} {
	return tplFns[f]
}

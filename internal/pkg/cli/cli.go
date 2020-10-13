package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

const (
	helpText = `Converts between input and output formats, including Go templates. Reads from stdin and writes to stdout.
Default Go template is "{{.}}"

Examples:
  Feed Kubernetes YAML into gq and render it as a Go template
  $ kubectl get namespaces -o json | gq '{{range .items}}{{println .metadata.name}}{{end}}'

  You can omit the {{ }} if the Go template would be entirely contained within it. Sprig functions and more are available.
  $ kubectl get secret demo-tls -o json | gq '(index (index .data "tls.crt" | b64dec | x509Decode) 0).NotBefore'

  Convert Terraform HCL (v1) into JSON (and feed that into jq for querying!)
  $ cat *.tf | gq -i hcl -o json | jq

Usage:
  gq [template string] [flags]

Flags:`
)

var (
	argTemplate = defaultGoTemplate

	flagVersion = flag.BoolP("version", "v", false, "Print program version")
	flagFile    = flag.StringP("file", "f", "-", "File to read input from. Defaults to stdin.")
	flagInput   = flag.StringP("input", "i", "json", "Input format. One of: json|yaml|toml|hcl")
	flagOutput  = flag.StringP("output", "o", "go-template", "Output format. One of: go-template|jsonpath|json|yaml|toml")
	flagSimple  = flag.BoolP("simple", "s", true, "Automatically wraps Go template in {{ }} if not already")
)

func init() {
	flag.Usage = func() {
		_, err := fmt.Fprintln(os.Stderr, helpText)
		check(err, "unable to print help text to stderr")
		flag.PrintDefaults()
	}
}

// Execute is the entrypoint command
func Execute(version string) {
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	// join all positional args into one string delimited by spaces
	if len(flag.Args()) > 0 {
		argTemplate = strings.Join(flag.Args(), " ")
	}

	// read input from stdin by default, or a filename passed as a flag
	var in []byte
	var err error
	if *flagFile == "" || *flagFile == "-" {
		*flagFile = "stdin"
		in, err = ioutil.ReadAll(os.Stdin)
	} else {
		in, err = ioutil.ReadFile(*flagFile)
	}
	check(err, "unable to read input from "+*flagFile)

	intermediate, err := input(in, *flagInput)
	check(err, "unable to parse input as "+*flagInput)

	out, err := output(intermediate, *flagOutput)
	check(err, "unable to render output as "+*flagOutput)

	fmt.Println(string(out))
}

// check fatally exits with an error message if an error exists.
func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

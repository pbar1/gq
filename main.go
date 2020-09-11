package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/Masterminds/sprig"
	flag "github.com/spf13/pflag"
)

var (
	version = "unknown"
	// _            = flag.Bool("help", false, "Help for gq")
	printVersion = flag.Bool("version", false, "Print program version")
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

	var v interface{}
	err = json.Unmarshal(in, &v)
	check(err, "unable to unmarshal json input from "+filename)

	tpl := template.Must(template.New("base").Funcs(sprig.FuncMap()).Parse(*outTemplate))

	err = tpl.Execute(os.Stdout, v)
	check(err, "unable to render template")
}

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}

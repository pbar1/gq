# GQ

Convert between input and output formats, including Go templates.

### Usage
```
Reads from stdin and writes to stdout. Can convert between input and output formats, including Go templates.

Examples:
  Feed Kubernetes YAML into gq and render it as a Go template
  $ kubectl get namespaces -o yaml | gq -i yaml -t '{{range .items}}{{.metadata.name}}{{printf "\n"}}{{end}}'

  Convert Terraform HCL into JSON (and feed that into jq for querying!)
  cat *.tf | gq -i hcl -o json | jq

Usage of gq:
  -i, --input string      Input format. One of: json|yaml|toml|hcl (default "json")
  -o, --output string     Output format. One of: go-template|json|yaml|toml (default "go-template")
  -t, --template string   Go template string (default "{{.}}")
  -v, --version           Print program version
```

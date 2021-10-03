# Usage

`gq` supports a few core features, which are explained in-depth in the following
chapters.

By default (ie, given no arguments or flags), `gq` will read from standard input
until it encounters an end-of-file (EOF), expecting the data to be in JSON format.
It then injects this input data into a trivial Go template (noted below) and renders the
result to standard output. Of course, this behavior can be modified using the following
options.

## Positional Arguments

### `[template string]`

Either a [Go template][1], or a JSONPath query ([using `kubectl` syntax][2]).

**Default:**
- Go template: `{{ . }}`
- JSONPath: `{ . }`

## Flags

### `-f, --file`

File to read input from.

**Default:** `-` (stdin)

### `-i, --input`

Input format.

Options:
- **`json` (default)**
- `yaml`
- `toml`
- `hcl`
  - **Note:** At this time, only HCL 1 (Terraform <= 0.11) is supported. HCL 2 (Terraform >= 0.12) support is planned.

### `-o, --output`

Output format.

Options:
- **`go-template` (default)**
- `jsonpath`
- `json`
- `yaml`
- `toml`

### `-l, --lines`

Apply the operation to each line, rather than the whole input as one.

**Default:** `false`

### `-s, --simple`

Automatically wraps the template string with the necessary delimiters if they do not exist.
For Go templates this is `{{ }}`, while for JSONPath it is `{ }`.

**Default:** `true`

### `-v, --version`

Print program version information and exit. This flag overrides all others.

**Default:** `false`

<!-- Sources -->

[1]: https://golang.org/pkg/text/template/
[2]: https://kubernetes.io/docs/reference/kubectl/jsonpath/

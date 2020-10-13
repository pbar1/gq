package cli

import (
	"bytes"
	"strings"

	json "github.com/json-iterator/go"
	"k8s.io/client-go/util/jsonpath"
)

const defaultJSONPathTemplate = "{}"

// jsonpathMarshal renders a JSONPath template given in Kubernetes CLI format.
// More information: https://kubernetes.io/docs/reference/kubectl/jsonpath/
// TODO: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/kubectl/pkg/cmd/get/customcolumn.go
func jsonpathMarshal(v interface{}) ([]byte, error) {
	// since Go templates are the default, need to reset default template when opting for JSONPath
	t := argTemplate
	if t == defaultGoTemplate {
		t = defaultJSONPathTemplate
	}

	if *flagSimple {
		if !strings.Contains(t, "{") {
			t = "{" + t + "}"
		}
	}

	j := jsonpath.New("base")
	j.EnableJSONOutput(true)
	buf := new(bytes.Buffer)
	if err := j.Parse(t); err != nil {
		return nil, err
	}

	if err := j.Execute(buf, v); err != nil {
		return nil, err
	}
	// return buf.Bytes(), nil

	// TODO: there should be a better way to do this
	b := buf.Bytes()
	hack := new([]interface{})
	if err := json.Unmarshal(b, hack); err != nil {
		return nil, err
	}
	if len(*hack) < 1 {
		return nil, nil
	}
	return json.Marshal((*hack)[0])
}

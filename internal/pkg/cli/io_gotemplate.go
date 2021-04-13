package cli

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/dgrijalva/jwt-go"
)

const defaultGoTemplate = "{{.}}"

// goTemplateMarshal renders a Go template. Supports Sprig functions.
func goTemplateMarshal(v interface{}) ([]byte, error) {
	t := argTemplate

	if !strings.Contains(t, "{{") {
		if *flagRange {
			t = "{{ range . }}{{" + t + "}}{{ end }}"
		} else if *flagSimple {
			t = "{{" + t + "}}"
		}
	}

	tpl, err := template.New("base").Funcs(funcMap()).Parse(t)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// funcMap returns a map of functions that can be called within a Go template.
// Inspired by Helm: https://github.com/helm/helm/blob/master/pkg/engine/funcs.go
func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	extra := template.FuncMap{
		"fnptr":        fnptr,
		"base64decode": base64Decode,
		"x509decode":   x509Decode,
		"jwtdecode":    jwtDecode,
	}
	for k, v := range extra {
		f[k] = v
	}
	return f
}

// fnptr returns a reference to the function in the Go template function table
// by the given name.
func fnptr(f string) interface{} {
	return funcMap()[f]
}

// base64Decode decodes a standard base64-encoded string. For compatibility with
// kubectl: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/cli-runtime/pkg/printers/template.go
func base64Decode(v string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %v", err)
	}
	return string(data), nil
}

// x509Decode decodes and parses a PEM formatted X.509 certificate bundle.
func x509Decode(pemData string) ([]x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("pem decoded block empty")
	}
	crts, err := x509.ParseCertificates(block.Bytes)
	if err != nil {
		return nil, err
	}
	crtsCopy := make([]x509.Certificate, 0)
	for _, c := range crts {
		crtsCopy = append(crtsCopy, *c)
	}
	return crtsCopy, nil
}

// jwtDecode decodes a JWT token string, without validating the signature.
func jwtDecode(tokenData string) (*jwt.Token, error) {
	parser := new(jwt.Parser)
	token, _, err := parser.ParseUnverified(tokenData, jwt.MapClaims{})
	return token, err
}

package cli

import (
	json "github.com/json-iterator/go"
	"github.com/tmccombs/hcl2json/convert"
)

func HCL2Input(b []byte, v interface{}) error {
	jsonBytes, err := convert.Bytes(b, "tmp.hcl", convert.Options{Simplify: *flagHCL2Simplify})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBytes, &v); err != nil {
		return err
	}
	return nil
}

package cli

import (
	"github.com/hashicorp/hcl"
)

func HCL1Input(data []byte, v interface{}) error {
	return hcl.Unmarshal(data, v)
}

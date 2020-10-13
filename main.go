package main

import "github.com/pbar1/gq/internal/pkg/cli"

var version = "unknown"

func main() {
	cli.Execute(version)
}

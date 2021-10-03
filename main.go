package main

import "github.com/pbar1/gq/internal/pkg/cli"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	cli.Execute(version, commit, date, builtBy)
}

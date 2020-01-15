package main

import "github.com/SimonBaeumer/monorepo-operator/cmd"

var (
	version string
	build   string
	ref     string
)

func main() {
	cmd.Execute(version, build, ref)
}

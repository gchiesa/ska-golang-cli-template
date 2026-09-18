// Package main is the entry point for the {{ .appName }} CLI.
package main

import (
	"fmt"
	"os"

	"{{ regexReplaceAll "^https?://" .githubRepo "" }}/cmd"
)

var version = ""

func main() {
	err := cmd.Execute(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

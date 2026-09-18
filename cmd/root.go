// Package cmd provides the CLI commands for {{ .appName }}.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{ .appName }}",
		Short: "{{ .appDescription }}",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(newVersionCmd(version)) // version subcommand
	cmd.AddCommand(newExampleCmd())        // example subcommand

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	err := newRootCmd(version).Execute()
	if err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}

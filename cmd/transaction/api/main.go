package api

import "github.com/spf13/cobra"

func BuildApiCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "run-api",
		DisableFlagsInUseLine: true,
		RunE: func(cobra *cobra.Command, args []string) error {
			return nil
		},
	}
}

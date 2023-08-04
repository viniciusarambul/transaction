package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusarambul/transaction/cmd/transaction/api"
)

func main() {
	err := buildRootCommand()
	if err != nil {
		fmt.Print(err)
	}
}

func buildRootCommand() error {
	rootCmd := &cobra.Command{
		Use:     "transaction",
		Version: "1.0",
	}

	rootCmd.AddCommand(api.BuildApiCommand())

	return nil
}

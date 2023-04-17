package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type projectInfo struct {
	name    string
	version string
	binName string
}

var project = projectInfo{
	name:    "go-brutils",
	version: "v0.0.0-dev",
	binName: "brutils",
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   project.binName,
		Short: "Utilitário para negócios específicos do Brasil",
		Long:  "Aplicação para uso das funções providas pela biblioteca go-brutils.",
	}

	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.AddCommand(newCPFCommand())
	cmd.AddCommand(newCNPJCommand())

	return cmd
}

func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

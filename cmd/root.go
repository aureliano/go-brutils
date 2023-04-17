package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-brutils",
		Short: "Utilitário para negócios específicos do Brasil",
		Long:  "Aplicação para uso das funções providas pela biblioteca go-brutils.",
	}

	cmd.CompletionOptions.DisableDefaultCmd = true

	return cmd
}

func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCPFCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cpf",
		Short: "Manipular CPF",
		Long:  "Aciona funcionalidades para manipulação de CPF.",
		Example: fmt.Sprintf(`  %s cpf gerar   [flags]
  %s cpf validar [flags]`, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return cmd
}

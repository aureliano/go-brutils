package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cnpj",
		Short: "Manipular CNPJ",
		Long:  "Aciona funcionalidades para manipulação de CNPJ.",
		Example: fmt.Sprintf(`  %s cnpj gerar   [flags]
  %s cnpj validar [flags]`, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return cmd
}

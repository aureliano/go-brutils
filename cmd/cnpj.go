package cmd

import (
	"fmt"
	"log"

	"github.com/aureliano/go-brutils/rfb"
	"github.com/spf13/cobra"
)

func newCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cnpj",
		Short: "Manipular CNPJ",
		Long:  "Aciona funcionalidades para manipulação de CNPJ.",
		Example: fmt.Sprintf(`  %s cnpj gerar     [flags]
  %s cnpj completar [flags]
  %s cnpj validar`, project.binName, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(newGerarCNPJCommand())

	return cmd
}

func newGerarCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gerar",
		Short: "Gerar um número de CNPJ",
		Long:  "Aciona funcionalidade para geração de um número de CNPJ.",
		Example: fmt.Sprintf(`  # Gerar um número de CNPJ.
  %s cnpj gerar`, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			gerarNumeroCNPJ(cmd)
		},
	}

	cmd.Flags().BoolP("formatar", "f", false, "Formatar a saída com a máscara ##.###.###/####-##")

	return cmd
}

func gerarNumeroCNPJ(cmd *cobra.Command) {
	formatar, _ := cmd.Flags().GetBool("formatar")
	cnpj, err := rfb.GerarCNPJ()
	if err != nil {
		log.Fatalln(err)
	}

	if formatar {
		fmt.Println(cnpj.Formatado())
	} else {
		fmt.Println(cnpj.Desformatado())
	}
}

package cmd

import (
	"fmt"
	"log"

	"github.com/aureliano/go-brutils/rfb"
	"github.com/spf13/cobra"
)

func newCPFCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cpf",
		Short: "Manipular CPF",
		Long:  "Aciona funcionalidades para manipulação de CPF.",
		Example: fmt.Sprintf(`  %s cpf gerar     [flags]
  %s cpf completar [flags]
  %s cpf validar   [flags]`, project.binName, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(newGerarCPFCommand())

	return cmd
}

func newGerarCPFCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gerar",
		Short: "Gerar um número de CPF",
		Long:  "Aciona funcionalidade para geração de um número de CPF.",
		Example: fmt.Sprintf(`  # Gerar um CPF de um Estado/Região Fiscal tomado aleatoriamente.
		%s cpf gerar
	  
		# Gerar um CPF da Região Fiscal de Minas Gerais.
		%s cpf gerar -e MG`, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			gerarNumeroCPF(cmd)
		},
	}

	cmd.Flags().BoolP("formatar", "f", false, "Formatar a saída com a máscara ###.###.###-##")
	cmd.Flags().StringP("estado", "e", "", "Estado/Região Fiscal do CPF (ex.: MG, AM, ES, AC, MT etc.)")

	return cmd
}

func gerarNumeroCPF(cmd *cobra.Command) {
	formatar, _ := cmd.Flags().GetBool("formatar")
	estado, _ := cmd.Flags().GetString("estado")
	var cpf rfb.CPF
	var err error

	if estado == "" {
		cpf, err = rfb.GerarCPF()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		uf := rfb.NewEstado(estado)
		cpf, err = rfb.GerarCPFParaUF(uf)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if formatar {
		fmt.Println(cpf.Formatado())
	} else {
		fmt.Println(cpf.Desformatado())
	}
}

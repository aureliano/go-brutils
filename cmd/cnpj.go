package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aureliano/go-brutils/rfb"
	"github.com/spf13/cobra"
)

func newCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cnpj",
		Short: "Manipulação de CNPJ",
		Long:  "Aciona funcionalidades para manipulação de CNPJ.",
		Example: fmt.Sprintf(`  %s cnpj gerar     [flags]
  %s cnpj completar [flags]
  %s cnpj validar`, project.binName, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(newGerarCNPJCommand())
	cmd.AddCommand(newCompletarCNPJCommand())
	cmd.AddCommand(newValidarCNPJCommand())

	return cmd
}

func newGerarCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gerar",
		Short: "Gera um número de CNPJ",
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

func newCompletarCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completar",
		Short: "Completa um número de CNPJ",
		Long: "Completa um número de CNPJ, ora preenchendo com zeros a esquerda " +
			"ora preenchendo os dígitos verificadores.",
		Example: fmt.Sprintf(`  %s cpf completar 1981621`, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			completarNumeroCNPJ(cmd, args)
		},
	}

	cmd.Flags().BoolP("formatar", "f", false, "Formatar a saída com a máscara ##.###.###/####-##")

	return cmd
}

func newValidarCNPJCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validar",
		Short: "Valida um número de CNPJ",
		Long:  "Valida um número de CNPJ.",
		Example: fmt.Sprintf(`  # Validar um número de CNPJ sem máscara.
  %s cpf validar 00000198162197
  
  # Validar um número de CNPJ com máscara.
  %s cpf validar 00.000.198/1621-97`, project.binName, project.binName),
		Run: func(cmd *cobra.Command, args []string) {
			validarNumeroCNPJ(cmd, args)
		},
	}

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

func completarNumeroCNPJ(cmd *cobra.Command, args []string) {
	formatar, _ := cmd.Flags().GetBool("formatar")

	if len(args) == 0 {
		cmd.Help()
		fmt.Println()
		log.Fatalln("Esperava um número inteiro >= 0.")
	} else if len(args) > 1 {
		cmd.Help()
		fmt.Println()
		log.Fatalln("Esperava apenas um número de entrada.")
	}

	numero, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("%s não é um número válido.\n", args[0])
	} else if numero == 0 {
		log.Fatalln("Esperava um número inteiro >= 0.")
	}

	cnpj := rfb.NewCNPJ(uint(numero))

	if formatar {
		fmt.Println(cnpj.Formatado())
	} else {
		fmt.Println(cnpj.Desformatado())
	}
}

func validarNumeroCNPJ(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		fmt.Println()
		log.Fatalln("Esperava o número do CNPJ a ser validado.")
	} else if len(args) > 1 {
		cmd.Help()
		fmt.Println()
		log.Fatalln("Esperava apenas um número de entrada.")
	}

	cnpj, err := rfb.NewCNPJFromStr(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	if cnpj.Valido() {
		fmt.Println("CNPJ válido.")
	} else {
		fmt.Println("CNPJ INválido!")
	}
}

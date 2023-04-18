package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

type projectInfo struct {
	name    string
	version string
	binName string
}

var version = "v0.0.0-dev"

var project = projectInfo{
	name:    "go-brutils",
	version: version,
	binName: "brutils",
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   project.binName,
		Short: "Utilitário para negócios específicos do Brasil",
		Long:  fmt.Sprintf("Aplicação para uso das funções providas pela biblioteca %s.", project.name),
		Run: func(cmd *cobra.Command, args []string) {
			exibirVersao(cmd)
		},
	}

	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.AddCommand(newCPFCommand())
	cmd.AddCommand(newCNPJCommand())

	cmd.Flags().BoolP("version", "v", false, "Exibe o número da versão deste programa")

	return cmd
}

func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func exibirVersao(cmd *cobra.Command) {
	versao, _ := cmd.Flags().GetBool("version")
	if versao {
		goVersion := runtime.Version()
		osName := runtime.GOOS
		osArch := runtime.GOARCH

		fmt.Printf("Version:       %s\n", version)
		fmt.Printf("Go version:    %s\n", goVersion)
		fmt.Printf("OS/Arch:       %s/%s\n", osName, osArch)
	} else {
		_ = cmd.Help()
	}
}

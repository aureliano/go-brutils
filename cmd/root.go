package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/aureliano/caravela"
	"github.com/aureliano/caravela/provider"
	"github.com/spf13/cobra"
)

type projectInfo struct {
	name           string
	version        string
	binName        string
	scmHost        string
	scmSsl         bool
	scmProjectPath string
}

var version = "v0.0.0-dev"

var project = projectInfo{
	name:           "go-brutils",
	version:        version,
	binName:        "brutils",
	scmHost:        "api.github.com",
	scmSsl:         true,
	scmProjectPath: "aureliano/go-brutils",
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
	cmd.AddCommand(newAtualizarCommand())

	cmd.Flags().BoolP("version", "v", false, "Exibe o número da versão deste programa")

	return cmd
}

func Execute() {
	buscarAtualizacoes()
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

func buscarAtualizacoes() {
	if !strings.HasSuffix(project.version, "-dev") {
		release, err := caravela.CheckUpdates(caravela.Conf{
			Version: project.version,
			Provider: provider.GithubProvider{
				Host:        project.scmHost,
				Ssl:         project.scmSsl,
				ProjectPath: project.scmProjectPath,
			},
		})

		if err != nil {
			fmt.Printf("Verificação de novas versões falhou! %s\n", err)
		} else if release.Name != "" {
			fmt.Println(strings.Repeat("*", 80))
			fmt.Printf("[ATENÇÃO] Há uma versão mais nova do %s disponível.\nPara atualizar para a versão %s execute o comando `%s atualizar`.\n", project.name, release.Name, project.binName)
			fmt.Println(strings.Repeat("*", 80))
			fmt.Println()
		}
	}
}

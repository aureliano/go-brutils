package cmd

import (
	"fmt"

	"github.com/aureliano/caravela"
	"github.com/aureliano/caravela/provider"
	"github.com/spf13/cobra"
)

func newAtualizarCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atualizar",
		Short: "Atualiza este programa",
		Long:  "Verifica se há uma versão mais nova deste programa e o atualiza.",
		Run: func(cmd *cobra.Command, args []string) {
			atualizar()
		},
	}

	return cmd
}
func atualizar() {
	release, err := caravela.Update(caravela.Conf{
		Version: project.version,
		Provider: provider.GithubProvider{
			Host:        project.scmHost,
			Ssl:         project.scmSsl,
			ProjectPath: project.scmProjectPath,
		},
	})

	if err != nil {
		fmt.Printf("Atualização do programa falhou! %s\n", err)
	} else {
		fmt.Printf("Release %s de %s.\n\n", release.Name, release.ReleasedAt.Format("02/01/2006 15:04:05"))
		fmt.Println(release.Description)
		fmt.Printf("\nAtualização da versão %s para %s realizada com sucesso!\n", project.version, release.Name)
	}
}

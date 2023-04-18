## Contribuindo
Sinta-se à vontade para criar issues, forks do repositório e enviar pull requests!

### Relatando problemas
Bugs, solicitações de funcionalidades e perguntas relacionadas ao desenvolvimento devem ser direcionadas ao nosso [GitHub rastreador de demandas](https://github.com/aureliano/go-brutils/issues). Ao relatar um bug, tente fornecer o máximo de contexto possível como seu sistema operacional, versão Go e qualquer outra coisa que possa ser relevante para o atendimento. Para solicitações de funcionalidades, explique o que você está tentando fazer e como a funcionalidade o ajudaria a fazer isso.

Bugs relacionados à segurança podem ser relatados no rastreador de demandas.

### Enviando um patch
1. Geralmente é melhor começar abrindo um novo problema descrevendo o bug ou a funcionalidade que você pretende corrigir. Mesmo se você achar que é relativamente menor, é útil saber no que as pessoas estão trabalhando. Mencione na edição inicial que você está planejando trabalhar naquele bug ou funcionalidade para que ele possa ser atribuído a você.

2. Siga o processo normal de [fork](https://help.github.com/articles/fork-a-repo) do projeto e configure uma nova ramificação para trabalhar. É importante que cada grupo de alterações seja feito em branches separados para garantir que um pull request inclua apenas os commits relacionados a esse bug ou funcionalidade.

3. Go torna muito simples garantir que o código esteja formatado corretamente, então sempre execute `go fmt` em seu código antes de submetê-lo. Você também deve executar `go vet` ou `make code-lint` no seu código. Isso o ajudará a encontrar problemas de estilo comuns em seu código e manterá o estilo consistente no projeto.

4. Quaisquer alterações significativas devem ser quase sempre acompanhadas de testes. O projeto já possui uma boa cobertura de teste, portanto, consulte alguns dos testes existentes se não tiver certeza de como fazê-lo. [gocov](https://github.com/axw/gocov) e [gocov-html](https://github.com/matm/gocov-html) são ferramentas valiosas para ver quais partes do seu código não estão sendo cobertas por seus testes.

5. Execute:
   - `make test`
   - `make code-lint`

O comando `make test` executará testes dentro do seu código. Isso o ajudará a identificar locais onde o código pode estar com defeito antes de fazer o commit.

E o comando `make code-lint` buscará por quebras de boas práticas ou de segurança do seu código, mantendo a formatação consistente do projeto.

Faça o possível para ter [mensagens de commit eloquentes](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html) para cada alteração. Isso fornece consistência em todo o projeto e garante que as mensagens de confirmação possam ser formatadas adequadamente por várias ferramentas git.

Por fim, envie os commits para o fork e envie um [pull request](https://help.github.com/articles/creating-a-pull-request).

**OBSERVAÇÃO**: não use force-push em PRs neste repositório, pois isso torna mais difícil para os revisores ver o que mudou desde a última revisão de código.

### Outras notas sobre a organização do código
Atualmente, todos os recursos expostos - verificar atualizações e atualizar - são definidos no arquivo principal no diretório base. Tais recursos são delegados aos provedores de acordo com o local onde os lançamentos foram publicados. Dito isso, um provedor do Github deve ser um arquivo `provider/github.go` assim como `provider/gitlab.go` para o Gitlab. Portanto, use isso como seu guia para saber onde colocar novos provedores.

Atualmente, todas as funcionalidades estão dispostas em pacotes segundo o seu escopo. Como exemplo, podemos citar o pacote rfb que contém as funcionalidades para manipulação de números identificadores da Receita Federal do Brasil. Portanto, use isso como seu guia para saber onde colocar novas funcionalidades.

### Guia do Mantenedor
**Sempre tente manter um histórico limpo e linear do git**. Com pouquíssimas exceções, a execução do `git log` não deve mostrar muitas ramificações e mesclagens.

Nunca use o botão "merge" do GitHub, pois ele sempre cria uma confirmação de merge. Em vez disso, verifique o pull request localmente ([estes atalhos do git ajudam](https://github.com/willnorris/dotfiles/blob/d640d010c23b1116bdb3d4dc12088ed26120d87d/git/.gitconfig#L13-L15)) e então faça cherry-picking ou rebase no branch main. Se houver pequenos commits de limpeza, especialmente como resultado de abordar comentários de revisão de código, estes devem ser reduzidos a um único commit (squash). Não se preocupe em fazer squash de commits que realmente mereçam ser separados. Se necessário, sinta-se à vontade para corrigir pequenas alterações adicionais no código ou na mensagem de confirmação que não valham a pena passar pela revisão do código.

Se você fez alterações como compactar commits, rebasing no main etc, o GitHub não reconhecerá que este é o mesmo commit para marcar o pull request como "merged". Em vez disso, altere a mensagem de confirmação para incluir uma linha "Fixes #0", referenciando o número do pull request. Isso seria um acréscimo a quaisquer outras linhas de "Correções" para fechar problemas relacionados. Se você esquecer de fazer isso, também pode deixar um comentário no pull request. Se você fez outras alterações, vale a pena observar isso também.

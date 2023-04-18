/*
go-brutils é uma biblioteca que implementa algumas regras de negócio inerentes ao Brasil.

Algumas dessas regras, tais como geração e validação de números identificadores da Receita Federal do Brasil,
foram implementas e exportadas em pacotes específicos do escopo funcional.

Abaixo segue uma visão geral de cada pacote com implementações dessas funcionalidades. Para mais detalhes,
acesse a documentação de cada pacote.

# RFB

O pacote rfb possui as funcionalidades próprias da Receita Federal do Brasil. Tais como geração e validação de
CPF e CNPJ.

# Aplicação de linha de comando

Algumas das funcionalidades desta biblioteca estarão acessíveis para testes na aplicação de linha de comando.
Com efeito, essa aplicação será disponibilizada no binário de fechamento da release
(ver https://github.com/aureliano/go-brutils/releases).

Em desenvolvimento basta executar o comando `go run main.go help` na raiz do projeto para se ter acesso à ajuda
do programa. Se quiser usar o binário tal como é disponibilizado no fechamento da release, execute o comando
`make snapshot` e veja no diretório dist qual é binário da sua distribuição do Sistema Operacional.
*/
package main

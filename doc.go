/*
go-brutils é uma biblioteca que implementa algumas regras de negócio inerentes ao Brasil.

Algumas dessas regras, tais como geração e validação de números identificadores da Receita Federal do Brasil,
foram implementas e exportadas em pacotes específicos do escopo funcional.

# RFB

O pacote rfb possui as funcionalidades próprias da Receita Federal do Brasil. Tais como geração e validação de
CPF e CNPJ.
Para mais detalhes, acesse a documentação do pacote rfb.

# Aplicação de linha de comando

Algumas das funcionalidades desta biblioteca estarão acessíveis para testes na aplicação de linha de comando.

Em desenvolvimento basta executar o comando `go run main.go help` na raiz do projeto para se ter acesso à ajuda
do programa. Se quiser usar o binário tal como é disponibilizado no fechamento da release, execute o comando
`make snapshot` e veja no diretório dist qual é binário da sua distribuição do Sistema Operacional.
*/
package main

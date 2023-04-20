# go-brutils

[![CI Pipeline](https://github.com/aureliano/go-brutils/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/aureliano/go-brutils/actions/workflows/build.yml?query=branch%3Amain)
[![Coverage](https://coveralls.io/repos/github/aureliano/go-brutils/badge.svg?branch=main)](https://coveralls.io/github/aureliano/go-brutils?branch=main)
[![go-brutils release (latest SemVer)](https://img.shields.io/github/v/release/aureliano/go-brutils?sort=semver)](https://github.com/aureliano/go-brutils/releases)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/aureliano/go-brutils)](https://pkg.go.dev/github.com/aureliano/go-brutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-brutils é uma biblioteca que implementa algumas regras de negócio inerentes ao Brasil.

Algumas dessas regras, tais como geração e validação de números identificadores da Receita Federal do Brasil, foram implementas e exportadas em pacotes específicos do escopo funcional.

Atualmente, **go-brutils requer a versão 1.17 ou superior do Go**. Seguimos a política de suporte de versões do Go e fazemos o melhor pra não quebrar as versões mais velhas do Go, mas devido a restrições de ferramentas, nem sempre testamos versões mais antigas.

## Instalação
Para instalar go-brutils, use `go get`:

`go get github.com/aureliano/go-brutils`

Ou instale uma versão específica:

`go get github.com/aureliano/go-brutils/v1`

Ou mesmo adicione-a como uma dependência do seu módulo:

`require github.com/aureliano/go-brutils v1`

Para instalar a aplicação de linha de comando, caso queira instalar no GOPATH:

`go install github.com/aureliano/go-brutils@v1`

Ficará acessível como `go-brutils`.

Para instalar via GitHub no Linux:

```sh
curl -OL https://github.com/aureliano/go-brutils/releases/download/v1.0.0/go-brutils_Linux_x86_64.tar.gz
tar xzf go-brutils_Linux_x86_64.tar.gz
```

Ficará acessível no diretório corrente como `./brutils`.

### Fique atualizado
Para atualizar para a última versão, use `go get -u github.com/aureliano/go-brutils`.

## Uso

### RFB

O pacote rfb possui as funcionalidades próprias da Receita Federal do Brasil. Tais como geração e validação de CPF e CNPJ.

### Iniciação dos NIs
Para cada NI há um tipo personalizado. A iniciação desses tipos deve ser feita através das funções adequadas. Há duas formas de se obter um NI: passando o número base (um inteiro não negativo) ou passando o número completo como string com ou sem máscara.

```go
// Passando o número base do CPF "000.001.234-39".
cpf := rfb.NewCPF(1234)

// Passando o número do CPF como texto sem máscara.
cpf, err := rfb.NewCPFFromStr("00000123439")

// Passando o número do CPF como texto com máscara.
cpf, err := rfb.NewCPFFromStr("000.001.234-39")

cpf.NumeroBase()           // 1234
cpf.DigitosVerificadores() // [3, 9]
cpf.Valido()               // true
cpf.Formatado()            // 000.001.234-39
cpf.Desformatado()         // 00000123439
```

## Aplicação de linha de comando

Algumas das funcionalidades desta biblioteca estarão acessíveis para testes na aplicação de linha de comando. Com efeito, essa aplicação será disponibilizada no binário de fechamento da [release](https://github.com/aureliano/go-brutils/releases).

Em desenvolvimento basta executar o comando `go run main.go help` na raiz do projeto para se ter acesso à ajuda do programa. Se quiser usar o binário tal como é disponibilizado no fechamento da release, execute o comando `make snapshot` e veja no diretório dist qual é binário da sua distribuição do Sistema Operacional.

## Contribuindo

Sinta-se a vontade para criar issues, forks e enviar pull requests. Mas antes, leia [este guia](./CONTRIBUTING.md) para obter orientações de como contribuir da melhor forma.

## Licença
Este projeto está disponível sob os termo da licença MIT, que pode ser encontrada no arquivo [LICENSE](./LICENSE).

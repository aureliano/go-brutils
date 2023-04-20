package rfb

import (
	"bytes"
	"strconv"
)

// NIRFB é a interface que os números identificadores devem implementar,
// pois provê as funcionalidades da Receita Federal do Brasil inerentes
// aos NIs. Ex.: CPF, CNPJ.
type NIRFB interface {
	// NumeroBase retorna o número base de um NI. Ou seja, os nove primeiros
	// dígitos do CPF ou os doze primeiros dígitos do CNPJ.
	//
	// Retorna o número base como um número inteiro >= zero.
	NumeroBase() uint

	// DigitosVerificadores retorna os dígitos verificadores de um NI. Ou seja,
	// os dois últimos dígitos ora do CPF ora do CNPJ, que são obtidos através
	// de um cálculo sobre os seus antecessores.
	//
	// Retorna os dois dígitos verificadores. Em caso de falha, retorna -1, -1.
	DigitosVerificadores() (int, int)

	// Valido aplica o algoritmo de validação dos dígitos verificadores sobre o
	// número base.
	//
	// Retorna verdadeiro caso o NI seja válido.
	Valido() bool

	// Formatado aplica uma máscara para formatação do NI.
	// Máscara para CPF: ###.###.###-##
	// Máscara para CNPJ: ##.###.###/####-##
	//
	// Retorna uma string com o NI formatado pela máscara.
	Formatado() string

	// Desformatado remove a máscara de formatação do NI.
	//
	// Retorna uma string com o NI sem máscara.
	Desformatado() string
}

// NewEstado instancia uma estrutura de dados do tipo Estado dada uma sigla de
// unidade federativa.
//
// Retorna um ponteiro para um tipo Estado.
func NewEstado(uf string) *Estado {
	e := newEstado(uf)
	if e == nil {
		e = &Estado{}
	}

	return e
}

// GerarCPF gera um CPF válido tomando, aleatoriamente, uma região fiscal.
//
// Retorna um CPF.
func GerarCPF() (CPF, error) {
	return gerarCPF()
}

// GerarCPFParaUF gera um CPF para uma região fiscal dada como parâmetro de entrada.
//
// Retorna um CPF.
func GerarCPFParaUF(uf *Estado) (CPF, error) {
	return gerarCPFParaUF(uf.UF)
}

// NewCPF instancia um CPF a partir do número base. Sendo o número base um inteiro
// maior ou igual a zero, composto dos nove primeiros dígitos do CPF.
//
// Retorna um CPF.
func NewCPF(numbBase uint) CPF {
	return newCPF(numbBase)
}

// NewCPF instancia um CPF a partir do número, com ou sem máscara. O parâmetro de entrada
// deve ser uma string representando um número de onze caracteres, com ou sem máscara.
//
// Retorna um CPF.
func NewCPFFromStr(str string) (CPF, error) {
	return newCPFFromStr(str)
}

// NumeroBase retorna o número base do CPF. Ou seja, os nove primeiros dígitos.
//
// Retorna o número base como um inteiro >= zero.
func (cpf CPF) NumeroBase() uint {
	if !cpfNumeral(cpf) {
		return 0
	}

	base, rf := recuperarNumeroBaseCPF(string(cpf))

	var b bytes.Buffer
	for _, d := range base {
		b.WriteString(strconv.Itoa(d))
	}

	b.WriteString(strconv.Itoa(rf))
	num, _ := strconv.Atoi(b.String())

	return uint(num)
}

// DigitosVerificadores retorna os dígitos verificadores do CPF. Ou seja,
// os dois últimos dígitos.
//
// Retorna os dois dígitos verificadores. Em caso de falha, retorna -1, -1.
func (cpf CPF) DigitosVerificadores() (int, int) {
	if !cpfNumeral(cpf) {
		return -1, -1
	}

	scpf := string(cpf)

	dv1, _ := strconv.Atoi(scpf[9:10])
	dv2, _ := strconv.Atoi(scpf[10:11])

	return dv1, dv2
}

// Valido aplica o algoritmo de validação dos dígitos verificadores sobre o
// número base.
//
// Retorna verdadeiro caso o CPF seja válido.
func (cpf CPF) Valido() bool {
	return cpfValido(cpf)
}

// Formatado aplica uma máscara para formatação do CPF (###.###.###-##)
// Exemplo: dado o número 12345678901 temos 123.456.789-01
//
// Retorna uma string com o CPF formatado pela máscara.
func (cpf CPF) Formatado() string {
	return formatarCPF(cpf)
}

// Desformatado remove a máscara de formatação do CPF.
//
// Retorna uma string com o CPF sem máscara.
func (cpf CPF) Desformatado() string {
	return desformatarCPF(cpf)
}

// GerarCNPJ gera um CNPJ válido com números gerados aleatoriamente.
//
// Retorna um CNPJ.
func GerarCNPJ() (CNPJ, error) {
	return gerarCNPJ()
}

// NewCNPJ instancia um CNPJ a partir do número base. Sendo o número base um inteiro
// maior ou igual a zero, composto dos doze primeiros dígitos do CNPJ.
//
// Retorna um CNPJ.
func NewCNPJ(numbBase uint) CNPJ {
	return newCNPJ(numbBase)
}

// NewCNPJ instancia um CNPJ a partir do número, com ou sem máscara. O parâmetro de entrada
// deve ser uma string representando um número de quatorze caracteres, com ou sem máscara.
//
// Retorna um CNPJ.
func NewCNPJFromStr(str string) (CNPJ, error) {
	return newCNPJFromStr(str)
}

// NumeroBase retorna o número base do CNPJ. Ou seja, os doze primeiros dígitos.
//
// Retorna o número base como um inteiro >= zero.
func (cnpj CNPJ) NumeroBase() uint {
	if !cnpjNumeral(cnpj) {
		return 0
	}

	base := recuperarNumeroBaseCNPJ(string(cnpj))

	var b bytes.Buffer
	for _, d := range base {
		b.WriteString(strconv.Itoa(d))
	}

	num, _ := strconv.Atoi(b.String())

	return uint(num)
}

// DigitosVerificadores retorna os dígitos verificadores do CNPJ. Ou seja,
// os dois últimos dígitos.
//
// Retorna os dois dígitos verificadores. Em caso de falha, retorna -1, -1.
func (cnpj CNPJ) DigitosVerificadores() (int, int) {
	if !cnpjNumeral(cnpj) {
		return -1, -1
	}

	scnpj := string(cnpj)

	dv1, _ := strconv.Atoi(scnpj[12:13])
	dv2, _ := strconv.Atoi(scnpj[13:14])

	return dv1, dv2
}

// Valido aplica o algoritmo de validação dos dígitos verificadores sobre o
// número base.
//
// Retorna verdadeiro caso o CNPJ seja válido.
func (cnpj CNPJ) Valido() bool {
	return cnpjValido(cnpj)
}

// Formatado aplica uma máscara para formatação do CNPJ (##.###.###/####-##)
// Exemplo: dado o número 12345678000123 temos 12.345.678/0001-23
//
// Retorna uma string com o CNPJ formatado pela máscara.
func (cnpj CNPJ) Formatado() string {
	return formatarCNPJ(cnpj)
}

// Desformatado remove a máscara de formatação do CNPJ.
//
// Retorna uma string com o CNPJ sem máscara.
func (cnpj CNPJ) Desformatado() string {
	return desformatarCNPJ(cnpj)
}

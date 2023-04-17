package rfb

import (
	"bytes"
	"strconv"
)

type NIRFB interface {
	NumeroBase() uint
	DigitosVerificadores() (int, int)
	Valido() bool
	Formatado() string
	Desformatado() string
}

func NewEstado(uf string) *Estado {
	e := newEstado(uf)
	if e == nil {
		e = &Estado{}
	}

	return e
}

func GerarCPF() (CPF, error) {
	return gerarCPF()
}

func GerarCPFParaUF(uf *Estado) (CPF, error) {
	return gerarCPFParaUF(uf.UF)
}

func NewCPF(numbBase uint) CPF {
	return newCPF(numbBase)
}

func NewCPFFromStr(str string) (CPF, error) {
	return newCPFFromStr(str)
}

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

func (cpf CPF) DigitosVerificadores() (int, int) {
	if !cpfNumeral(cpf) {
		return -1, -1
	}

	scpf := string(cpf)

	dv1, _ := strconv.Atoi(scpf[9:10])
	dv2, _ := strconv.Atoi(scpf[10:11])

	return dv1, dv2
}

func (cpf CPF) Valido() bool {
	return cpfValido(cpf)
}

func (cpf CPF) Formatado() string {
	return formatarCPF(cpf)
}

func (cpf CPF) Desformatado() string {
	return desformatarCPF(cpf)
}

func GerarCNPJ() (CNPJ, error) {
	return gerarCNPJ()
}

func NewCNPJ(numbBase uint) CNPJ {
	return newCNPJ(numbBase)
}

func NewCNPJFromStr(str string) (CNPJ, error) {
	return newCNPJFromStr(str)
}

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

func (cnpj CNPJ) DigitosVerificadores() (int, int) {
	if !cnpjNumeral(cnpj) {
		return -1, -1
	}

	scnpj := string(cnpj)

	dv1, _ := strconv.Atoi(scnpj[12:13])
	dv2, _ := strconv.Atoi(scnpj[13:14])

	return dv1, dv2
}

func (cnpj CNPJ) Valido() bool {
	return cnpjValido(cnpj)
}

func (cnpj CNPJ) Formatado() string {
	return formatarCNPJ(cnpj)
}

func (cnpj CNPJ) Desformatado() string {
	return desformatarCNPJ(cnpj)
}

package rfb

import (
	"bytes"
	"strconv"
)

type CPFInt interface {
	NumeroBase() uint
	DigitosVerificadores() (int, int)
	Estado() *Estado
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

func GerarCPFParaUF(uf string) (CPF, error) {
	return gerarCPFParaUF(uf)
}

func NewCPF(numbBase uint) CPF {
	return newCPF(numbBase)
}

func (cpf CPF) NumeroBase() uint {
	if !cpfNumeral(cpf) {
		return 0
	}

	base, rf := recuperarNumeroBase(string(cpf))

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

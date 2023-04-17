package rfb

import (
	"bytes"
	"strconv"
)

type CPFInt interface {
	NumeroBase() uint
	CodigoVerificador() uint
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
	base, rf := recuperarNumeroBase(string(cpf))

	var b bytes.Buffer
	for _, d := range base {
		b.WriteString(strconv.Itoa(d))
	}

	b.WriteString(strconv.Itoa(rf))
	num, _ := strconv.Atoi(b.String())

	return uint(num)
}

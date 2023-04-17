package rfb_test

import (
	"errors"
	"testing"

	"github.com/aureliano/go-brutils/rfb"
	"github.com/stretchr/testify/assert"
)

func TestNewEstadoInvalido(t *testing.T) {
	e := rfb.NewEstado("")
	assert.Equal(t, "", e.UF)
	assert.Equal(t, "", e.Nome)
	assert.Equal(t, 0, e.RegiaoFiscal)

	e = rfb.NewEstado("br")
	assert.Equal(t, "", e.UF)
	assert.Equal(t, "", e.Nome)
	assert.Equal(t, 0, e.RegiaoFiscal)
}

func TestNewEstado(t *testing.T) {
	e := rfb.NewEstado("mg")
	assert.Equal(t, "MG", e.UF)
	assert.Equal(t, "Minas Gerais", e.Nome)
	assert.Equal(t, 6, e.RegiaoFiscal)
}

func TestGerarCPF(t *testing.T) {
	cpf, err := rfb.GerarCPF()
	assert.Nil(t, err)
	assert.Regexp(t, `^\d{11}$`, cpf)
}

func TestGerarCPFParaUf(t *testing.T) {
	uf := &rfb.Estado{UF: "mg"}
	cpf, err := rfb.GerarCPFParaUF(uf)
	assert.Nil(t, err)
	assert.Regexp(t, `^\d{11}$`, cpf)

	uf.UF = "br"
	_, err = rfb.GerarCPFParaUF(uf)
	assert.True(t, errors.Is(err, rfb.ErrUFDesconhecida))
}

func TestNewCPF(t *testing.T) {
	cpf := rfb.NewCPF(1234)
	assert.Equal(t, rfb.CPF("00000123439"), cpf)
}

func TestNumeroBaseCPF(t *testing.T) {
	cpf := rfb.CPF("12345678900")
	expected := uint(123456789)
	actual := cpf.NumeroBase()

	assert.Equal(t, expected, actual)

	cpf = rfb.CPF("00005678900")
	expected = uint(56789)
	actual = cpf.NumeroBase()

	assert.Equal(t, expected, actual)

	cpf = rfb.CPF("")
	expected = uint(0)
	actual = cpf.NumeroBase()

	assert.Equal(t, expected, actual)

	cpf = rfb.CPF("a1111122223")
	expected = uint(0)
	actual = cpf.NumeroBase()

	assert.Equal(t, expected, actual)
}

func TestDigitosVerificadoresCPF(t *testing.T) {
	cpf := rfb.CPF("12345678954")
	edv1, edv2 := 5, 4
	dv1, dv2 := cpf.DigitosVerificadores()

	assert.Equal(t, edv1, dv1)
	assert.Equal(t, edv2, dv2)

	cpf = rfb.CPF("345678954")
	edv1, edv2 = -1, -1
	dv1, dv2 = cpf.DigitosVerificadores()

	assert.Equal(t, edv1, dv1)
	assert.Equal(t, edv2, dv2)
}

func TestValidoCPF(t *testing.T) {
	cpf := rfb.CPF("12345678954")
	assert.False(t, cpf.Valido())

	cpf = rfb.CPF("")
	assert.False(t, cpf.Valido())

	cpf = rfb.CPF("93976464279")
	assert.True(t, cpf.Valido())
}

func TestFormatadoCPF(t *testing.T) {
	cpf := rfb.CPF("12345678954")
	assert.Equal(t, "123.456.789-54", cpf.Formatado())
}

func TestDesformatadoCPF(t *testing.T) {
	cpf := rfb.CPF("123.456.789-54")
	assert.Equal(t, "12345678954", cpf.Desformatado())
}

func TestGerarCNPJ(t *testing.T) {
	cnpj, err := rfb.GerarCNPJ()
	assert.Nil(t, err)
	assert.Regexp(t, `^\d{14}$`, cnpj)
}

func TestNewCNPJ(t *testing.T) {
	cnpj := rfb.NewCNPJ(1234)
	assert.Equal(t, rfb.CNPJ("00000000123439"), cnpj)
}

func TestNumeroBaseCNPJ(t *testing.T) {
	cnpj := rfb.CNPJ("12345678900000")
	expected := uint(123456789000)
	actual := cnpj.NumeroBase()

	assert.Equal(t, expected, actual)

	cnpj = rfb.CNPJ("00005678900000")
	expected = uint(56789000)
	actual = cnpj.NumeroBase()

	assert.Equal(t, expected, actual)

	cnpj = rfb.CNPJ("")
	expected = uint(0)
	actual = cnpj.NumeroBase()

	assert.Equal(t, expected, actual)

	cnpj = rfb.CNPJ("a1111122223222")
	expected = uint(0)
	actual = cnpj.NumeroBase()

	assert.Equal(t, expected, actual)
}

func TestDigitosVerificadoresCNPJ(t *testing.T) {
	cnpj := rfb.CNPJ("12345678000154")
	edv1, edv2 := 5, 4
	dv1, dv2 := cnpj.DigitosVerificadores()

	assert.Equal(t, edv1, dv1)
	assert.Equal(t, edv2, dv2)

	cnpj = rfb.CNPJ("3456789500001")
	edv1, edv2 = -1, -1
	dv1, dv2 = cnpj.DigitosVerificadores()

	assert.Equal(t, edv1, dv1)
	assert.Equal(t, edv2, dv2)
}

func TestValidoCNPJ(t *testing.T) {
	cnpj := rfb.CNPJ("12345678954225")
	assert.False(t, cnpj.Valido())

	cnpj = rfb.CNPJ("")
	assert.False(t, cnpj.Valido())

	cnpj = rfb.CNPJ("00009006000120")
	assert.True(t, cnpj.Valido())
}

func TestFormatadoCNPJ(t *testing.T) {
	cnpj := rfb.CNPJ("00009006000120")
	assert.Equal(t, "00.009.006/0001-20", cnpj.Formatado())
}

func TestDesformatadoCNPJ(t *testing.T) {
	cnpj := rfb.CNPJ("00.009.006/0001-20")
	assert.Equal(t, "00009006000120", cnpj.Desformatado())
}

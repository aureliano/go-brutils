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
	cpf, err := rfb.GerarCPFParaUF("mg")
	assert.Nil(t, err)
	assert.Regexp(t, `^\d{11}$`, cpf)

	_, err = rfb.GerarCPFParaUF("br")
	assert.True(t, errors.Is(err, rfb.ErrUFDesconhecida))
}

func TestNewCPF(t *testing.T) {
	cpf := rfb.NewCPF(1234)
	assert.Equal(t, rfb.CPF("00000123431"), cpf)
}

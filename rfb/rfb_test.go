package rfb_test

import (
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

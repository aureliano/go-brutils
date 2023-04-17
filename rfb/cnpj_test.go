package rfb

import (
	"fmt"
	"testing"

	"github.com/aureliano/go-brutils/number"
	"github.com/stretchr/testify/assert"
)

func TestGerarCNPJ(t *testing.T) {
	mpGerarUnidadeDecimalCNPJ = number.GerarUnidadeDecimal

	cnpj, _ := gerarCNPJ()
	assert.Regexp(t, cnpjNumeralRegex, cnpj)
}

func TestGerarCNPJErroGeracaoNumBase(t *testing.T) {
	mpGerarUnidadeDecimalCNPJ = number.GerarUnidadeDecimal
	mpGerarNumeroBaseCNPJ = func() ([]int, error) { return nil, ErrGeracaoCNPJ }

	_, err := gerarCNPJ()
	assert.ErrorIs(t, err, ErrGeracaoCNPJ)

	mpGerarNumeroBaseCNPJ = gerarNumeroBaseCNPJ
}

func TestGerarNumeroBaseCNPJErro(t *testing.T) {
	mpGerarUnidadeDecimalCNPJ = func() (int, error) { return -1, fmt.Errorf("any error") }

	_, err := gerarNumeroBaseCNPJ()
	assert.ErrorIs(t, err, ErrGeracaoCNPJ)

	mpGerarUnidadeDecimalCNPJ = number.GerarUnidadeDecimal
}

func TestGerarDigitosVerificadoresCNPJ(t *testing.T) {
	dv1, dv2 := gerarDigitosVerificadoresCNPJ([]int{1, 1, 4, 4, 4, 7, 7, 7, 0, 0, 0, 1})
	assert.Equal(t, 6, dv1)
	assert.Equal(t, 1, dv2)

	dv1, dv2 = gerarDigitosVerificadoresCNPJ([]int{6, 2, 4, 4, 4, 7, 7, 7, 0, 0, 0, 1})
	assert.Equal(t, 0, dv1)
	assert.Equal(t, 0, dv2)
}

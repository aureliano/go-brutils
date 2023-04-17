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

func TestNewCNPJ(t *testing.T) {
	type testCase struct {
		name     string
		input    uint
		expected CNPJ
	}
	testCases := []testCase{
		{name: "Matriz com menos dígitos", input: 90060001, expected: "00009006000120"},
		{name: "Matriz com todos os dígitos", input: 111300130001, expected: "11130013000100"},
		{name: "Filial unidade", input: 981123280002, expected: "98112328000285"},
		{name: "Filial dezena", input: 793805000021, expected: "79380500002108"},
		{name: "Filial centena", input: 693451290321, expected: "69345129032170"},
		{name: "Filial milhar", input: 564599554321, expected: "56459955432112"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := newCNPJ(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
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

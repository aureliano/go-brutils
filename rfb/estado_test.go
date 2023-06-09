package rfb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEstado(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected *Estado
	}
	testCases := []testCase{
		{
			name:     "deve retornar Minas Gerais",
			input:    "MG",
			expected: &Estado{"MG", "Minas Gerais", 6},
		},
		{
			name:     "deve retornar Pará",
			input:    "pa",
			expected: &Estado{"PA", "Pará", 2},
		},
		{
			name:     "deve retornar Alagoas",
			input:    "Al",
			expected: &Estado{"AL", "Alagoas", 4},
		},
		{
			name:     "deve retornar nulo",
			input:    "BR",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := newEstado(tc.input)
			assert.True(t, estadosSaoIguais(tc.expected, actual))
		})
	}
}

func estadosSaoIguais(e1, e2 *Estado) bool {
	switch {
	case e1 == nil && e2 == nil:
		return true
	case (e1 != nil && e2 == nil) || (e1 == nil && e2 != nil):
		return false
	default:
		return (e1.Nome == e2.Nome) && (e1.UF == e2.UF) && (e1.RegiaoFiscal == e2.RegiaoFiscal)
	}
}

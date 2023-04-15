package rfb

import "testing"

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
			if !estadosSaoIguais(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func estadosSaoIguais(e1, e2 *Estado) bool {
	if e1 == nil && e2 == nil {
		return true
	} else if (e1 != nil && e2 == nil) || (e1 == nil && e2 != nil) {
		return false
	} else {
		return (e1.Nome == e2.Nome) && (e1.UF == e2.UF) && (e1.RegiaoFiscal == e2.RegiaoFiscal)
	}
}

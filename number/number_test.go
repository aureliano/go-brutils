package number_test

import (
	"testing"

	"github.com/aureliano/go-brutils/number"
	"github.com/stretchr/testify/assert"
)

func TestExtrairNumeros(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected string
	}
	testCases := []testCase{
		{name: "empty string", input: "", expected: ""},
		{name: "only text", input: "abcdef", expected: ""},
		{name: "only number", input: "12345", expected: "12345"},
		{name: "formatted CPF", input: "012.345.678-90", expected: "01234567890"},
		{name: "formatted CNPJ", input: "01.234.567/0001-00", expected: "01234567000100"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := number.ExtrairNumeros(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGerarUnidadeDecimal(t *testing.T) {
	for i := 0; i < 100; i++ {
		ud, _ := number.GerarUnidadeDecimal()
		assert.True(t, ud >= 0)
		assert.True(t, ud < 10)
	}
}

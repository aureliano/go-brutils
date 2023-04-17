package number

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

var numbersRegex = regexp.MustCompilePOSIX("[0-9]+")

// ExtrairNumeros remove todo caractere que não for um número decimal da cadeia de texto.
//
// Retorna apenas os números.
func ExtrairNumeros(text string) string {
	nums := numbersRegex.FindAllString(text, -1)
	return strings.Join(nums, "")
}

// GerarUnidadeDecimal gera, aleatoriamente, um número decimal entre zero e nove.
//
// Retorna um inteiro decimal de zero a nove ou erro caso falhe o sorteio.
func GerarUnidadeDecimal() (int, error) {
	const maxNum = 10
	n, err := rand.Int(rand.Reader, big.NewInt(maxNum))

	return int(n.Int64()), err
}

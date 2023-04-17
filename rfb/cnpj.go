package rfb

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/aureliano/go-brutils/number"
)

type CNPJ string

var mpGerarUnidadeDecimalCNPJ = number.GerarUnidadeDecimal
var mpGerarNumeroBaseCNPJ = gerarNumeroBaseCNPJ
var cnpjNumeralRegex = regexp.MustCompile(`^\d{14}$`)
var cnpjFormatadoRegex = regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)

const cnpjSize = 14
const cnpjNumBaseSize = cnpjSize - 2

func gerarCNPJ() (CNPJ, error) {
	base, err := mpGerarNumeroBaseCNPJ()
	if err != nil {
		return "", err
	}

	dv1, dv2 := gerarDigitosVerificadoresCNPJ(base)

	return CNPJ(writeCNPJ(base, dv1, dv2)), nil
}

func newCNPJ(numBase uint) CNPJ {
	base := fmt.Sprintf("%012d", numBase)[:cnpjNumBaseSize]
	ibase := recuperarNumeroBaseCNPJ(base)

	dv1, dv2 := gerarDigitosVerificadoresCNPJ(ibase)

	return CNPJ(writeCNPJ(ibase, dv1, dv2))
}

func newCNPJFromStr(str string) (CNPJ, error) {
	if cnpjFormatadoRegex.MatchString(str) {
		nums := number.ExtrairNumeros(str)
		return CNPJ(nums), nil
	}

	if !cnpjNumeral(CNPJ(str)) {
		return "", ErrCNPJInvalido
	}

	return CNPJ(str), nil
}

func cnpjValido(cnpj CNPJ) bool {
	if !cnpjNumeral(cnpj) {
		return false
	}

	strCnpj := string(cnpj)
	base := recuperarNumeroBaseCNPJ(strCnpj)

	dv1, dv2 := gerarDigitosVerificadoresCNPJ(base)

	edv1, _ := strconv.Atoi(strCnpj[12:13])
	edv2, _ := strconv.Atoi(strCnpj[13:14])

	return (edv1 == dv1) && (edv2 == dv2)
}

func cnpjNumeral(cnpj CNPJ) bool {
	return cnpjNumeralRegex.MatchString(string(cnpj))
}

func formatarCNPJ(cnpj CNPJ) string {
	scnpj := string(cnpj)

	if len(cnpj) != cnpjSize {
		return scnpj
	}

	return fmt.Sprintf("%s.%s.%s/%s-%s", scnpj[0:2], scnpj[2:5], scnpj[5:8], scnpj[8:12], scnpj[12:14])
}

func desformatarCNPJ(cnpj CNPJ) string {
	scnpj := string(cnpj)
	return number.ExtrairNumeros(scnpj)
}

func writeCNPJ(base []int, dv1, dv2 int) string {
	var b bytes.Buffer
	for _, d := range base {
		b.WriteString(strconv.Itoa(d))
	}

	b.WriteString(strconv.Itoa(dv1))
	b.WriteString(strconv.Itoa(dv2))

	return b.String()
}

func gerarNumeroBaseCNPJ() ([]int, error) {
	num := make([]int, cnpjNumBaseSize)

	for i := 0; i < cnpjNumBaseSize; i++ {
		n, err := mpGerarUnidadeDecimalCNPJ()
		if err != nil {
			return nil, ErrGeracaoCNPJ
		}

		num[i] = n
	}

	return num, nil
}

func gerarDigitosVerificadoresCNPJ(base []int) (int, int) {
	sum := 0
	const minUnit = 2
	const modSeed = 11
	const lastPos = 2
	const blk1Size = 4
	const blk2Size = 8

	for i := 0; i < cnpjNumBaseSize; i++ {
		d := base[i]

		if i < blk1Size {
			sum += d * (blk1Size + 1 - i)
		} else {
			sum += d * (blk2Size + 5 - i)
		}
	}

	dv1 := sum % modSeed
	if dv1 < minUnit {
		dv1 = 0
	} else {
		dv1 = modSeed - dv1
	}

	sum = 0
	for i := 0; i < cnpjNumBaseSize; i++ {
		d := base[i]

		if i < (blk1Size + 1) {
			sum += d * (blk1Size + 2 - i)
		} else {
			sum += d * (blk2Size + 6 - i)
		}
	}
	sum += dv1 * lastPos

	dv2 := sum % modSeed
	if dv2 < minUnit {
		dv2 = 0
	} else {
		dv2 = modSeed - dv2
	}

	return dv1, dv2
}

func recuperarNumeroBaseCNPJ(cnpj string) []int {
	base := make([]int, cnpjNumBaseSize)

	for i, n := range strings.Split(cnpj, "")[:cnpjNumBaseSize] {
		v, _ := strconv.Atoi(n)
		base[i] = v
	}

	return base
}

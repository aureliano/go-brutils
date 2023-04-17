package rfb

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/aureliano/go-brutils/number"
)

type CPF string

var mpGenRandomDecimalUnit = number.GerarUnidadeDecimal
var mpGerarNumeroBaseCPF = gerarNumeroBaseCPF
var mpGerarCodigoRegiaoFiscal = gerarCodigoRegiaoFiscal
var cpfNumeralRegex = regexp.MustCompile(`^\d{11}$`)

const cpfSize = 11
const cpfNumBaseSize = 9

func gerarCPF() (CPF, error) {
	return gerarCPFParaUF("")
}

func gerarCPFParaUF(uf string) (CPF, error) {
	base, err := mpGerarNumeroBaseCPF()
	if err != nil {
		return "", err
	}

	rf, err := mpGerarCodigoRegiaoFiscal(uf)
	if err != nil {
		return "", err
	}

	dv1, dv2 := gerarDigitosVerificadoresCPF(base, rf)

	return CPF(writeCPF(base, rf, dv1, dv2)), nil
}

func newCPF(numBase uint) CPF {
	base := fmt.Sprintf("%09d", numBase)[:cpfNumBaseSize]
	ibase, rf := recuperarNumeroBase(base)

	dv1, dv2 := gerarDigitosVerificadoresCPF(ibase, rf)

	return CPF(writeCPF(ibase, rf, dv1, dv2))
}

func cpfValido(cpf CPF) bool {
	if !cpfNumeral(cpf) {
		return false
	}

	strCpf := string(cpf)
	base, rf := recuperarNumeroBase(strCpf)

	dv1, dv2 := gerarDigitosVerificadoresCPF(base, rf)

	edv1, _ := strconv.Atoi(strCpf[9:10])
	edv2, _ := strconv.Atoi(strCpf[10:11])

	return (edv1 == dv1) && (edv2 == dv2)
}

func cpfNumeral(cpf CPF) bool {
	return cpfNumeralRegex.MatchString(string(cpf))
}

func formatarCPF(cpf CPF) string {
	scpf := string(cpf)

	if len(cpf) != cpfSize {
		return scpf
	}

	return fmt.Sprintf("%s.%s.%s-%s", scpf[0:3], scpf[3:6], scpf[6:9], scpf[9:11])
}

func desformatarCPF(cpf CPF) string {
	scpf := string(cpf)
	return number.ExtrairNumeros(scpf)
}

func writeCPF(base []int, rf, dv1, dv2 int) string {
	var b bytes.Buffer
	for _, d := range base {
		b.WriteString(strconv.Itoa(d))
	}

	b.WriteString(strconv.Itoa(rf))
	b.WriteString(strconv.Itoa(dv1))
	b.WriteString(strconv.Itoa(dv2))

	return b.String()
}

func gerarNumeroBaseCPF() ([]int, error) {
	digitos := 8
	num := make([]int, digitos)

	for i := 0; i < digitos; i++ {
		n, err := mpGenRandomDecimalUnit()
		if err != nil {
			return nil, ErrGeracaoCPF
		}

		num[i] = n
	}

	return num, nil
}

func gerarCodigoRegiaoFiscal(uf string) (int, error) {
	if uf == "" {
		n, err := mpGenRandomDecimalUnit()
		if err != nil {
			return -1, ErrGeracaoCPF
		}

		return n, nil
	}

	estado := newEstado(uf)

	if estado == nil {
		return -1, ErrUFDesconhecida
	}

	return estado.RegiaoFiscal, nil
}

func gerarDigitosVerificadoresCPF(base []int, rf int) (int, int) {
	sum := 0
	const baseSize = 8
	const maxUnit = 9
	const ten = 10
	const lastPos = 2

	for i := 0; i < baseSize; i++ {
		d := base[i]
		seed := cpfSize - 1 - i
		sum += d * seed
	}
	sum += rf * lastPos

	dv1 := (sum * ten) % cpfSize
	if dv1 > maxUnit {
		dv1 = 0
	}

	sum = 0
	for i := 0; i < baseSize; i++ {
		d := base[i]
		seed := cpfSize - i
		sum += d * seed
	}
	sum += rf * (lastPos + 1)
	sum += dv1 * lastPos

	dv2 := (sum * ten) % cpfSize
	if dv2 > maxUnit {
		dv2 = 0
	}

	return dv1, dv2
}

func recuperarNumeroBase(cpf string) ([]int, int) {
	base := make([]int, cpfNumBaseSize-1)
	rf := -1

	for i, n := range strings.Split(cpf, "")[:cpfNumBaseSize] {
		v, _ := strconv.Atoi(n)
		if i == cpfNumBaseSize-1 {
			rf = v
		} else {
			base[i] = v
		}
	}

	return base, rf
}

package rfb

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type CPF string

var mpGenRandomDecimalUnit = genRandomDecimalUnit
var mpGerarNumeroBase = gerarNumeroBase
var mpGerarCodigoRegiaoFiscal = gerarCodigoRegiaoFiscal

func gerarCPF() (CPF, error) {
	return gerarCPFParaUF("")
}

func gerarCPFParaUF(uf string) (CPF, error) {
	base, err := mpGerarNumeroBase()
	if err != nil {
		return "", err
	}

	rf, err := mpGerarCodigoRegiaoFiscal(uf)
	if err != nil {
		return "", err
	}

	dv1, dv2 := gerarDigitosVerificadores(base, rf)

	return CPF(writeCPF(base, rf, dv1, dv2)), nil
}

func newCPF(numBase uint) CPF {
	const maxDigitos = 9
	base := fmt.Sprintf("%09d", numBase)[:maxDigitos]
	ibase := make([]int, maxDigitos-1)
	rf := -1

	for i, n := range strings.Split(base, "") {
		v, _ := strconv.Atoi(n)
		if i == maxDigitos-1 {
			rf = v
		} else {
			ibase[i] = v
		}
	}

	dv1, dv2 := gerarDigitosVerificadores(ibase, rf)

	return CPF(writeCPF(ibase, rf, dv1, dv2))
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

func gerarNumeroBase() ([]int, error) {
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

func gerarDigitosVerificadores(base []int, rf int) (int, int) {
	sum := 0
	const baseSize = 8
	const maxUnit = 9
	const cpfSize = 11
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

func genRandomDecimalUnit() (int, error) {
	const maxNum = 10
	n, err := rand.Int(rand.Reader, big.NewInt(maxNum))

	return int(n.Int64()), err
}

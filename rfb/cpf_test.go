package rfb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGerarCPF(t *testing.T) {
	mpGenRandomDecimalUnit = genRandomDecimalUnit

	cpf, _ := gerarCPF()
	fmt.Println(cpf)
	assert.Regexp(t, `^\d{11}$`, cpf)
}

func TestGerarCPFParaUF(t *testing.T) {
	mpGenRandomDecimalUnit = genRandomDecimalUnit

	cpf, _ := gerarCPFParaUF("")
	assert.Regexp(t, `^\d{11}$`, cpf)

	cpf, _ = gerarCPFParaUF("mg")
	assert.Regexp(t, `^\d{11}$`, cpf)

	_, err := gerarCPFParaUF("br")
	assert.ErrorIs(t, err, ErrUFDesconhecida)
}

func TestGerarCPFParaUFErroGeracaoNumBase(t *testing.T) {
	mpGenRandomDecimalUnit = genRandomDecimalUnit
	mpGerarNumeroBase = func() ([]int, error) { return nil, ErrGeracaoCPF }

	_, err := gerarCPFParaUF("mg")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarNumeroBase = gerarNumeroBase
}

func TestGerarCPFParaUFErroCodRecFiscal(t *testing.T) {
	mpGenRandomDecimalUnit = genRandomDecimalUnit
	mpGerarNumeroBase = gerarNumeroBase
	mpGerarCodigoRegiaoFiscal = func(uf string) (int, error) { return -1, ErrGeracaoCPF }

	_, err := gerarCPFParaUF("mg")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarCodigoRegiaoFiscal = gerarCodigoRegiaoFiscal
}

func TestNewCPF(t *testing.T) {
	cpf := newCPF(1234)
	assert.Equal(t, CPF("00000123431"), cpf)

	cpf = newCPF(12345678900)
	assert.Equal(t, CPF("12345678906"), cpf)
}

func TestGerarNumeroBaseErro(t *testing.T) {
	mpGenRandomDecimalUnit = func() (int, error) { return -1, fmt.Errorf("any error") }

	_, err := gerarNumeroBase()
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGenRandomDecimalUnit = genRandomDecimalUnit
}

func TestGerarCodigoRegiaoFiscalErro(t *testing.T) {
	mpGenRandomDecimalUnit = func() (int, error) { return -1, fmt.Errorf("any error") }

	_, err := gerarCodigoRegiaoFiscal("")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGenRandomDecimalUnit = genRandomDecimalUnit
}

func TestGerarDigitosVerificadores(t *testing.T) {
	dv1, dv2 := gerarDigitosVerificadores([]int{4, 9, 9, 9, 9, 9, 9, 9}, 6)
	assert.Equal(t, 0, dv1)
	assert.Equal(t, 8, dv2)

	dv1, dv2 = gerarDigitosVerificadores([]int{1, 2, 3, 4, 5, 6, 7, 1}, 6)
	assert.Equal(t, 4, dv1)
	assert.Equal(t, 0, dv2)
}

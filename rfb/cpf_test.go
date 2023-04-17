package rfb

import (
	"fmt"
	"testing"

	"github.com/aureliano/go-brutils/number"
	"github.com/stretchr/testify/assert"
)

func TestGerarCPF(t *testing.T) {
	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal

	cpf, _ := gerarCPF()
	assert.Regexp(t, cpfNumeralRegex, cpf)
}

func TestGerarCPFParaUF(t *testing.T) {
	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal

	cpf, _ := gerarCPFParaUF("")
	assert.Regexp(t, cpfNumeralRegex, cpf)

	cpf, _ = gerarCPFParaUF("mg")
	assert.Regexp(t, cpfNumeralRegex, cpf)

	_, err := gerarCPFParaUF("br")
	assert.ErrorIs(t, err, ErrUFDesconhecida)
}

func TestGerarCPFParaUFErroGeracaoNumBase(t *testing.T) {
	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal
	mpGerarNumeroBaseCPF = func() ([]int, error) { return nil, ErrGeracaoCPF }

	_, err := gerarCPFParaUF("mg")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarNumeroBaseCPF = gerarNumeroBaseCPF
}

func TestGerarCPFParaUFErroCodRecFiscal(t *testing.T) {
	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal
	mpGerarNumeroBaseCPF = gerarNumeroBaseCPF
	mpGerarCodigoRegiaoFiscal = func(uf string) (int, error) { return -1, ErrGeracaoCPF }

	_, err := gerarCPFParaUF("mg")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarCodigoRegiaoFiscal = gerarCodigoRegiaoFiscal
}

func TestNewCPF(t *testing.T) {
	type testCase struct {
		name     string
		input    uint
		expected CPF
	}
	testCases := []testCase{
		{name: "AC - complete CPF", input: 939764642, expected: "93976464279"},
		{name: "AC - complete CPF", input: 42646692, expected: "04264669260"},
		{name: "AL - complete CPF", input: 759190674, expected: "75919067470"},
		{name: "AL - complete CPF", input: 108138284, expected: "10813828481"},
		{name: "AM - complete CPF", input: 676657132, expected: "67665713220"},
		{name: "AM - complete CPF", input: 525859542, expected: "52585954200"},
		{name: "AP - complete CPF", input: 610055082, expected: "61005508224"},
		{name: "AP - complete CPF", input: 75244372, expected: "07524437200"},
		{name: "BA - complete CPF", input: 977904285, expected: "97790428578"},
		{name: "BA - complete CPF", input: 873918735, expected: "87391873578"},
		{name: "CE - complete CPF", input: 840604033, expected: "84060403350"},
		{name: "CE - complete CPF", input: 877372843, expected: "87737284389"},
		{name: "DF - complete CPF", input: 849199221, expected: "84919922191"},
		{name: "DF - complete CPF", input: 1981621, expected: "00198162197"},
		{name: "ES - complete CPF", input: 224026487, expected: "22402648708"},
		{name: "ES - complete CPF", input: 659271467, expected: "65927146708"},
		{name: "GO - complete CPF", input: 908386811, expected: "90838681174"},
		{name: "GO - complete CPF", input: 766183131, expected: "76618313171"},
		{name: "MA - complete CPF", input: 127675473, expected: "12767547367"},
		{name: "MA - complete CPF", input: 771728933, expected: "77172893344"},
		{name: "MG - complete CPF", input: 768852376, expected: "76885237612"},
		{name: "MG - complete CPF", input: 466728576, expected: "46672857613"},
		{name: "MS - complete CPF", input: 919104901, expected: "91910490164"},
		{name: "MS - complete CPF", input: 500871141, expected: "50087114151"},
		{name: "MT - complete CPF", input: 349423791, expected: "34942379130"},
		{name: "MT - complete CPF", input: 191837891, expected: "19183789162"},
		{name: "PA - complete CPF", input: 699825174, expected: "69982517422"},
		{name: "PA - complete CPF", input: 21145414, expected: "02114541410"},
		{name: "PB - complete CPF", input: 49500774, expected: "04950077490"},
		{name: "PB - complete CPF", input: 770979574, expected: "77097957492"},
		{name: "PE - complete CPF", input: 28062704, expected: "02806270413"},
		{name: "PE - complete CPF", input: 709154264, expected: "70915426498"},
		{name: "PI - complete CPF", input: 575332883, expected: "57533288300"},
		{name: "PI - complete CPF", input: 129342293, expected: "12934229327"},
		{name: "PR - complete CPF", input: 205308099, expected: "20530809907"},
		{name: "PR - complete CPF", input: 612897719, expected: "61289771979"},
		{name: "RJ - complete CPF", input: 15230717, expected: "01523071761"},
		{name: "RJ - complete CPF", input: 891013227, expected: "89101322729"},
		{name: "RN - complete CPF", input: 164983864, expected: "16498386460"},
		{name: "RN - complete CPF", input: 84748444, expected: "08474844401"},
		{name: "RS - complete CPF", input: 721950280, expected: "72195028009"},
		{name: "RS - complete CPF", input: 946309750, expected: "94630975039"},
		{name: "RO - complete CPF", input: 670486362, expected: "67048636211"},
		{name: "RO - complete CPF", input: 429055532, expected: "42905553227"},
		{name: "RR - complete CPF", input: 481043172, expected: "48104317202"},
		{name: "RR - complete CPF", input: 597758502, expected: "59775850207"},
		{name: "SC - complete CPF", input: 916413999, expected: "91641399996"},
		{name: "SC - complete CPF", input: 174139309, expected: "17413930943"},
		{name: "SE - complete CPF", input: 357683065, expected: "35768306501"},
		{name: "SE - complete CPF", input: 583252955, expected: "58325295538"},
		{name: "SP - complete CPF", input: 516498188, expected: "51649818823"},
		{name: "SP - complete CPF", input: 105557108, expected: "10555710866"},
		{name: "TO - complete CPF", input: 606581011, expected: "60658101188"},
		{name: "TO - complete CPF", input: 925461721, expected: "92546172107"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := newCPF(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewCPFFromStr(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected CPF
	}
	testCases := []testCase{
		{name: "CPF sem formatação", input: "93976464279", expected: "93976464279"},
		{name: "CPF com formatação", input: "042.646.692-60", expected: "04264669260"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := newCPFFromStr(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewCPFFromStrErro(t *testing.T) {
	_, err := newCPFFromStr("123")
	assert.ErrorIs(t, err, ErrCPFInvalido)
}

func TestCpfValido(t *testing.T) {
	assert.False(t, cpfValido("123"))
	assert.False(t, cpfValido("123987987988878"))
	assert.False(t, cpfValido("12345678901"))
	assert.False(t, cpfValido("78545214665"))
	assert.False(t, cpfValido("00475223690"))
	assert.False(t, cpfValido("87554128963"))
	assert.False(t, cpfValido("11470025308"))

	assert.True(t, cpfValido("92546172107"))
	assert.True(t, cpfValido("60658101188"))
	assert.True(t, cpfValido("10555710866"))
	assert.True(t, cpfValido("51649818823"))
	assert.True(t, cpfValido("58325295538"))
}

func TestFormatarCPF(t *testing.T) {
	fmtd := formatarCPF("92546172107")
	assert.Equal(t, "925.461.721-07", fmtd)

	fmtd = formatarCPF("00475223690")
	assert.Equal(t, "004.752.236-90", fmtd)

	fmtd = formatarCPF("12")
	assert.Equal(t, "12", fmtd)

	fmtd = formatarCPF("123456789123")
	assert.Equal(t, "123456789123", fmtd)
}

func TestDesformatarCPF(t *testing.T) {
	cpf := desformatarCPF("925.461.721-07")
	assert.Equal(t, "92546172107", cpf)

	cpf = desformatarCPF("004.752.236-90")
	assert.Equal(t, "00475223690", cpf)

	cpf = desformatarCPF("12")
	assert.Equal(t, "12", cpf)

	cpf = desformatarCPF("123456789123")
	assert.Equal(t, "123456789123", cpf)
}

func TestGerarNumeroBaseCPFErro(t *testing.T) {
	mpGerarUnidadeDecimalCPF = func() (int, error) { return -1, fmt.Errorf("any error") }

	_, err := gerarNumeroBaseCPF()
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal
}

func TestGerarCodigoRegiaoFiscalErro(t *testing.T) {
	mpGerarUnidadeDecimalCPF = func() (int, error) { return -1, fmt.Errorf("any error") }

	_, err := gerarCodigoRegiaoFiscal("")
	assert.ErrorIs(t, err, ErrGeracaoCPF)

	mpGerarUnidadeDecimalCPF = number.GerarUnidadeDecimal
}

func TestGerarDigitosVerificadoresCPF(t *testing.T) {
	dv1, dv2 := gerarDigitosVerificadoresCPF([]int{4, 9, 9, 9, 9, 9, 9, 9}, 6)
	assert.Equal(t, 0, dv1)
	assert.Equal(t, 3, dv2)

	dv1, dv2 = gerarDigitosVerificadoresCPF([]int{1, 2, 3, 4, 5, 6, 7, 1}, 6)
	assert.Equal(t, 4, dv1)
	assert.Equal(t, 5, dv2)
}

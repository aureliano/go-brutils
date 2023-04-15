package rfb

import "strings"

type Estado struct {
	UF           string
	Nome         string
	RegiaoFiscal int
}

const (
	rf1  = 1
	rf2  = 2
	rf3  = 3
	rf4  = 4
	rf5  = 5
	rf6  = 6
	rf7  = 7
	rf8  = 8
	rf9  = 9
	rf10 = 10
)

var estados = []Estado{
	{
		UF:           "DF",
		Nome:         "Distrito Federal",
		RegiaoFiscal: rf1,
	},
	{
		UF:           "GO",
		Nome:         "Goiás",
		RegiaoFiscal: rf1,
	},
	{
		UF:           "MS",
		Nome:         "Mato Grosso do Sul",
		RegiaoFiscal: rf1,
	},
	{
		UF:           "MT",
		Nome:         "Mato Grosso",
		RegiaoFiscal: rf1,
	},
	{
		UF:           "TO",
		Nome:         "Tocantins",
		RegiaoFiscal: rf1,
	},
	{
		UF:           "AC",
		Nome:         "Acre",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "AM",
		Nome:         "Amazonas",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "AP",
		Nome:         "Amapá",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "PA",
		Nome:         "Pará",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "RO",
		Nome:         "Rondônia",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "RR",
		Nome:         "Roraima",
		RegiaoFiscal: rf2,
	},
	{
		UF:           "CE",
		Nome:         "Ceará",
		RegiaoFiscal: rf3,
	},
	{
		UF:           "MA",
		Nome:         "Maranhão",
		RegiaoFiscal: rf3,
	},
	{
		UF:           "PI",
		Nome:         "Piauí",
		RegiaoFiscal: rf3,
	},
	{
		UF:           "AL",
		Nome:         "Alagoas",
		RegiaoFiscal: rf4,
	},
	{
		UF:           "PB",
		Nome:         "Paraíba",
		RegiaoFiscal: rf4,
	},
	{
		UF:           "PE",
		Nome:         "Pernambuco",
		RegiaoFiscal: rf4,
	},
	{
		UF:           "RN",
		Nome:         "Rio Grande do Norte",
		RegiaoFiscal: rf4,
	},
	{
		UF:           "BA",
		Nome:         "Bahia",
		RegiaoFiscal: rf5,
	},
	{
		UF:           "SE",
		Nome:         "Sergipe",
		RegiaoFiscal: rf5,
	},
	{
		UF:           "MG",
		Nome:         "Minas Gerais",
		RegiaoFiscal: rf6,
	},
	{
		UF:           "ES",
		Nome:         "Espírito Santo",
		RegiaoFiscal: rf7,
	},
	{
		UF:           "RJ",
		Nome:         "Rio de Janeiro",
		RegiaoFiscal: rf7,
	},
	{
		UF:           "SP",
		Nome:         "São Paulo",
		RegiaoFiscal: rf8,
	},
	{
		UF:           "PR",
		Nome:         "Paraná",
		RegiaoFiscal: rf9,
	},
	{
		UF:           "SC",
		Nome:         "Santa Catarina",
		RegiaoFiscal: rf9,
	},
	{
		UF:           "RS",
		Nome:         "Rio Grande do Sul",
		RegiaoFiscal: rf10,
	},
}

func newEstado(uf string) *Estado {
	chave := strings.ToUpper(uf)
	for _, estado := range estados {
		if chave == estado.UF {
			return &estado
		}
	}

	return nil
}

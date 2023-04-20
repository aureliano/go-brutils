/*
No pacote rfb é onde se encontram as funcionalidades para manipulação de números identificadores
da Receita Federal do Brasil. Tais como CPF e CNPJ.

Cada tipo de número identificador implementa a interface NIRFB, que possui os métodos para tratamento
dos respectivos NIs.

# NI-RFB

Interface para manipulação de números identificadores da RFB.

	type NIRFB interface {
		NumeroBase() uint
		DigitosVerificadores() (int, int)
		Valido() bool
		Formatado() string
		Desformatado() string
	}

# Iniciação dos NIs

Para cada NI há um tipo personalizado. A iniciação desses tipos deve ser feita através das funções adequadas.
Há duas formas de se obter um NI: passando o número base (um inteiro não negativo) ou passando o número completo
como string com ou sem máscara.

	// Passando o número base do CPF "000.001.234-39".
	cpf := rfb.NewCPF(1234)

	// Passando o número do CPF como texto sem máscara.
	cpf, err := rfb.NewCPFFromStr("00000123439")

	// Passando o número do CPF como texto com máscara.
	cpf, err := rfb.NewCPFFromStr("000.001.234-39")

Para a iniciação do NI como string, deve-se atentar à possibilidade de falha.

	// Falhará se o número tiver o total de dígitos diferente de onze (11).
	cpf, err := rfb.NewCPFFromStr("000123439")
	if errors.Is(err, ErrCPFInvalido) {
		// Faz alguma coisa.
	}

	// Falhará se a máscara não seguir o padrão "###.###.###-##".
	cpf, err := rfb.NewCPFFromStr("000001.234-39")
	if errors.Is(err, ErrCPFInvalido) {
		// Faz alguma coisa.
	}

# Geração de NIs

Existe também suporte a geração de NIs. Para CPF também é possível gerar pra um Estado/Região Fiscal.

	// Geração de CPF de uma região fiscal tomada aleatoriamente.
	cpf, _ := GerarCPF()

	// Geração de CPF para uma região fiscal definida.
	uf := NewEstado("mg")
	cpf, _ := GerarCPFParaUF(uf)

	// Falha na geração de CPF.
	uf := NewEstado("br")
	cpf, err := GerarCPFParaUF(uf)

	if errors.Is(err, ErrUFDesconhecida) {
		// Faz alguma coisa.
	}

# CPF - Implementação

	type CPF string

	cpf := rfb.NewCPF(1234)

	cpf.NumeroBase()           // 1234
	cpf.DigitosVerificadores() // [3, 9]
	cpf.Valido()               // true
	cpf.Formatado()            // 000.001.234-39
	cpf.Desformatado()         // 00000123439

	cpf, _ = rfb.NewCPFFromStr("00000123439") // Poderia ser 000.001.234-39

	cpf.NumeroBase()           // 1234
	cpf.DigitosVerificadores() // [3, 9]
	cpf.Valido()               // true
	cpf.Formatado()            // 000.001.234-39
	cpf.Desformatado()         // 00000123439

# CNPJ - Implementação

	type CNPJ string

	cnpj := rfb.NewCNPJ(1234)

	cnpj.NumeroBase()           // 1234
	cnpj.DigitosVerificadores() // [3, 9]
	cnpj.Valido()               // true
	cnpj.Formatado()            // 00.000.000/1234-39
	cnpj.Desformatado()         // 00000000123439

	cnpj, _ = rfb.NewCPFFromStr("00000000123439") // Poderia ser 00.000.000/1234-39

	cnpj.NumeroBase()           // 1234
	cnpj.DigitosVerificadores() // [3, 9]
	cnpj.Valido()               // true
	cnpj.Formatado()            // 00.000.000/1234-39
	cnpj.Desformatado()         // 00000000123439
*/
package rfb

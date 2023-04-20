/*
O pacote number é um pacote com funções utilitárias para manipulação de números.

# ExtrairNumeros

Função para remoção de todo caractere que não for um número decimal da cadeia de texto.
Pode ser usado pra remover máscaras de números cadastrais como CPF e CNPJ.

	numeros := number.ExtrairNumeros("123.456.789-01")
	fmt.Println(numeros) // Saída: 12345678901

	numeros = number.ExtrairNumeros("12.345.678/0001-23")
	fmt.Println(numeros) // Saída: 12345678000123

# GerarUnidadeDecimal

Função para geração aleatória de um número decimal de zero a nove.

	ud := GerarUnidadeDecimal()
	fmt.Println(ud) // Saída: um número inteiro de zero a nove (0-9).
*/
package number

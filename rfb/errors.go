package rfb

import "errors"

var (
	ErrUFDesconhecida = errors.New("uf desconhecida")
	ErrGeracaoCPF     = errors.New("erro na geração do cpf")
	ErrCPFInvalido    = errors.New("cpf inválido (deve ser número - com ou sem máscara - de onze dígitos)")
	ErrGeracaoCNPJ    = errors.New("erro na geração do cnpj")
	ErrCNPJInvalido   = errors.New("cnpj inválido (deve ser número - com ou sem máscara - de quatorze dígitos)")
)

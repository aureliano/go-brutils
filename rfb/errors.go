package rfb

import "errors"

var (
	ErrUFDesconhecida = errors.New("uf desconhecida")
	ErrGeracaoCPF     = errors.New("erro na geração do cpf")
)

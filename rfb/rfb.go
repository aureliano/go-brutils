package rfb

func NewEstado(uf string) *Estado {
	e := newEstado(uf)
	if e == nil {
		e = &Estado{}
	}

	return e
}

func GerarCPF() (CPF, error) {
	return gerarCPF()
}

func GerarCPFParaUF(uf string) (CPF, error) {
	return gerarCPFParaUF(uf)
}

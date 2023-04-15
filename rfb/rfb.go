package rfb

func NewEstado(uf string) *Estado {
	e := newEstado(uf)
	if e == nil {
		e = &Estado{}
	}

	return e
}

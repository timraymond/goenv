package mocks

type Builder struct {
	MkdirAllF func(string) error
	ChdirF    func(string) error
}

func (b *Builder) MkdirAll(path string) error {
	return b.MkdirAllF(path)
}

func (b *Builder) Chdir(path string) error {
	return b.ChdirF(path)
}

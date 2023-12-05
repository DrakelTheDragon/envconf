package internal

type Var[T any] struct {
	dest   *T
	parser Parser[T]
}

func NewVar[T any](dest *T, parser ParseFunc[T]) Var[T] { return Var[T]{dest: dest, parser: parser} }

func (v Var[T]) Parse(name string) error { return ParseVar(name, NewVarParser(v.dest, v.parser)) }

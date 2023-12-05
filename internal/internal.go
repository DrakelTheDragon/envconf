package internal

import (
	"fmt"
	"os"
)

type VarParser interface {
	ParseVar(string) error
}

type VarParserFunc func(string) error

func (f VarParserFunc) ParseVar(s string) error { return f(s) }

func GetVar(name string) (string, error) {
	val, exists := os.LookupEnv(name)
	if !exists {
		return "", fmt.Errorf("environment variable %s not set", name)
	}
	return val, nil
}

func ParseVar(name string, parser VarParser) error {
	val, err := GetVar(name)
	if err != nil {
		return err
	}
	return parser.ParseVar(val)
}

func NewVarParser[T any](dest *T, f Parser[T]) VarParserFunc {
	return func(s string) error {
		val, err := f.Parse(s)
		if err != nil {
			return err
		}
		*dest = val
		return nil
	}
}

func BoolVarParser(dest *bool) VarParserFunc { return NewVarParser(dest, ParseFunc[bool](ParseBool)) }
func IntVarParser(dest *int) VarParserFunc   { return NewVarParser(dest, ParseFunc[int](ParseInt)) }

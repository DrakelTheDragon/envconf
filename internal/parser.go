package internal

import "strconv"

type Parser[T any] interface {
	Parse(string) (T, error)
}

type ParseFunc[T any] func(string) (T, error)

func (f ParseFunc[T]) Parse(s string) (T, error) { return f(s) }

func Parse[T any](s string, f ParseFunc[T]) (T, error) { return f(s) }

func ParseBool(s string) (bool, error) {
	val, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return val, nil
}

func ParseInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return val, nil
}

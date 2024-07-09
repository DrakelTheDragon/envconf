package envconf

import "go.drakelthedragon.dev/envconf/internal"

type StringVar struct{ internal.Var[string] }

func NewStringVar(dest *string) *StringVar {
	return &StringVar{Var: internal.NewVar(dest, func(s string) (string, error) { return s, nil })}
}

type BoolVar struct{ internal.Var[bool] }

func NewBoolVar(dest *bool) *BoolVar {
	return &BoolVar{Var: internal.NewVar[bool](dest, internal.ParseBool)}
}

type IntVar struct{ internal.Var[int] }

func NewIntVar(dest *int) *IntVar {
	return &IntVar{Var: internal.NewVar(dest, internal.ParseInt)}
}

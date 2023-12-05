package envconf

type VarParser interface {
	ParseVar(string) error
}

type Vars map[string]VarParser

func (v Vars) Parse() error {
	for name, parser := range v {
		if err := parser.ParseVar(name); err != nil {
			return err
		}
	}
	return nil
}

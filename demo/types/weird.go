package types

type Person struct {
	Name string
	ssn  string
}

func (p *Person) GetName() string {
	return p.Name
}

type HasName interface {
	GetName() string
}

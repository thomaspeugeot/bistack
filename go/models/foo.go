package models

type Foo struct {
	Name string
}

func (foo *Foo) SetName(name string) {
	foo.Name = name
}

package models

type Bar struct {
	Name string
}

func (bar *Bar) SetName(name string) {
	bar.Name = name
}

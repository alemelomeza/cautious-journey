package domain

type Parent struct {
	Name string
	Relationship string
}

func (p *Parent) Validate() bool {
	if p.Relationship == "mother" || p.Relationship == "father" {
		return true
	}
	return false
}
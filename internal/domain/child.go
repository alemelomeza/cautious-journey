package domain

type Child struct {
	Name    string
	Age     int
	Parents []Parent

}

func (c *Child) ValidateAge() bool {
	if c.Age > 0 && c.Age <= 5 {
		return true
	}
	return false
}

func (c *Child) Validate() bool {
	if len(c.Parents) != 0 && c.ValidateAge() {
		return true
	}
	return false
}

package domain

import "testing"

func TestValidateAge(t *testing.T) {
	t.Run("should be true when the child have 3 years old", func(t *testing.T) {
		c := Child{
			Name: "Nelson",
			Age:  3,
		}
		want := true
		got := c.ValidateAge()
		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("Shoould be False when Child have over 3 years old", func(t *testing.T) {
		c := Child{
			Name: "Bryan",
			Age:  29,
		}
		want := false
		got := c.ValidateAge()
		if want != got {
			t.Errorf("want: %v , got %v", want, got)
		}
	})
}

func TestValidate(t *testing.T) {
	t.Run("Happy PAth: validate return true when age and Parents is complete ", func(t *testing.T) {
		parents := []Parent{
			Parent{Name: "Nelson", Relationship: "father"},
			Parent{Name: "Andres", Relationship: "mother"},
		}
		c := Child{
			Name:    "Bryan Sepulveda",
			Age:     4,
			Parents: parents,
		}
		want := true
		got := c.Validate()
		if want != got {
			t.Errorf("want: %v. got: %v", want, got)
		}
	})

	t.Run("Unhappy Path: return false when age and. Parents is not valid", func(t *testing.T) {
		p := []Parent{}
		c := Child{Name: "Alejo", Age: 40, Parents: p}
		want := false
		got := c.Validate()

		if want != got {
			t.Errorf("want %v, got: %v", want, got)
		}

	})
}

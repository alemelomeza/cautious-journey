package domain

import (
	"testing"
)

func TestParentValidate(t *testing.T) {

	t.Run("sould be return false when parent relataionship is Capital", func(t *testing.T) {
		want := false
		parent := Parent{Name: "Bryan", Relationship: "Mother"}
		got := parent.Validate()
		if want != got {
			t.Errorf("Want: %v. got:%v  ", want, got)
		}

	})

	t.Run("sould be return true when parent relataionship ", func(t *testing.T) {
		want := true
		parent := Parent{Name: "Bryan", Relationship: "mother"}
		got := parent.Validate()
		if want != got {
			t.Errorf("Want: %v. got:%v  ", want, got)
		}

	})

}

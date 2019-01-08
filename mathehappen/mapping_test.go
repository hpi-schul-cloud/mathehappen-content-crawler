package mathehappen

import (
	"testing"
)

func TestFilterSpaceBeforeColon(t *testing.T) {
	data := "Bruchzahlen : Bruchdarstellungen"
	expected := "Bruchzahlen: Bruchdarstellungen"

	actual := filterSpaceBeforeColon(data)
	if expected != actual {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

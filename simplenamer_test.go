package namecon_test

import (
	"fmt"
	"testing"

	"github.com/honeytrap/namecon"
)

// TestSimpleNamer validates the use of the provided namer to match the
// giving template rule set provided.
func TestSimpleNamer(t *testing.T) {
	namer := namecon.GenerateNamer(namecon.SimpleNamer{}, "API-%s-%s")

	firstName := namer("Trappa")
	secondName := namer("Honey")

	fmt.Printf("First: %q -> %q ", firstName, secondName)

	if firstName == secondName {
		t.Fatalf("Should have successfully generated new unique generation names")
	}
	t.Logf("Should have successfully generated new unique generation names")
}

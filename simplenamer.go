package namecon

import "fmt"

var (
	maxCharacter      = 60
	maxBaseCharacter  = 10
	maxAvailableSpace = maxCharacter - maxBaseCharacter
)

// SimpleNamer defines a struct which implements the NameGenerator interface and generates
// a giving name based on the template and a giving set of directives.
type SimpleNamer struct{}

// Generate returns a new name from the provided template and base to generate based
// on the rules of the simple rules.
func (s SimpleNamer) Generate(template string, base string) string {
	if len(base) > maxBaseCharacter {
		base = base[:maxBaseCharacter]
	}

	gened := fmt.Sprintf(template, base, String(maxBaseCharacter))

	if len(gened) > maxCharacter {
		gened = gened[:maxCharacter]
	}

	return gened
}

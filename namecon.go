// Package namecon defines a package for naming things, it provides us a baseline
// package for providing naming standards for things we use.
package namecon

// NameGenerator defines a type whch exposes a method which generates a new name
// based on the provided parameters.
type NameGenerator interface {
	Generate(template string, base string) string
}

// GenerateNamer defines a function which accepts a base generator and uses the provided
// template and returns a function which will use this template and generator for all
// name generation.
func GenerateNamer(gen NameGenerator, template string) func(string) string {
	return func(base string) string {
		return gen.Generate(template, base)
	}
}

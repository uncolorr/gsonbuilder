package builder

type class struct {
	name       string
	properties []property
}

// Set name and set first character to uppercase
func (class *class) setNameWithFormat(name string) {
	class.name = upperCaseFirst(name)
}

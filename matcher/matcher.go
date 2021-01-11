package matcher

// Matcher provides the structure for matcher operations.
type Matcher struct {
	Describe string
	matches  func(actual interface{}) bool
}

// Match is matching wrapper function.
func (m *Matcher) Match(actual interface{}) bool {
	return m.matches(actual)
}

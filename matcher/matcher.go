package matcher

import "testing"

// Matcher provides the structure for matcher operations.
type Matcher struct {
	Describe string
	matches  func(t *testing.T, actual interface{}) bool
}

// Match is matching wrapper function.
func (m *Matcher) Match(t *testing.T, actual interface{}) bool {
	return m.matches(t, actual)
}

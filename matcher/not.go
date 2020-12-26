package matcher

//Not negates the given matcher.
func Not(matcher *Matcher) *Matcher {
	m := new(Matcher)
	m.Describe = "not(" + matcher.Describe + ")"
	m.matches = func(actual interface{}) bool {
		return !matcher.matches(actual)
	}
	return m
}

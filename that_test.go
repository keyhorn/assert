package assert

import (
	"testing"

	"github.com/keyhorn/assert/matcher"
)

func TestThatIsEqualTo(t *testing.T) {
	That(t, "abcde", matcher.EqualTo("abcde"))
}

func TestThatIsNotEqualTo(t *testing.T) {
	That(t, "abcde", matcher.Not(matcher.EqualTo("ABCDE")))
}

func TestThatIsNil(t *testing.T) {
	That(t, nil, matcher.Nil())
}

func TestThatIsNotNil(t *testing.T) {
	That(t, "abcde", matcher.Not(matcher.Nil()))
}

func TestThatIsEmpty(t *testing.T) {
	That(t, "", matcher.Empty())
}

func TestThatIsNotEmpty(t *testing.T) {
	That(t, "abcde", matcher.Not(matcher.Empty()))
}

func TestThatIsLessThan(t *testing.T) {
	That(t, 99, matcher.LessThan(100))
}

func TestThatIsLessThanOrEqualTo(t *testing.T) {
	That(t, 99, matcher.LessThanOrEqualTo(100))
	That(t, 100, matcher.LessThanOrEqualTo(100))
}

func TestThatIsGreaterThan(t *testing.T) {
	That(t, 101, matcher.GreaterThan(100))
}

func TestThatIsGreaterThanOrEqualTo(t *testing.T) {
	That(t, 100, matcher.GreaterThanOrEqualTo(100))
	That(t, 101, matcher.GreaterThanOrEqualTo(100))
}

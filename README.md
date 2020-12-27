# assert

Assertion library for Go

## Install

Use `go get` to install this library.

```shell
go get -u github.com/keyhorn/assert
```

## API Document

See [GoDoc](https://godoc.org/github.com/keyhorn/assert) for full doument.

## Usage

```golang
import (
  "testing"

  "github.com/keyhorn/assert"
  "github.com/keyhorn/assert/matcher"
)

func TestExample(t *testing.T) {
  assert.That(t, "Lorem ipsum", matcher.EqualTo("Lorem ipsum"))
}
```

## Matcher list

| Marcher              | Description                                                                                               |
| -------------------- | --------------------------------------------------------------------------------------------------------- |
| AnyOf                | Takes some matchers and checks if at least one of the matchers return true.                               |
| AllOf                | Takes some matchers and checks if at least one of the matchers return true.                               |
| Not                  | Negates the given matcher.                                                                                |
| EqualTo              | Checks if two values are equal.                                                                           |
| Nil                  | Matches if the actual value is nil.                                                                       |
| Empty                | Empty matches if the actual is "empty".                                                                   |
| LessThan             | Compares two values that are numeric or string values, and when called returns true if actual < expected. |
| LessThanOrEqualTo    | A short hand matcher for anyOf(LessThan(x), equalTo(x))                                                   |
| GreaterThan          | Compares two values that are numeric or string values, and when called returns true if actual > expected. |
| GreaterThanOrEqualTo | A short hand matcher for anyOf(GreaterThan(x), EqualTo(x))                                                |
| StartsWith           | Returns a matcher that matches if the given string is prefixed with the expected string.                  |
| EndsWith             | Returns a matcher that matches if the given string is suffixed with the expected string.                  |
| Key                  | checks if actual has a key == expected.                                                                   |
| AllKeys              | checks if map actual has all keys == expecteds.                                                           |

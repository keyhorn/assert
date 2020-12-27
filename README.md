# assert

[![Test Status](https://github.com/keyhorn/assert/workflows/test/badge.svg)](https://github.com/keyhorn/assert/actions?query=workflow%3Atest)
[![Lint Status](https://github.com/keyhorn/assert/workflows/lint/badge.svg)](https://github.com/keyhorn/assert/actions?query=workflow%3Alint)
[![Coverage Status](https://coveralls.io/repos/github/keyhorn/assert/badge.svg?branch=main)](https://coveralls.io/github/keyhorn/assert?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/keyhorn/assert)](https://goreportcard.com/report/github.com/keyhorn/assert)
[![BCH compliance](https://bettercodehub.com/edge/badge/keyhorn/assert?branch=main)](https://bettercodehub.com/)
[![Documentation](https://godoc.org/github.com/keyhorn/assert?status.svg)](http://godoc.org/github.com/keyhorn/assert)
[![License](https://img.shields.io/github/license/keyhorn/assert.svg?maxAge=2592000)](https://github.com/keyhorn/assert/LICENSE)

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
| AllOf                | Takes some matchers and checks if all of the matchers return true.                               |
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

## License

This library is licensed under MIT license. See [LICENSE](https://github.com/keyhorn/assert/LICENSE) for details.
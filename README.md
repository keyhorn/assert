# assert

[![Test Status](https://github.com/keyhorn/assert/workflows/test/badge.svg)](https://github.com/keyhorn/assert/actions?query=workflow%3Atest)
[![Lint Status](https://github.com/keyhorn/assert/workflows/lint/badge.svg)](https://github.com/keyhorn/assert/actions?query=workflow%3Alint)
[![Coverage Status](https://coveralls.io/repos/github/keyhorn/assert/badge.svg?branch=main)](https://coveralls.io/github/keyhorn/assert?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/keyhorn/assert)](https://goreportcard.com/report/github.com/keyhorn/assert)
[![BCH compliance](https://bettercodehub.com/edge/badge/keyhorn/assert?branch=main)](https://bettercodehub.com/)
[![Documentation](https://godoc.org/github.com/keyhorn/assert?status.svg)](http://godoc.org/github.com/keyhorn/assert)
[![License](https://img.shields.io/github/license/keyhorn/assert.svg?maxAge=2592000)](https://github.com/keyhorn/assert/LICENSE)
[![Release](https://img.shields.io/github/release/keyhorn/assert.svg?label=Release)](https://github.com/keyhorn/assert/releases)

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

### Basic matchers

| Marcher | Description                                                                 |
| ------- | --------------------------------------------------------------------------- |
| AnyOf   | Takes some matchers and checks if at least one of the matchers return true. |
| AllOf   | Takes some matchers and checks if all of the matchers return true.          |
| Not     | Negates the given matcher.                                                  |
| EqualTo | Returns true if two values are equal.                                       |
| Nil     | Returns true if the actual value is nil.                                    |
| Empty   | Returns true if the actual is "empty".                                      |

### Matchers for string

| Marcher    | Description                                                                                                          |
| ---------- | -------------------------------------------------------------------------------------------------------------------- |
| StartsWith | Compares two values that are string values, and returns true if the given string is starts with the expected string. |
| EndsWith   | Compares two values that are string values, and returns true if the given string is ends with the expected string.   |
| Contains   | Returns true if the argument is a string that contains a specific substring.                                         |

### Matchers for numerical value

| Marcher              | Description                                                                                   |
| -------------------- | --------------------------------------------------------------------------------------------- |
| LessThan             | Compares two values that are numeric or string values, and returns true if actual < expected. |
| LessThanOrEqualTo    | A short hand matcher for AnyOf(LessThan(x), EqualTo(x))                                       |
| GreaterThan          | Compares two values that are numeric or string values, and returns true if actual > expected. |
| GreaterThanOrEqualTo | A short hand matcher for AnyOf(GreaterThan(x), EqualTo(x))                                    |

### Matchers for list(array and slice)

| Marcher            | Description                                                                     |
| ------------------ | ------------------------------------------------------------------------------- |
| ContainsIn         | Returns true if all expected values contains in actual list at specified order. |
| ContainsInAnyOrder | Returns true if all expected values contains in actual list at any order.       |

### Matchers for map

| Marcher  | Description                                           |
| -------- | ----------------------------------------------------- |
| HasEntry | Returns true if actual has entry with a conditions.   |
| HasKey   | Returns true if actual has a key with a conditions.   |
| HasValue | Returns true if actual has a value with a conditions. |

## License

This library is licensed under MIT license. See [LICENSE](https://github.com/keyhorn/assert/LICENSE) for details.

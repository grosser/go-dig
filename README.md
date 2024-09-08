# go-dig

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/mnogu/go-dig)](https://pkg.go.dev/mod/github.com/mnogu/go-dig)
[![GitHub Actions](https://github.com/mnogu/go-dig/workflows/Go/badge.svg)](https://github.com/mnogu/go-dig/actions?query=workflow%3AGo)

Access values in a deeply nested maps or slices.

(Go version of [`Hash#dig`](https://docs.ruby-lang.org/en/2.7.0/Hash.html#method-i-dig) and [`Array#dig`](https://docs.ruby-lang.org/en/2.7.0/Array.html#method-i-dig) in Ruby)

## Install

```
$ go get -u github.com/mnogu/go-dig
```

## Examples

```
value, err := dig.Dig(nested, "foo", "bar", "baz") // get value out of {foo: {bar: {baz: value}}}
value, err := dig.Dig(nested, "foo", 0, "baz")     // get value out of {foo: [{baz: value}]}
```

### Lookup in nested maps

[Go playground](https://go.dev/play/p/c1t82Gfmice)

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mnogu/go-dig"
)

func main() {
	var jsonBlob = []byte(`{"foo": {"bar": {"baz": 1}}}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		fmt.Println(err)
	}
	
	// successful lookup
	value, err := dig.Dig(v, "foo", "bar", "baz")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("foo.bar.baz =", value) // foo.bar.baz = 1

	// failed lookup
	value, err = dig.Dig(v, "foo", "qux", "quux")
	if err != nil {
		fmt.Println(err) // key qux not found in <nil>
	}
	fmt.Println("foo.qux.quux =", value) // foo.qux.quux = <nil>
}
```

### Lookup in nested slices

[Go Playground](https://go.dev/play/p/79CNgEoX6v-)

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mnogu/go-dig"
)

func main() {
	var jsonBlob = []byte(`{"foo": [10, 11, 12]}`)
	var v interface{}
	if err := json.Unmarshal(jsonBlob, &v); err != nil {
		fmt.Println(err)
	}

	// successful lookup
	value, err := dig.Dig(v, "foo", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("foo.1 =", value) // foo.1 = 11

	// failed lookup slice style in map
	value, err = dig.Dig(v, "foo", 1, 0)
	if err != nil {
		fmt.Println(err) // 11 isn't a slice
	}
	fmt.Println("foo.1.0 =", value) // foo.1.0 = <nil>

	// failed lookup map style in slice
	value, err = dig.Dig(v, "foo", "bar")
	if err != nil {
		fmt.Println(err) // [10 11 12] isn't a map
	}
	fmt.Println("foo.bar =", value) // foo.bar = <nil>
}
```

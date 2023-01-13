# Random name and slug generator

coolname is a Golang port of a [Python package][upstream] by [Alexander Lukanin].
All hard work of creating and maintaing word lists happens upstream.

[upstream]: https://github.com/alexanderlukanin13/coolname
[Alexander Lukanin]: https://github.com/alexanderlukanin13


## Documentation

See [package docs at pkg.go.dev][docs]

[docs]: https://pkg.go.dev/github.com/sio/coolname


## Installation

```console
$ go get github.com/sio/coolname
```

## Usage

### Importing the package

```go
import "github.com/sio/coolname"
```

### Generating cool names

```
>>> coolname.Slug()
"vegan-outrageous-bumblebee-of-discourse"
nil

>>> coolname.SlugN(2)
"crimson-caracal"
nil

>>> coolname.Generate()
[]string{
  "kind",
  "romantic",
  "markhor",
  "of",
  "luxury",
}
nil
```


### Advanced configuration

Several tunable knobs are provided, check the [docs] and source code to be able to:

  - Provide custom word lists
  - Create new lists by combining existing ones via plain concatenation or via
    cartesian product (see [config.json], custom JSON inputs are supported!)
  - Use custom random number generator

[config.json]: data/config.json


## License and copyright

### Golang port (this package)

Copyright 2023 Vitaly Potyarkin, `Apache-2.0 License`

### Word lists and JSON config

Copyright 2015-2023 [Alexander Lukanin][upstream] and contributors, `BSD-2-Clause License`

# vocab

A Go program to test your knowledge of some vocab.

## Installation

Install with:

```bash
go install github.com/mybearworld/vocab@1.0.0
```

## Usage

```
vocab ./path/to/vocab.json [mode]
```

The vocab.json file contains data in this format:

```json
[
  ["source", "target"],
  ["apple", "manzana"],
  ["orange", "naranja"],
  ...
]
```

The mode can be:

- `reverse`: Asks the questions the other way around.

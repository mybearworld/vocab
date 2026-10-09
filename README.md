# vocab

A Go program to test your knowledge of some vocab.

## Installation

Install with:

```bash
go install github.com/mybearworld/vocab@1.3.2
```

> [!NOTE]
> If you are using Termux, set the `VOCAB_IS_TERMUX` environment variable to
> ensure command line arguments are passed properly.

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

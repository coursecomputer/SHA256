[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

# SHA-256
[EN] Implementation of the SHA-256 hash algorithm

[FR] Implémentation de l'algorithme de hash SHA-256

## Explanation
 - [English](documentation/explanation.en.md)
 - [Francais](documentation/explanation.fr.md)

## Technology
* go [v1.14](https://golang.org/)

## Usage
CLI:
```bash
go test -v ./test
```

CODE:
```go
import "github.com/coursecomputer/SHA256/source"

var sha256 = SHA256.Init()

sha256.Update([]byte("..."))
hash := sha256.Digest()
```

## Links
* https://en.wikipedia.org/wiki/SHA-2
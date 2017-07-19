# go-bedrock
The bedrock for your next go project.

[![Go Report Card](https://goreportcard.com/badge/github.com/mattrmiller/go-bedrock)](https://goreportcard.com/report/github.com/mattrmiller/go-bedrock)
[![godocs](https://img.shields.io/badge/godocs-reference-blue.svg)](https://godoc.org/github.com/mattrmiller/go-bedrock)

# Mission Statement
To provide a neatly written, stable and tested, library that will accelerate the development of your next go project.

# Documentation
- [go-bedrock/brconfig](https://godoc.org/github.com/mattrmiller/go-bedrock/brconfig)
- [go-bedrock/brcrypt](https://godoc.org/github.com/mattrmiller/go-bedrock/brcrypt)
- [go-bedrock/brencoding](https://godoc.org/github.com/mattrmiller/go-bedrock/brencoding)
- [go-bedrock/brstrings](https://godoc.org/github.com/mattrmiller/go-bedrock/brstrings)
- [go-bedrock/brtesting](https://godoc.org/github.com/mattrmiller/go-bedrock/brtesting)

# Install
```
go get github.com/mattrmiller/go-bedrock
```

```
import (
    bedrock "github.com/mattrmiller/go-bedrock"
)
```

# Rules For Contributing
- Please make sure all changed files are run through gofmt
- Submit a PR for review
- Your name will be added below to Contributors

# Helpful Bash Additions
Test recursively, all go tests, in current folder:
```
alias go-testall='go test $(go list ./... | grep -v vendor)'
```

# Author
[Matthew R. Miller](https://github.com/mattrmiller)

# Contributors
[Matthew R. Miller](https://github.com/mattrmiller)

# License
[MIT License](LICENSE)

# go-bedrock
The bedrock for your next go project.

# Mission Statement
To provide a neatly written, stable and tested, library that will accelerate the development of your next go project.

# Documentation
- [brstrings](https://godoc.org/github.com/mattrmiller/go-bedrock/brstrings){:target="_blank"}
- [brtesting](https://godoc.org/github.com/mattrmiller/go-bedrock/brtesting){:target="_blank"}
- [config](https://godoc.org/github.com/mattrmiller/go-bedrock/config){:target="_blank"}

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

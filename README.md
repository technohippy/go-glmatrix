[![Build Status](https://secure.travis-ci.org/technohippy/go-glmatrix.png?branch=master)](http://travis-ci.org/technohippy/go-glmatrix)
[![Coverage Status](https://coveralls.io/repos/technohippy/go-glmatrix/badge.svg?branch=master)](https://coveralls.io/r/technohippy/go-glmatrix?branch=master)
[![GoDoc](https://godoc.org/github.com/technohippy/go-glmatrix?status.svg)](https://godoc.org/github.com/technohippy/go-glmatrix)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/technohippy/go-glmatrix/blob/master/LICENSE.txt)

# go-glmatrix

go-glmatrix is a golang version of [glMatrix](http://glmatrix.net/), which is ``designed to perform vector and matrix operations stupidly fast''.

# Usage

```go
package main

import (
	"fmt"
	glm "github.com/technohippy/go-glmatrix"
)

func main() {
	pers := glm.Mat4Create() // pers is just a float64 slice
	glm.Mat4Perspective(pers, glm.ToRadian(45), 640./480, 0.1, 200.)
	fmt.Println(glm.Mat4Str(pers))
}
```

# Document

- See [https://pkg.go.dev/](https://pkg.go.dev/github.com/technohippy/go-glmatrix)
- or See [the documentation for glMatrix](http://glmatrix.net/docs/)

# License

MIT
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Build Status](https://secure.travis-ci.org/technohippy/go-glmatrix.png?branch=master)](http://travis-ci.org/technohippy/go-glmatrix)
[![Coverage Status](https://coveralls.io/repos/technohippy/go-glmatrix/badge.svg?branch=master)](https://coveralls.io/r/technohippy/go-glmatrix?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/technohippy/go-glmatrix)](https://goreportcard.com/report/github.com/technohippy/go-glmatrix)
[![GoDoc](https://godoc.org/github.com/technohippy/go-glmatrix?status.svg)](https://godoc.org/github.com/technohippy/go-glmatrix)
[![license](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/technohippy/go-glmatrix/blob/master/LICENSE.txt)

# go-glmatrix

go-glmatrix is a golang version of [glMatrix](http://glmatrix.net/), which is ``designed to perform vector and matrix operations stupidly fast''.

## Usage

```go
package main

import (
	"fmt"
	glm "github.com/technohippy/go-glmatrix"
)

func main() {
	// config
	rotateQuat := glm.QuatFromValues(1, 2, 3, 4)
	translateVec3 := glm.Vec3FromValues(1, 2, 3)
	scale := []float64{4, 5, 6}
	rotateOrigin := []float64{7, 8, 9}

	// construct matrix
	transMat := glm.Mat4Create()
	glm.Mat4Identity(transMat)
	glm.Mat4Translate(transMat, transMat, translateVec3)
	glm.Mat4Translate(transMat, transMat, rotateOrigin)
	rotateMat := glm.Mat4Create()
	rotateMat = glm.Mat4FromQuat(rotateMat, rotateQuat)
	matrix := glm.Mat4Multiply(glm.Mat4Create(), transMat, rotateMat)
	glm.Mat4Scale(matrix, matrix, scale)
	negativeOrigin := glm.Vec3Negate(glm.Vec3Create(), rotateOrigin)
	glm.Mat4Translate(matrix, matrix, negativeOrigin)

	// transform position
	position := glm.Vec3FromValues(10, 20, 30)
	glm.Vec3TransformMat4(position, position, matrix)
	fmt.Printf(glm.Vec3Str(position)) // => vec3(1280, -290, -42)
}
```

## Document

- See [https://pkg.go.dev/](https://pkg.go.dev/github.com/technohippy/go-glmatrix)
- or See [the documentation for glMatrix](http://glmatrix.net/docs/)

## License

MIT
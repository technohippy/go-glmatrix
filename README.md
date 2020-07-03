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

See [the documentation for glMatrix](http://glmatrix.net/docs/)

# License

MIT
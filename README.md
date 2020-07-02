go-glmatrix is a golang version of [glMatrix](http://glmatrix.net/)

```go
package main

import (
	"fmt"
	ggm "github.com/technohippy/go-glmatrix"
)

func main() {
	perspective := ggm.Mat4Create()
	ggm.Mat4Perspective(perspective, ggm.ToRadian(45), 640./480, 0.1, 200.)
	fmt.Println(ggm.Mat4Str(perspective))
}
```
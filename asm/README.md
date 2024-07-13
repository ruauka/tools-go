# Asm

## Overview

Package for working with `slices` using `Go Assembly` and `SIMD` optimization.

Optimization will only work for `amd64` architecture with `AVX`.

For `arm64` the functions will be applied automatically from `rslices` package. 

## Content

- [Sum32](#sum32)
- [Sum64](#sum64)
- [Mul32](#mul32)
- [Mul64](#mul64)
- [Mul64Simd](#mul64Simd)

### Sum32

Sums the values `[]float32`.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    s := []float32{1, 2, 3}
    var res float32 = asm.Sum32(s)
    fmt.Println(res) // 6
}
```

### Sum64

Sums the values `[]float64`.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    s := []float64{1, 2, 3}
    var res float64 = asm.Sum64(s)
    fmt.Println(res) // 6
}
```

### Mul32

Multiply arguments `[]float32` element-wise. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
    	s1 = []float32{1, 2, 3}
        s2 = []float32{2, 4, 6}
    )
    
    asm.Mul32(s1, s2)
    fmt.Println(s1) // [2 8 18]
}
```

### Mul64

Multiply arguments `[]float64` element-wise. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
    	s1 = []float64{1, 2, 3}
        s2 = []float64{2, 4, 6}
    )
    
    asm.Mul64(s1, s2)
    fmt.Println(s1) // [2 8 18]
}
```

### Mul64Simd

Multiply arguments `[]float64` element-wise with `simd` optimization.
First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        out = []float64{0, 0, 0}
        s1  = []float64{1, 2, 3}
        s2  = []float64{2, 4, 6}
    )
    
    asm.Mul64Simd(out, s1, s2)
    fmt.Println(out) // [2 8 18]
}
```
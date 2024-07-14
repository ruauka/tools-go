# Asm

## Overview

Package for working with `slices` using `Go Assembly` and `SIMD` optimization.

Optimization only for `amd64` architecture with `AVX`.

For `arm64` the functions will be applied automatically from `rslices` package. 

## Content

- [Sum32](#sum32)
- [Sum64](#sum64)
- [Mul32](#mul32)
- [Mul64](#mul64)
- [Mul64Simd](#mul64Simd)
- [MulNum32](#mulNum32)
- [MulNum64](#mulNum64)
- [Add32](#add32)
- [Add64](#add64)
- [AddNum32](#addNum32)
- [AddNum64](#addNum64)
- [MaximumNum32](#maximumNum32)
- [MaximumNum64](#MaximumNum64)

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

### MulNum32

Multiply arguments `[]float32` element-wise with number. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float32{1, 2, 3}
        num float32 = 2
    )
    
    asm.MulNum32(s1, num)
    fmt.Println(s1) // [2 4 6]
}
```

### MulNum64

Multiply arguments `[]float64` element-wise with number. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float64{1, 2, 3}
        num float64 = 2
    )
    
    asm.MulNum64(s1, num)
    fmt.Println(s1) // [2 4 6]
}
```

### Add32

Add number from `[]float32` for each element in `[]float32`. First argument will be changed.

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
    
    asm.Add32(s1, s2)
    fmt.Println(s1) // [3 6 9]
}
```

### Add64

Add number from `[]float64` for each element in `[]float64`. First argument will be changed.

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
    
    asm.Add64(s1, s2)
    fmt.Println(s1) // [3 6 9]
}
```

### AddNum32

Add number `float32` for each element in `[]float32`. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float32{1, 2, 3}
        num float32 = 2
    )
    
    asm.AddNum32(s1, num)
    fmt.Println(s1) // [3 4 5]
}
```

### AddNum64

Add number `float64` for each element in `[]float64`. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float64{1, 2, 3}
        num float64 = 2
    )
    
    asm.AddNum64(s1, num)
    fmt.Println(s1) // [3 4 5]
}
```

### MaximumNum32

Element-wise maximum of `[]float32` elements with number. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float32{1, 2, 3}
        num float32 = 2
    )
    
    asm.MaximumNum32(s1, num)
    fmt.Println(s1) // [0 2 3]
}
```

### MaximumNum64

Element-wise maximum of `[]float64` elements with number. First argument will be changed.

```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/asm"
)

func main() {
    var (
        s1          = []float64{1, 2, 3}
        num float64 = 2
    )
    
    asm.MaximumNum64(s1, num)
    fmt.Println(s1) // [0 2 3]
}
```

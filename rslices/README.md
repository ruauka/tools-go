# RSlices

## Overview
Package with slices functions.

## Content
- [Sum](#sum)
- [Mul](#mul)
- [MulNum](#mulNum)
- [Add](#add)
- [AddNum](#addnum)
- [MaximumNum](#maximumnum)
- [IsIntersect](#isintersect)
- [Intersection](#intersection)
- [Concat](#concat)

### Sum
Sums the values in a collection. If collection is empty 0 is returned.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32 = []float32{1, 2, 3}
        s64 = []float64{1, 2, 3}
    )
    
    resFloats32 := rslices.Sum(s32)
    fmt.Println(resFloats32) // 6
    
    resFloats64 := rslices.Sum(s64)
    fmt.Println(resFloats64) // 6
}
```

### Mul
Multiplying slice by slice.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32_1 = []float32{1, 2, 3}
        s32_2 = []float32{2, 2, 2}
    
        s64_1 = []float64{1, 2, 3}
        s64_2 = []float64{2, 2, 2}
    )
    
    rslices.Mul(s32_1, s32_2)
    fmt.Println(s32_1) // [2 4 6]
    
    rslices.Mul(s64_1, s64_2)
    fmt.Println(s64_1) // [2 4 6]
}
```

### MulNum
Multiplying slice by number.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32_1          = []float32{1, 2, 3}
        num_32 float32 = 2
        
        s64_1          = []float64{1, 2, 3}
        num_64 float64 = 2
    )
    
    rslices.MulNum(s32_1, num_32)
    fmt.Println(s32_1) // [2 4 6] 
    
    rslices.MulNum(s64_1, num_64)
    fmt.Println(s64_1) // [2 4 6]
}
```

### Add
Adding slice to slice.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32_1 = []float32{1, 2, 3}
        s32_2 = []float32{2, 2, 2}
    
        s64_1 = []float64{1, 2, 3}
        s64_2 = []float64{2, 2, 2}
    )
    
    rslices.Add(s32_1, s32_2)
    fmt.Println(s32_1) // [3 4 5]
    
    rslices.Add(s64_1, s64_2)
    fmt.Println(s64_1) // [3 4 5]
}
```

### AddNum
Adding slice by number.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32_1          = []float32{1, 2, 3}
        num_32 float32 = 2
    
        s64_1          = []float64{1, 2, 3}
        num_64 float64 = 2
    )
    
    rslices.AddNum(s32_1, num_32)
    fmt.Println(s32_1) // [3 4 5] 
    
    rslices.AddNum(s64_1, num_64)
    fmt.Println(s64_1) // [3 4 5]
}
```

### MaximumNum
Set 0 if a < 0.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        s32_1          = []float32{1, 2, 3}
        num_32 float32 = 2
    
        s64_1          = []float64{1, 2, 3}
        num_64 float64 = 2
    )
    
    rslices.MaximumNum(s32_1, num_32)
    fmt.Println(s32_1) // [0 2 3]
    
    rslices.MaximumNum(s64_1, num_64)
    fmt.Println(s64_1) // [0 2 3]
}
```

### IsIntersect
Return ok or not for intersection slice.
```go
package main

import (
    "fmt"
    
    "github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        ints1 = []int{1, 2, 3}
        ints2 = []int{3, 4, 5}
    
        floats1 = []float64{1.1, 2.2, 3.3}
        floats2 = []float64{11.1, 22.2, 44.4}
    
        str1 = []string{"aaa", "bbb", "ccc"}
        str2 = []string{"aaa", "bbb", "ddd"}
    )
    
    resInts := rslices.IsIntersect(ints1, ints2)
    fmt.Println(resInts) // true
    
    resFloats := rslices.IsIntersect(floats1, floats2)
    fmt.Println(resFloats) // false
    
    resStrs := rslices.IsIntersect(str1, str2)
    fmt.Println(resStrs) // true
}
```

### Intersection
Find intersection of two arrays. Returns new slice.
```go
package main

import (
    "fmt"

	"github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        intsL = []int{1, 2, 3}
        intsR = []int{1, 2, 4}
        
        floatsL = []float64{1.1, 2.2, 3.3}
        floatsR = []float64{1.1, 2.2, 4.4}
        
        strL = []string{"aaa", "bbb", "ccc"}
        strR = []string{"aaa", "bbb", "ddd"}
    )
    
    resInts := rslices.Intersection(intsL, intsR)
    fmt.Println(resInts) // [1 2]
    
    resFloats := rslices.Intersection(floatsL, floatsR)
    fmt.Println(resFloats) // [1.1 2.2]
    
    resStrs := rslices.Intersection(strL, strR)
    fmt.Println(resStrs) // [aaa bbb]
}
```

### Concat
Concatenation of multiple slices.

```go
package main

import (
    "fmt"

	"github.com/ruauka/tools-go/rslices"
)

func main() {
    var (
        ints1 = []int{1, 2, 3}
        ints2 = []int{4, 5, 6}
        ints3 = []int{7, 8, 9}
    
        strs1 = []string{"1", "2", "3"}
        strs2 = []string{"4", "5", "6"}
        strs3 = []string{"7", "8", "9"}
    )
    
    ints := rslices.Concat(ints1, ints2, ints3)
    fmt.Println(ints) // [1 2 3 4 5 6 7 8 9]
    
    strs := rslices.Concat(strs1, strs2, strs3)
    fmt.Println(strs) // [1 2 3 4 5 6 7 8 9]
}
```
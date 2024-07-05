# Attrs-go
![](https://img.shields.io/github/go-mod/go-version/ruauka/attrs-go)
[![Coverage Status](https://codecov.io/github/ruauka/attrs-go/coverage.svg?branch=master)](https://codecov.io/gh/ruauka/attrs-go)
[![build](https://github.com/ruauka/attrs-go/actions/workflows/pipeline.yml/badge.svg)](https://github.com/ruauka/attrs-go/actions/workflows/pipeline.yml)
[![GoDoc](https://godoc.org/github.com/ruauka/attrs-go?status.svg)](https://godoc.org/github.com/ruauka/attrs-go)

## Overview
Tool for working with structure fields:
 - analog of Python 'getattr' and 'setattr'
 - some useful funcs to changing and rounding struct fields

## Content

- [Installation](#installation)
- [Usage](#usage)
    - [GetAttr](#getattr)
    - [SetAttr](#setattr)
    - [Round](#round)
    - [SetStructAttrs](#setstructattrs)
    - [RoundFloatStruct](#roundfloatstruct)
    - [Intersection](#intersection)
    - [SlicesConcat](#slicesconcat)
    - [TimeDelta](#timedelta)

## Installation
To install the package run
```bash
go get -u github.com/ruauka/attrs-go
```

## Usage

### GetAttr
Get struct field value.

```go
package main

import (
    "fmt"
    
	"github.com/ruauka/tools-go/attrs"
)

type User struct {
    Username string
}

func main() {
    user := User{Username: "username value"}
    
    value, _ := attrs.GetAttr(user, "Username")
    fmt.Println(value) // username value
}
```

### SetAttr
Set new value at structure field.

```go
package main

import (
    "fmt"

  "github.com/ruauka/tools-go/attrs"
)

type User struct {
    Username string
}

func main() {
    u := &User{Username: "username value"}
    
    if err := attrs.SetAttr(u, "new username value", "Username"); err != nil {
      fmt.Println(err)
    }
    
    fmt.Println(u.Username) // new username value
}
```

### Round
Float64 and Float32 rounder to certain precision.
```go
package main

import (
    "fmt"
    "reflect"

  "github.com/ruauka/tools-go/attrs"
)

var (
    val32 float32 = 0.12345
    val64 float64 = 0.12345
)

func main() {
    res32 := attrs.Round(val32, 3)
    fmt.Println(res32)                 // 0.123
    fmt.Println(reflect.TypeOf(res32)) // float32
  
    res64 := attrs.Round(val64, 3)
    fmt.Println(res64)                 // 0.123
    fmt.Println(reflect.TypeOf(res64)) // float64
}
```

### SetStructAttrs
Update current structure fields with the values of the new structure fields.


```go
package main

import (
    "fmt"

  "github.com/ruauka/tools-go/attrs"
)

type User struct {
    Username string // will change by pte
    Age      int    // will change by value
    Married  bool   // will be the same
}

type NewUser struct {
    Username *string `json:"username"`
    Age      int     `json:"age"`
    Married  *bool   `json:"married"` // nil
}

func main() {
    user := &User{
        Username: "username",
        Age:      30,
        Married:  true,
    }
    
    newUserName := "new_username"
    
    newUser := NewUser{
        Username: &newUserName,
        Age:      35,
        Married:  nil,
    }
    
    fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married) // username, 30, true
    
    if err := attrs.SetStructAttrs(user, newUser); err != nil {
        fmt.Println(err)
    }
    
    fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married) // new_username, 35, true
}
```

### RoundFloatStruct
Round up float struct fields to certain precision.

```go
package main

import (
    "fmt"

  "github.com/ruauka/tools-go/attrs"
)

type Foo struct {
    Field1 float32
    Field2 float64
    Field3 []float32
    Field4 []float64
    Field5 [3]float32
    Field6 [3]float64
    Field7 int    // will be the same
    Field8 string // will be the same
}

func main() {
    foo := &Foo{
        Field1: 1.1111,
        Field2: 2.2222,
        Field3: []float32{1.1111, 2.2222, 3.3333},
        Field4: []float64{4.4444, 5.5555, 7.7777},
        Field5: [3]float32{1.1111, 2.2222, 3.3333},
        Field6: [3]float64{4.4444, 5.5555, 7.7777},
        Field7: 7,
        Field8: "field8",
    }

    fmt.Printf("%+v\n", *foo)
    // {
    //Field1:1.1111 Field2:2.2222
    //Field3:[1.1111 2.2222 3.3333] Field4:[4.4444 5.5555 7.7777]
    //Field5:[1.1111 2.2222 3.3333] Field6:[4.4444 5.5555 7.7777]
    //Field7:7 Field8:field8
    //}
    
    if err := attrs.RoundFloatStruct(foo, 3); err != nil {
        fmt.Println(err)
    }
    
    fmt.Printf("%+v", *foo)
    // {
    //Field1:1.112 Field2:2.223
    //Field3:[1.112 2.223 3.334] Field4:[4.445 5.556 7.778]
    //Field5:[1.112 2.223 3.334] Field6:[4.445 5.556 7.778]
    //Field7:7 Field8:field8
    //}
}
```

### Intersection
Find intersection of two arrays. Returns new slice.
```go
package main

import (
    "fmt"

  "github.com/ruauka/tools-go/attrs"
)

var (
    intsL = []int{1, 2, 3}
    intsR = []int{1, 2, 4}
    
    floatsL = []float64{1.1, 2.2, 3.3}
    floatsR = []float64{1.1, 2.2, 4.4}
    
    strL = []string{"aaa", "bbb", "ccc"}
    strR = []string{"aaa", "bbb", "ddd"}
)

func main() {
    resInts, _ := attrs.Intersection(intsL, intsR)
    fmt.Println(resInts) // [1 2]
    
    resFloats, _ := attrs.Intersection(floatsL, floatsR)
    fmt.Println(resFloats) // [1.1 2.2]
    
    resStrs, _ := attrs.Intersection(strL, strR)
    fmt.Println(resStrs) // [aaa bbb]
}
```

### SlicesConcat
Concatenation of multiple slices.

```go
package main

import (
    "fmt"

  "github.com/ruauka/tools-go/attrs"
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
    
    ints := attrs.SlicesConcat(ints1, ints2, ints3)
    fmt.Println(ints) // [1 2 3 4 5 6 7 8 9]
    
    strs := attrs.SlicesConcat(strs1, strs2, strs3)
    fmt.Println(strs) // [1 2 3 4 5 6 7 8 9]
}
```

### TimeDelta
The difference between the two dates for each value

```go
package main

import (
    "fmt"
    "time"
    
    attrs "github.com/ruauka/attrs-go"
)

func main() {
    var (
        from = time.Date(2022, 5, 25, 1, 1, 1, 1, time.UTC)
        to   = time.Date(2023, 5, 25, 1, 1, 1, 1, time.UTC)
    )
    
    res := attrs.Elapsed(from, to)
    fmt.Println(res)           // &{1 0 0 0 0 0 0 12 365 8760 525600 31536000}
    fmt.Println(res.TotalDays) // 365
}
```
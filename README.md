# Tools-go
![](https://img.shields.io/github/go-mod/go-version/ruauka/attrs-go)
[![Coverage Status](https://codecov.io/github/ruauka/attrs-go/coverage.svg?branch=master)](https://codecov.io/gh/ruauka/attrs-go)
[![build](https://github.com/ruauka/attrs-go/actions/workflows/pipeline.yml/badge.svg)](https://github.com/ruauka/attrs-go/actions/workflows/pipeline.yml)
[![GoDoc](https://godoc.org/github.com/ruauka/attrs-go?status.svg)](https://godoc.org/github.com/ruauka/attrs-go)

## Overview
The tool contains some useful functions:
- Attrs - changing and rounding struct fields (`getattr`, `setattr`, etc...)
- rTime - counting time

## Content

- [Installation](#installation)
- [Usage](#usage)
  - [Attrs](attrs/README.md)
  - [rTime](rtime/README.md)
  - [Intersection](#intersection)
  - [SlicesConcat](#slicesconcat)
  

## Installation
To install the package run
```bash
go get -u github.com/ruauka/tools-go
```

## Usage


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
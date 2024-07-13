# RTime

## Overview
Package with time functions:
- Calculating the time difference.
  - Months
  - Days

## Content
- [Months](#months)
- [Days](#days)

### Months
The difference between the two dates.
```go
package main

import (
    "fmt"
    "time"
    
    "github.com/ruauka/tools-go/rtime"
)

func main() {
    var (
        d1 = time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC)
        d2 = time.Date(2023, 1, 30, 0, 0, 0, 0, time.UTC)
    )
    
    res := rtime.Months(d1, d2)
    fmt.Println(res) // 10
}
```

### Days
The number of days between two dates.
```go
package main

import (
    "fmt"
    "time"
    
    "github.com/ruauka/tools-go/rtime"
)

func main() {
    var (
        d1 = time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC)
        d2 = time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC)
    )
    
    res := rtime.Days(d1, d2)
    fmt.Println(res) // 365
}
```

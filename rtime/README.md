# RTime

## Overview
Package with time functions:
- Calculating the time difference. 
  - Years
  - Months
  - Days
  - Hours
  - Minutes
  - Seconds
  - Nanoseconds

## Content
- [TimeDelta](#timedelta)

### TimeDelta
The difference between the two dates for each value

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/ruauka/tools-go/rtime"
)

func main() {
    var (
        from = time.Date(2022, 5, 25, 1, 1, 1, 1, time.UTC)
        to   = time.Date(2023, 5, 25, 1, 1, 1, 1, time.UTC)
    )
    
    res := rtime.Elapsed(from, to)
    fmt.Println(res)           // &{1 0 0 0 0 0 0 12 365 8760 525600 31536000}
    fmt.Println(res.TotalDays) // 365
}
```

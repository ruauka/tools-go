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
- [Months](#months)
- [Days](#days)

### TimeDelta
The difference between the two dates for each value.
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

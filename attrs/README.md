# Attrs

## Overview
Package for working with structure fields:
- Analog of Python `getattr` and `setattr`
- Funcs to changing and rounding struct fields

## Content

- [GetAttr](#getattr)
- [SetAttr](#setattr)
- [Round](#round)
- [SetStructAttrs](#setstructattrs)
- [RoundStructFloatFields](#roundStructFloatFields)

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

### RoundStructFloatFields
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
    
    if err := attrs.RoundStructFloatFields(foo, 3); err != nil {
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

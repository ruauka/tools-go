## Project description
Tool for working with structure fields.

## Install
```bash
go get github.com/ruauka/attrs-go@v1.0.4
```

## Usage
<ins>**GetAttr**</ins> - get struct field value. 

Args:
 - obj, fieldName - value param. Struct fields can be ptr or value.

```go
func GetAttr(obj interface{}, fieldName string) (interface{}, error)
```

```go
type User struct {
    Username string
}

user := User{Username: "username value"}

value, err := attrs_go.GetAttr(user, "Username")
fmt.Println(value) // username value
```

<ins>**SetAttr**</ins>  - set new value on structure field.

Args:
- obj - ptr param. Struct fields can be ptr or value.
- fieldName, newValue - value param.

```go
func SetAttr(obj interface{}, fieldName string, newValue interface{}) error
```

```go
type User struct {
    Username string
}

u := User{Username: "username value"}

if err := attrs_go.SetAttr(&u, "Username", "new username value"); err != nil {
    fmt.Println(err)
}

fmt.Println(u.Username) // new username value
```

<ins>**SetStructAttrs**</ins>  - updates current structure fields with the values of the new structure fields.
Args:
- curObj - ptr param. Struct fields can be ptr or value.
- newObj - value param. Struct fields can be ptr or value.

```go
func SetStructAttrs(curObj, newObj interface{}) error
```

```go
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

user := &User{
    Username: "username",
    Age:      30,
    Married:  nil,
}

newUserName := "new_username"
newUser := NewUser{
    Username: &newUserName,
    Age:      35,
}

fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married) // username, 30, false

if err := attrs_go.SetStructAttrs(user, newUser); err != nil {
    fmt.Println(err)
}

fmt.Printf("%s, %d, %v\n", user.Username, user.Age, user.Married) // new_username, 35, false
```
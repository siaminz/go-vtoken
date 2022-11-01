# github.com/ItsAminZamani/go-vtoken

## Usage


### Get the go-vtoken module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/ItsAminZamani/go-vtoken@v0.1.0
```

```go
package main

import (
    "fmt"
    "time"

    vtoken "github.com/ItsAminZamani/go-vtoken"
)

func main() {
    token, err := vtoken.New(identifier="user@domain.com", expires=time.Hour)
    if err != nil {
		panic(err)
	}
    fmt.Println(token)
    id, err := vtoken.Verify(token)
    if err != nil {
		panic(err)
	}
    fmt.Println(id.GetIdentifier())
}
```

## Expiration

```go
id.IsExpired()
```
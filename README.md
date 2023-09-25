# github.com/siaminz/go-vtoken

## Usage


### Get the go-vtoken module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/siaminz/go-vtoken@v0.2.0
```

```go
package main

import (
    "fmt"
    "time"

    vtoken "github.com/siaminz/go-vtoken"
)

func main() {
    token, err := vtoken.New("user@domain.com", time.Hour)
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

## Simple Token
Generate a token with custom length. this kind of tokens can't be verified.

```go
package main

import (
    "fmt"
    "time"

    vtoken "github.com/siaminz/go-vtoken"
)

func main() {
    token := vtoken.GenerateSimpleToken(5)
    fmt.Println(token)
}
```
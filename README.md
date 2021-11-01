# Cell

Cell is a Go package that creates new instances by string in running time.

Getting Started
===============

## Installing

To start using CELL, install Go and run `go get`:

```sh
$ go get -u github.com/projectred/cell
```

## Regist

Befort using CELL, you should regist the Types that you are going to create in runnig time.

```go
type A interface {
	Value() int
}

type DefaultA struct {
	v int `tag:"v"`
}

func (a *DefaultA) Fill(k string, v interface{},tag string){
    if k == "v"{
        a.V = v.(string)
    }
}

func (a *DefaultA) Value() int { return a.V }

func init() {
    cell.Regist("defaultA", &DefaultA{v: 1})
}
```

## Running

```go
package main

import "github.com/projectred/cell"

func main() {
	a := cell.Spilt("defaultA", nil).(A)
    a.Value() // 1
    a = cell.Spilt("defaultA", &SplitOptions{[]string{"v"}, []interface{}{1024}).(A)
    a.Value() // 1024
}
```

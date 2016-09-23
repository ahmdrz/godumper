# godumper
> Dump array of interfaces to CSV files.

[![Build Status](https://travis-ci.org/ahmdrz/godumper.svg?branch=master)](https://travis-ci.org/ahmdrz/godumper)
[![GoDoc](https://godoc.org/github.com/ahmdrz/godumper?status.svg)](https://godoc.org/github.com/ahmdrz/godumper)

### How to use ?

```bash
go get github.com/ahmdrz/godumper
```

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/godumper"
)

type Message struct {
	Id       int
	UserFrom int
	UserTo   int
	Text     string
	Time     int64
	IsReaded bool
}

var array = []Message{
	Message{
		Id:       810625,
		Time:     1475430311,
		Text:     "Hi dude, Are you okay ?",
		UserFrom: 812311,
		UserTo:   722311,
		IsReaded: true,
	},
	Message{
		Id:       187236,
		UserFrom: 722311,
		UserTo:   812311,
		Text:     "Hey buddy,I'm fine",
		Time:     1475430322,
		IsReaded: false,
	},
	Message{
		Id:       715623,
		UserFrom: 722311,
		UserTo:   812311,
		Text:     "😆😆😆 😛",
		Time:     1475430621,
		IsReaded: false,
	},
}

func main() {
	dumper, err := godumper.New(Message{})
	if err != nil {
		panic(err)
	}
	dumper, err = dumper.Dump(array)
	if err != nil {
		panic(err)
	}
	err = dumper.Save("result.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("The array dumped to result.csv")
}
```

### License 

Read the license file

### Contribe 

Just fork and make pull request. I will happy if you tell me bugs and help me to improve it.

# godumper
> Dump array of interfaces to CSV/XSLX files.

[![Build Status](https://travis-ci.org/ahmdrz/godumper.svg?branch=master)](https://travis-ci.org/ahmdrz/godumper)
[![GoDoc](https://godoc.org/github.com/ahmdrz/godumper?status.svg)](https://godoc.org/github.com/ahmdrz/godumper)

- [x] Dump array of interfaces
- [x] Dump slice of interfaces
- [ ] Dump map of interfaces
- [ ] ...

### How to use ?

```bash
go get github.com/ahmdrz/godumper
```

First of all...

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/godumper"
)

type Message struct {
	Id       int `dump:"index"`
	UserFrom int
	UserTo   int
	Text     string
	Time     int64
	IsReaded bool
}
```

**You can use `dump:"index"` for change header name**

If you want to dump a slice

```go
var slice = []Message{
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
	dumper, err := godumper.New(Message{}, godumper.CSV)
	if err != nil {
		panic(err)
	}
	
	err = dumper.Dump(slice)
	if err != nil {
		panic(err)
	}
	
	err = dumper.Save("result.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("The slice dumped to result.csv")
}
```

Or if you want to dump an array...

```go
var array = [3]Message{
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
	dumper, err := godumper.New(Message{}, godumper.XSLX)
	if err != nil {
		panic(err)
	}
	
	err = dumper.Dump(array)
	if err != nil {
		panic(err)
	}
	
	err = dumper.Save("result.xslx")
	if err != nil {
		panic(err)
	}
	fmt.Println("The array dumped to result.xslx")
}
```

### DumpAndSave

If you use `Save` function , another `[][]string` will created and memory used by not important datas.

So you can use `DumpAndSave` for save memory , see below :

```go
func main() {
	dumper, err := godumper.New(Message{}, godumper.CSV)
	if err != nil {
		panic(err)
	}
	err = dumper.DumpAndSave(array, "result.csv")
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

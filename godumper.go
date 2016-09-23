// godumper project godumper.go
package godumper

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Dumper struct {
	Header []string
	Type   interface{}
	Body   [][]string
}

func New(item interface{}) (*Dumper, error) {
	dumper := reflect.ValueOf(item)
	if dumper.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Invalid input, input must be a struct")
	}
	header := make([]string, 0)
	for i := 0; i < dumper.NumField(); i++ {
		header = append(header, dumper.Type().Field(i).Name)
	}
	return &Dumper{
		Type:   item,
		Header: header,
	}, nil
}

func (dumper *Dumper) Dump(datasets interface{}) (*Dumper, error) {
	items := reflect.ValueOf(datasets)
	if items.Kind() == reflect.Slice {
		result := make([][]string, items.Len())
		for i := range result {
			result[i] = make([]string, len(dumper.Header))
		}

		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				for j := 0; j < v.NumField(); j++ {
					if j >= len(result[i]) {
						return nil, fmt.Errorf("index out of range , each indexes must same as header")
					}
					result[i][j] = toString(v.Field(j).Interface())
				}
				dumper.Body = result
			} else {
				return nil, fmt.Errorf("not struct", item.Kind())
			}
		}
	} else {
		return nil, fmt.Errorf("input should be a slice")
	}

	return dumper, nil
}

func toString(item interface{}) string {
	switch item.(type) {
	case string:
		return item.(string)
	case int:
		return strconv.Itoa(item.(int))
	case int64:
		return strconv.FormatInt(item.(int64), 10)
	case bool:
		if item.(bool) == true {
			return "true"
		} else {
			return "false"
		}
	}
	return ""
}

func (dumper *Dumper) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write(dumper.Header)
	if err != nil {
		return err
	}

	for _, value := range dumper.Body {
		err = writer.Write(value)
		if err != nil {
			return err
		}
	}

	defer writer.Flush()
	return nil
}

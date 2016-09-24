// This library can dump slice , array and etc to CSV files.
package godumper

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// This struct contains Header , Type and Body.
// Header is header of CSV file.
// Body is 2D array for the body of CSV file.
// And Type is the type of interface which you want to dump.
type Dumper struct {
	Header []string
	Type   interface{}
	Body   [][]string
}

// This method return an error only if
//   Invalid input, input must be a struct
// And return the Dumper struct as base of this library.
// The input must be a struct.
//   Todo : allow method to receive map values.
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

// This method return error if :
// 	index out of range (body must be a child of header)
//  not struct input for body
//  input was not a slice or array
// And after called , the body is ready for save or something else...
func (dumper *Dumper) Dump(datasets interface{}) error {
	items := reflect.ValueOf(datasets)
	if items.Kind() == reflect.Slice || items.Kind() == reflect.Array {
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
						return fmt.Errorf("index out of range , each indexes must same as header")
					}
					result[i][j] = toString(v.Field(j).Interface())
				}
				dumper.Body = result
			} else {
				return fmt.Errorf("not struct", item.Kind())
			}
		}
	} else {
		return fmt.Errorf("input should be a slice or array")
	}

	return nil
}

// Change interface to string for saving to CSV file and making the body.
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

// This method return error if :
//	can't create file.
//  can't write header.
//  can't write body
// And after called , Header and Body will available in file.
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

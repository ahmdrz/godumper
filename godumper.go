// This library can dump slice , array and etc to CSV files.
package godumper

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"

	"github.com/tealeg/xlsx"
)

// Dumper , This struct contains Header , Type and Body.
// Header is header of CSV file.
// Body is 2D array for the body of CSV file.
// And Type is the type of interface which you want to dump.
type Dumper struct {
	Header     []string
	Type       interface{}
	Body       [][]string
	outputType int
}

const (
	XSLX int = 10 << iota
	CSV
)

// New , This method return an error only if
//   Invalid input, input must be a struct
// And return the Dumper struct as base of this library.
// The input must be a struct.
//   Todo : allow method to receive map values.
func New(item interface{}, outputType int) (*Dumper, error) {
	if outputType != XSLX && outputType != CSV {
		return nil, fmt.Errorf("Invalid output type")
	}
	dumper := reflect.ValueOf(item)
	if dumper.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Invalid input, input must be a struct")
	}
	header := make([]string, 0)
	for i := 0; i < dumper.NumField(); i++ {
		field := dumper.Type().Field(i)
		tag := field.Tag.Get("dump")
		// TODO : ignore `-` tag.
		if len(tag) > 0 {
			header = append(header, tag)
		} else {
			header = append(header, field.Name)
		}
	}
	return &Dumper{
		Type:       item,
		Header:     header,
		outputType: outputType,
	}, nil
}

// Dump This method return error if :
// 	index out of range (body must be a child of header)
//  not struct input for body
//  input was not a slice or array
// And after called , the body is ready for save or something else...
func (dumper *Dumper) Dump(inputSet interface{}) error {
	items := reflect.ValueOf(inputSet)
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
				return fmt.Errorf("not struct %v ", item.Kind())
			}
		}
	} else {
		return fmt.Errorf("input should be a slice or array")
	}

	return nil
}

// Change interface to string for saving to CSV file and making the body.
func toString(item interface{}) string {
	return fmt.Sprintf("%v", item)
}

// Save This method return error if :
//	can't create file.
//  can't write header.
//  can't write body
// And after called , Header and Body will available in file.
func (dumper *Dumper) Save(filename string) error {
	if dumper.outputType == CSV {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()

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

		return nil
	} else {
		var err error

		file := xlsx.NewFile()
		sheet, err := file.AddSheet("SHEET1")
		if err != nil {
			return err
		}
		headerRow := sheet.AddRow()
		for _, headerItem := range dumper.Header {
			cell := headerRow.AddCell()
			cell.Value = headerItem
		}
		for _, row := range dumper.Body {
			singleRow := sheet.AddRow()
			for _, rowCell := range row {
				cell := singleRow.AddCell()
				cell.Value = rowCell
			}
		}
		err = file.Save(filename)
		if err != nil {
			return err
		}
		return nil
	}
}

// DumpAndSave This method return error if :
// 	index out of range (body must be a child of header)
//  not struct input for body
//  input was not a slice or array
// And after called , the body is ready for save or something else...
func (dumper *Dumper) DumpAndSave(inputSet interface{}, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(dumper.Header)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(inputSet)
	if items.Kind() == reflect.Slice || items.Kind() == reflect.Array {
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			result := make([]string, len(dumper.Header))
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				for j := 0; j < v.NumField(); j++ {
					if j >= len(dumper.Header) {
						return fmt.Errorf("index out of range , each indexes must same as header")
					}
					result[j] = toString(v.Field(j).Interface())
				}
				err = writer.Write(result)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("not struct %v", item.Kind())
			}
		}
	} else {
		return fmt.Errorf("input should be a slice or array")
	}

	return nil
}

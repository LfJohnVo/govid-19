package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
)

//iterate a json with struct
func Desglose(x Stats) {
	v := reflect.ValueOf(x)
	typeOfS := v.Type()

	values := make([]interface{}, v.NumField())

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Field", "Value"})

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		data, _ := WriteTable(i, typeOfS.Field(i).Name, v.Field(i).Interface())
		for _, v := range data {
			table.Append(v)
		}

	}//end for

	table.Render() // Send output


}//end func

//draw table
func WriteTable(i int, field string, value interface{}) ([][]string, [][]string) {
	y := strconv.Itoa(i)
	valor := fmt.Sprintf("%v", value)
	data := [][]string{
		[]string{y, field, valor},
	}

	return data, nil
}//end func

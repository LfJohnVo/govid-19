package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
)

//iterate a json with struct
func DesgloseStats(x Stats) {
	v := reflect.ValueOf(x)
	typeOfS := v.Type()

	values := make([]interface{}, v.NumField())

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Field", "Value"})

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		data, _ := WriteTableStats(i, typeOfS.Field(i).Name, v.Field(i).Interface())
		for _, v := range data {
			table.Append(v)
		}

	} //end for

	table.Render() // Send output

} //end func

//iterate a json with struct
func SumarioGlobal(g Sum) {
	x := g.Global
	fmt.Println("")
	fmt.Println("----- GLOBAL DATA -----")
	v := reflect.ValueOf(x)
	typeOfS := v.Type()

	values := make([]interface{}, v.NumField())

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Field", "Value"})

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		//fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		data, _ := WriteTableStats(i, typeOfS.Field(i).Name, v.Field(i).Interface())
		for _, v := range data {
			table.Append(v)
		}

	} //end for

	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output

} //end func

func SumarioCountries(h Sum) {
	x := h.Countries
	//fmt.Println(x)
	fmt.Println("")
	fmt.Println("----- COUNTRIES DATA -----")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Country", "Country Code", "Slug", "NewConfirmed", "TotalConfirmed", "NewDeaths", "TotalDeaths", "NewRecovered", "TotalRecovered"})
	for _, element := range x {
		data, _ := WriteTableCountries(element.Country, element.CountryCode, element.Slug, element.NewConfirmed, element.TotalConfirmed, element.NewDeaths, element.TotalDeaths, element.NewRecovered, element.TotalRecovered)
		for _, v := range data {
			table.Append(v)
		}
	}
	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output
} //Summary func

func SumarioDate(h Sum) {
	x := h.Date
	//fmt.Println(x)
	fmt.Println("----- DATE TIME -----")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date"})
	data, _ := WriteTableDate(x)
	for _, v := range data {
		table.Append(v)
	}

	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output

}

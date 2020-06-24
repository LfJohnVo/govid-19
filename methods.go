package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
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

	table.SetAlignment(tablewriter.ALIGN_CENTER)
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

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output
} //Summary func

func Paisessum(Countries Countries) {
	x := Countries
	dic := make(map[string]string)
	//fmt.Println("%v\n", x)
	fmt.Println("")
	fmt.Println("----- COUNTRIES DATA [10 OF " + strconv.Itoa(len(x)) + "] -----")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Country", "Slug", "ISO2"})

	for index, element := range x {
		data, _ := WriteTableCountry(element.Country, element.Slug, element.ISO2)
		for _, v := range data {
			table.Append(v)
		}
		if index > 10 {
			break
		}
	}

	for _, e := range x {
		dic[e.Country] = e.Slug
	}

	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output

	fmt.Println("Search slug country option (type some chars):")
	var pais string
	fmt.Scanf("%s", &pais)
	for _, e := range dic {
		if strings.Contains(e, strings.ToLower(pais)) {
			fmt.Println(e)
		}
	}
}

func Paissum(Country Country) {
	x := Country
	//fmt.Println("%v\n", x)
	fmt.Println("")
	fmt.Println("----- COUNTRY DATA [LAST DAY] -----")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Confirmed", "Deaths", "Recovered", "Active", "Date"})

	for index, element := range x {
		data, _ := GenericTableCountry(element.Confirmed, element.Deaths, element.Recovered, element.Active, element.Date)
		if index+1 == len(x) {
			for _, v := range data {
				table.Append(v)
			}
		}
	}

	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output

}

func DesgloseAll(All Alldatastruct) {
	x := All
	fmt.Println("")
	fmt.Println("----- ALL WORLD DATA -----")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Country", "CountryCode", "Lat", "Lon", "Confirmed", "Deaths", "Recovered", "Active", "Date", "LocationID"})
	//fmt.Printf("%v", x)
	for _, element := range x {
		data, _ := WriteTableAlldata(element.Country, element.CountryCode, element.Lat, element.Lon, element.Confirmed, element.Deaths, element.Recovered, element.Active, element.Date, element.LocationID)
		for _, v := range data {
			table.Append(v)
		}
	}
	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output
}

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

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render() // Send output

}

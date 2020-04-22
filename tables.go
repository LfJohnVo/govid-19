package main

import (
	"fmt"
	"strconv"
	"time"
)

//draw table
func WriteTableStats(i int, field string, value interface{}) ([][]string, [][]string) {
	y := strconv.Itoa(i)
	valor := fmt.Sprintf("%v", value)
	data := [][]string{
		[]string{y, field, valor},
	}

	return data, nil
}//end func

//draw table
func WriteTableCountries(country string, code string, slug string, confirmed int, totalConfirmed int, deaths int, totalDeaths int, recovered int, totalRecovered int) ([][]string, [][]string) {
	//y := strconv.Itoa(i)
	//valor := fmt.Sprintf("%v", value)
	data := [][]string{
		[]string{country, code, slug, strconv.Itoa(confirmed), strconv.Itoa(totalConfirmed), strconv.Itoa(deaths), strconv.Itoa(totalDeaths), strconv.Itoa(recovered), strconv.Itoa(totalRecovered)},
	}

	return data, nil
}//end func

//draw table
func WriteTableDate(Date time.Time) ([][]string, [][]string) {
	data := [][]string{
		[]string{Date.String()},
	}

	return data, nil
}//end func
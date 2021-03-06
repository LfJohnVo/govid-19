package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Menu() {

	var input int
	n, err := fmt.Scanln(&input)

	if n < 1 || err != nil {
		fmt.Println("Invalid input")
		time.Sleep(5 * time.Millisecond)
		return
	}
	switch input {
	case 1:
		Summary("summary")
	case 2:
		Paises("countries")
	case 3:
		Alldata("all")
	case 4:
		Estadisticas("stats")
	case 5:
		Version("version")
	default:
		os.Exit(2)
	}

	fmt.Println("")
	fmt.Println("**** Press 1 to continue ****")
	fmt.Println("**** Press any button to exit ****")
	var input2 int
	yeah, err := fmt.Scanln(&input2)
	if yeah < 1 || err != nil {
		os.Exit(2)
	} else {
		fmt.Printf("Returning to the main menu")
		bar := progressbar.New(100)
		for i := 0; i < 100; i++ {
			bar.Add(1)
			time.Sleep(10 * time.Millisecond)
		}
		defer time.Sleep(5 * time.Millisecond)
		fmt.Println("")
	}

} //Menu func

func connect(op string) []byte {
	resp, err := http.Get("https://api.covid19api.com/" + op)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
} //connect func

func Summary(opcion string) {
	fmt.Println("1.- A summary of new and total cases per country updated daily:")
	bodyBytes := connect(opcion)
	var todoSum Sum
	json.Unmarshal(bodyBytes, &todoSum)
	x := todoSum
	SumarioCountries(x)
	SumarioGlobal(x)
	SumarioDate(x)
}

func Paises(opcion string) {
	fmt.Println("2.- Returns all the available countries and provinces, as well as the country slug for per country requests:")
	bodyBytes := connect(opcion)
	var todoCountries Countries
	json.Unmarshal(bodyBytes, &todoCountries)
	x := todoCountries
	//fmt.Printf("%+v\n", x)
	Paisessum(x)
	fmt.Println("Type slug option:")
	var slug string
	fmt.Scanf("%s", &slug)
	Pais("country/" + slug)
}

func Pais(opcion string) {
	fmt.Println("2.1 - Returns all cases by case type for a country. Country must be the country_slug from /countries:")
	bodyBytes := connect(opcion)
	var todoCountry Country
	json.Unmarshal(bodyBytes, &todoCountry)
	x := todoCountry
	//fmt.Printf("%+v\n", x)
	Paissum(x)

}

func Alldata(opcion string) {
	fmt.Println("3.- Returns all daily data. This call results in 10MB of data being returned and should be used infrequently.\n\n")
	fmt.Println("This can take a little more of time\n\n")
	bodyBytes := connect(opcion)
	var todoData Alldatastruct
	json.Unmarshal(bodyBytes, &todoData)
	x := todoData
	DesgloseAll(x)
}

func Estadisticas(opcion string) {
	fmt.Println("4.- This route returns the usage of the API. This is not for any COVID related statistics:")
	bodyBytes := connect(opcion)
	var todoStats Stats
	json.Unmarshal(bodyBytes, &todoStats)

	x := todoStats
	DesgloseStats(x)

} //Estadisticas func

func Version(opcion string) {
	fmt.Println("5.- This route returns OMS Api Version:")
	bodyBytes := connect(opcion)
	fmt.Printf("%v", bodyBytes)

} //version func

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Menu(){

	var input int
	n, err := fmt.Scanln(&input)

	if n < 1 || err != nil {
		fmt.Println("Invalid input")
		return
	}
	switch input {
	case 1:
		Estadisticas("stats")
	case 2:

	default:
		os.Exit(2)
	}

	fmt.Printf("Do you want to choose another option? Y/N")
	fmt.Println("")
	var volv string
	opt, err := fmt.Scanln(&volv)
	fmt.Println(opt)
	if opt == 'Y' || opt == 'y' {
		fmt.Println("")
		fmt.Println("Regresando al menú en 5 segundos")
		defer time.Sleep(5*time.Second)
		fmt.Println("")
	} else if opt == 'N' || opt == 'n' {
		os.Exit(2)
	} else {
		fmt.Println("Invalid input")
		time.Sleep(2*time.Second)
		fmt.Println("Returning to the main menu")
		return
	}

}//Menu func

func connect(op string) []byte {
	resp, err := http.Get("https://api.covid19api.com/"+op)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}//connect func

func Estadisticas(opcion string) {
	fmt.Println("1.- This route returns the usage of the API. This is not for any COVID related statistics:")
	bodyBytes := connect(opcion)
	// Convert response body to string
	//bodyString := string(bodyBytes)
	//fmt.Println("API Response as String:\n" + bodyString)
	// Convert response body to Todo struct
	var todoStats Stats
	json.Unmarshal(bodyBytes, &todoStats)
	fmt.Printf("API Response as struct %+v\n", todoStats)
}//Estadisticas func
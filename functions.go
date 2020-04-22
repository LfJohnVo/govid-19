package main

import (
	"encoding/json"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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
		Estadisticas("stats")
	case 2:

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

func Estadisticas(opcion string) {
	fmt.Println("1.- This route returns the usage of the API. This is not for any COVID related statistics:")
	bodyBytes := connect(opcion)
	var todoStats Stats
	json.Unmarshal(bodyBytes, &todoStats)

	x := todoStats
	Desglose(x)

} //Estadisticas func

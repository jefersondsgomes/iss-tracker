package main

import (
	"fmt"
	"os"

	"github.com/jefersondsgomes/iss-tracker/services"
)

func main() {
	fmt.Println(">---------- Tracking International Space Station ----------<")

	response, err := services.TrackISS()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Println("Date: ", services.FormatDate(response.Date))
	fmt.Println("Longitude: ", response.ISSPosition.Longitude)
	fmt.Println("Latitude: ", response.ISSPosition.Latitude)
}

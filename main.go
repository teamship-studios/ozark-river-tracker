package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/api/router"
	"github.com/river-folk/ozark-river-tracker/pkg/usgs"
)

func main() {
	fmt.Println("Ozark river tracker!")

	test, err := usgs.ReadGauge("07067000", []string{usgs.GaugeHeight, usgs.Discharge, "91110"})
	if err != nil {
		fmt.Println(err)
		return
	}

	discharge, _ := test.GetMostRecentDischarge()
	height, _ := test.GetMostRecentGaugeHeight()
	lat, lon, _ := test.GetCoordinates()

	fmt.Println(discharge)
	fmt.Println(height)
	fmt.Println(lat)
	fmt.Println(lon)
	// fmt.Println(test)

	connection, err := repository.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	http := gin.Default()

	router.Setup(http, connection)

	// TODO Set with ENV var or maybe just use "localhost"
	err = http.Run("127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
	}
}

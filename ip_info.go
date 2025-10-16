package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfo struct {
	Query        string  `json:"query"`
	Country      string  `json:"country"`
	RegionName   string  `json:"regionName"`
	City         string  `json:"city"`
	ISP          string  `json:"isp"`
	Org          string  `json:"org"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	Timezone     string  `json:"timezone"`
	CountryCode  string  `json:"countryCode"`
}

func main() {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		fmt.Println("❌ Failed to fetch IP info:", err)
		return
	}
	defer resp.Body.Close()

	var info IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Println("❌ Failed to parse response:", err)
		return
	}

	fmt.Println("\n🌐 IP Information")
	fmt.Println("---------------------------")
	fmt.Printf("Public IP     : %s\n", info.Query)
	fmt.Printf("Country       : %s (%s)\n", info.Country, info.CountryCode)
	fmt.Printf("Region / City : %s, %s\n", info.RegionName, info.City)
	fmt.Printf("ISP / Org     : %s / %s\n", info.ISP, info.Org)
	fmt.Printf("Location      : %.4f, %.4f\n", info.Lat, info.Lon)
	fmt.Printf("Timezone      : %s\n", info.Timezone)
}

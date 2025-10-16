package main

import (
	"fmt"
	"net/http"
	"time"
)

type Site struct {
	URL    string
	Status string
}

func checkSite(site *Site) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(site.URL)
	if err != nil {
		site.Status = "❌ DOWN"
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		site.Status = "✅ UP"
	} else {
		site.Status = fmt.Sprintf("⚠️ %d", resp.StatusCode)
	}
}

func main() {
	sites := []Site{
		{"https://google.com", ""},
		{"https://github.com", ""},
		{"https://supernovabusiness.in", ""},
	}

	fmt.Println("🌐 Website Uptime Monitor")
	fmt.Println("----------------------------")

	for {
		for i := range sites {
			checkSite(&sites[i])
			fmt.Printf("%-25s %s\n", sites[i].URL, sites[i].Status)
		}
		fmt.Println("----------------------------")
		time.Sleep(15 * time.Second) // repeat every 15 seconds
	}
}

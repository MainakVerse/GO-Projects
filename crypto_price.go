package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PriceResponse map[string]map[string]float64

func main() {
	var coin, currency string
	fmt.Print("Enter crypto symbol (e.g. bitcoin, ethereum): ")
	fmt.Scanln(&coin)
	fmt.Print("Enter currency (e.g. usd, inr, eur): ")
	fmt.Scanln(&currency)

	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
		strings.ToLower(coin),
		strings.ToLower(currency),
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Network error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("⚠️ API error:", resp.Status)
		return
	}

	var data PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("❌ Error parsing response:", err)
		return
	}

	price := data[strings.ToLower(coin)][strings.ToLower(currency)]
	if price == 0 {
		fmt.Println("⚠️ Invalid coin or currency.")
		return
	}

	fmt.Printf("\n💸 %s price: %.2f %s\n", strings.Title(coin), price, strings.ToUpper(currency))
}

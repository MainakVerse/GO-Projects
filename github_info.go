package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubUser struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	PublicRepos int  `json:"public_repos"`
	Followers int   `json:"followers"`
	Following int   `json:"following"`
	Bio       string `json:"bio"`
	HTMLURL   string `json:"html_url"`
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanln(&username)

	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("⚠️ User not found or API error.")
		return
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("❌ Error parsing response:", err)
		return
	}

	fmt.Println("\n👤 GitHub User Info")
	fmt.Println("---------------------------")
	fmt.Printf("Username   : %s\n", user.Login)
	fmt.Printf("Name       : %s\n", user.Name)
	fmt.Printf("Company    : %s\n", user.Company)
	fmt.Printf("Location   : %s\n", user.Location)
	fmt.Printf("Repos      : %d\n", user.PublicRepos)
	fmt.Printf("Followers  : %d | Following: %d\n", user.Followers, user.Following)
	fmt.Printf("Bio        : %s\n", user.Bio)
	fmt.Printf("Profile URL: %s\n", user.HTMLURL)
}

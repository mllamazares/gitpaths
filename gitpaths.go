package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)
const yellow = "\033[33m"
const reset = "\033[0m"

const banner = yellow + `
        _ __            __  __     
  ___ _(_) /____  ___ _/ /_/ /  ___
 / _ \/ / __/ _ \/ _ \/ __/ _ \(_-<
 \_, /_/\__/ .__/\_,_/\__/_//_/___/
/___/     /_/                     ` + reset + ` v0.0.1

Made by https://linkedin.com/in/mllamazares

---

`

type ApiResponse struct {
	Tree []struct {
		Path string `json:"path"`
	} `json:"tree"`
}

func main() {
	repoUrl := flag.String("u", "", "GitHub repository URL")
	outputFile := flag.String("o", "", "Output file (optional)")
	branch := flag.String("b", "master", "Branch name (optional, default: master)")
	silent := flag.Bool("silent", false, "Omit banner and sysout printing")
	help := flag.Bool("help", false, "Display help")

	flag.Parse()

	fmt.Print(banner)

	// Display help if -h is provided
	if *help {
		flag.Usage()
		return
	}

	// Validate input
	if *repoUrl == "" {
		fmt.Println("Repository URL is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Extract owner and repo name from URL provided
	parts := strings.Split(strings.TrimPrefix(*repoUrl, "https://github.com/"), "/")
	if len(parts) != 2 {
		fmt.Println("Invalid repository URL format.")
		os.Exit(1)
	}
	owner := parts[0]
	repo := parts[1]

	// Get the GitHub API URL
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", owner, repo, *branch)

	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to retrieve data: %s\n", resp.Status)
		os.Exit(1)
	}

	// Parse the JSON response
	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// Create the output file if defined
	var file *os.File
	if *outputFile != "" {
		file, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	// Print to sysout and optionally write the paths to the file
	for _, item := range apiResponse.Tree {
		if !*silent {
			fmt.Println(item.Path)
		}
		if file != nil {
			_, err := file.WriteString(item.Path + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				os.Exit(1)
			}
		}
	}
}

# gitpaths

![](https://img.shields.io/badge/license-MIT-green)
[![](https://img.shields.io/badge/LinkedIn-0077B5?logo=linkedin&logoColor=white)](https://www.linkedin.com/in/mllamazares/)
[![Watch on GitHub](https://img.shields.io/github/watchers/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/watchers)
[![Star on GitHub](https://img.shields.io/github/stars/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/stargazers)
[![Tweet](https://img.shields.io/twitter/url/https/github.com/mllamazares/gitpaths.svg?style=social)](https://twitter.com/intent/tweet?text=Check%20out%20gitpaths%21%20https%3A%2F%2Fgithub.com%2Fmllamazares%2Fgitpaths)

`gitpaths`  a lightweight tool written in Go that retrieves the folder structure of a GitHub repository without needing to clone it, making it easier to create custom wordlists for fuzzing. Check out the [sample workflow](#sample-workflow).

![gitpaths demo screenshot](demo.png)

## Features

- Fetches file paths from a repository using the GitHub API. 🔭
- Supports specifying a branch within the target repository. 🎯
- Minimal dependencies and lightweight. 🪶
- Blazing fast! 🚀

## Installation

`gitpaths` requires Go 1.21 or later to install successfully. Simply run the following command to get it:

```bash
go install -v github.com/mllamazares/gitpaths@latest
```

## Usage

```bash
gitpaths -h
```

This will display help for the tool. Here are all the parameters it supports:

```
Usage of gitpaths:
  -u string
    	GitHub repository URL
  -b string
    	Branch name (optional) (default "master")
  -o string
    	Output file (optional)
  -silent
    	Omit sysout printing
  -h	Display help
```

### Sample Workflow

1. Use a tool like [Wappalyzer](https://www.wappalyzer.com/), [httpx](https://github.com/projectdiscovery/httpx), [nuclei](https://github.com/projectdiscovery/nuclei), or [whatweb](https://github.com/urbanadventurer/WhatWeb) to detect the tech stack of a given endpoint, such as a web app or WordPress plugin.
2. Search for the corresponding GitHub repository and branch (if applicable) using Google.
3. Create a custom wordlist with the command: `gitpaths -u https://github.com/example/plugin -b version2.1.23 -o plugin_wordlist.txt`.
4. Finally, use your favorite fuzzing tool with that wordlist to test which endpoints are reachable on your target: `ffuf -u https://target.com/FUZZ -w plugin_wordlist.txt`.

## TODO
- [ ] Auto-detect technology and GitHub repo.
- [ ] Accept GitHub token to avoid rate limiting so it can be run at scale. 
- [ ] Mix with `ffuf` for an end-to-end fuzzing experience.
- [ ] Clean and refactor code.

Feel free to send a PR! 
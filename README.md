# 🗺️ gitpaths

![License: MIT](https://img.shields.io/badge/license-MIT-green)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?logo=linkedin&logoColor=white)](https://www.linkedin.com/in/mllamazares/)
[![Watch on GitHub](https://img.shields.io/github/watchers/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/watchers)
[![Star on GitHub](https://img.shields.io/github/stars/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/stargazers)
[![Tweet](https://img.shields.io/twitter/url/https/github.com/mllamazares/gitpaths.svg?style=social)](https://twitter.com/intent/tweet?text=Check%20out%20gitpaths%21%20https%3A%2F%2Fgithub.com%2Fmllamazares%2Fgitpaths)

`gitpaths` is a lightweight tool written in Go that lists the folder structure of a GitHub repo without requiring a clone. It simplifies the process of creating custom wordlists for fuzzing.

![gitpaths demo screenshot](demo.png)

## Features

- Fetches file paths from a repository using the GitHub API. 🔭
- Allows specifying a branch within the target repository. 🎯
- Minimal dependencies and lightweight. 🪶
- Extremely fast! 🚀

## Installation

`gitpaths` requires Go 1.21 or later. To install, simply run:

```bash
go install -v github.com/mllamazares/gitpaths@latest
```

## Usage

```bash
gitpaths -h
```

This command displays help for the tool. Here are the available parameters:

```
Usage of gitpaths:
  -u string
    	GitHub repository URL
  -b string
    	Branch name (optional) (default "master")
  -o string
    	Output file (optional)
  -silent
    	Suppress output to the terminal
  -h	Display help
```

### Sample Workflow

1. Use a tool like Wappalyzer, httpx, nuclei, or whatweb to identify the tech stack of the target endpoint, such as a CRM, CMS, or WordPress plugin.
2. Find the corresponding GitHub repository (and branch, if needed) using Google.
3. Generate a custom wordlist: `gitpaths -u https://github.com/example/pluginXYZ -b version2.1.23 -o pluginXYZ_wordlist.txt`
4. Use your preferred fuzzing tool with that wordlist to test which endpoints are reachable on your target: `ffuf -u https://target.com/plugins/FUZZ -w pluginXYZ_wordlist.txt`

## TODO
- [ ] Auto-detect technology and GitHub repo.
- [ ] Accept GitHub tokens to avoid rate limiting for large-scale operations.
- [ ] Integrate with `ffuf` for a complete fuzzing experience.
- [ ] Clean and refactor code.

Contributions are welcome! Feel free to submit a PR. 🙌

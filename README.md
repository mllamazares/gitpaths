# gitpaths

![](https://img.shields.io/badge/license-MIT-green)
[![](https://img.shields.io/badge/LinkedIn-0077B5?logo=linkedin&logoColor=white)](https://www.linkedin.com/in/mllamazares/)
[![Watch on GitHub](https://img.shields.io/github/watchers/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/watchers)
[![Star on GitHub](https://img.shields.io/github/stars/mllamazares/gitpaths.svg?style=social)](https://github.com/mllamazares/gitpaths/stargazers)
[![Tweet](https://img.shields.io/twitter/url/https/github.com/mllamazares/gitpaths.svg?style=social)](https://twitter.com/intent/tweet?text=Check%20out%20gitpaths%21%20https%3A%2F%2Fgithub.com%2Fmllamazares%2Fgitpaths)

`gitpaths` is a lightweight tool written in Go that retrieves the folder structure of a GitHub repository without needing to clone the entire content.

It's perfect for quickly creating custom wordlists for fuzzing open-source apps or plugins, helping you identify which endpoints in the default repository are accessible.

![gitpaths demo screenshot](demo.png)

## Features

- Fetches file paths from a repository using the GitHub API.
- No API key required.
- Supports specifying a branch within the target repository.
- Minimal dependencies and lightweight.
- Blazing fast! 🚀

## Installation

`gitpaths` requires Go 1.23 or later to install successfully. Simply run the following command to get it:

```shell
go install -v github.com/mllamazares/gitpaths@latest
```

## Usage

```
gitpaths -h
```

This will display help for the tool. Here are all the parameters it supports:

```
Usage of gitpaths:
  -b string
    	Branch name (optional) (default "master")
  -h	Display help
  -o string
    	Output file (optional)
  -silent
    	Omit banner and sysout printing
  -u string
    	GitHub repository URL
```

### Examples

To create a wordlist, run the following command:

```
gitpaths -u https://github.com/example/plugin -o plugin_wordlist.txt
```

You can then use that wordlist to test which endpoints are reachable on your target:

```
ffuf -u https://supersecret.com/FUZZ -w plugin_wordlist.txt
```

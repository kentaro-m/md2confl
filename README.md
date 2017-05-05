# md2confl
[![CircleCI](https://circleci.com/gh/kentaro-m/md2confl/tree/master.svg?style=shield)](https://circleci.com/gh/kentaro-m/md2confl/tree/master)
[![GitHub release](https://img.shields.io/github/release/kentaro-m/md2confl.svg)](https://github.com/kentaro-m/md2confl/releases)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/kentaro-m/md2confl/blob/master/LICENSE)

md2confl is a CLI tool to convert the markdown text to confluence wiki format.

## Demo
![](./demo.gif)

## Installation

### Homebrew
```
$ brew tap kentaro-m/homebrew-md2confl
$ brew install md2confl
```

### Golang
```
$ go get github.com/kentaro-m/md2confl
```

## Usage
```
$ md2confl help
md2confl - Convert markdown text to confluence wiki

Usage:
  md2confl [command]

Available Commands:
  copy        Copy a confluence wiki text converted from markdown.
  preview     Show a confluence wiki text converted from markdown.
  version     Print the version number of md2confl
  help        Help about any command

Flags:
  -h, --help   help for md2confl

Use "md2confl [command] --help" for more information about a command.
```

## Example
```
$ md2confl copy ~/sample.md
Copied to clipboard!

$ md2confl preview ~/sample.md
h1. Hello World

{code:language=go}package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
{code}
```

## License
MIT

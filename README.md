# go-dotenv
[![CircleCI](https://circleci.com/gh/bmf-san/go-dotenv/tree/master.svg?style=svg)](https://circleci.com/gh/bmf-san/go-dotenv/tree/master)

A dotenv built with golang.

# Installation
`go get github.com/bmf-san/go-dotenv`

# Get Started
Add a .env file in your project.

```
FOO=foo
BAR=bar
```

Then you can load env file like this. 

```golang
package main

import (
	"fmt"
	"os"

	"github.com/bmf-san/go-dotenv"
)

func main() {
	if err := dotenv.LoadEnv(); err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println(os.Getenv("FOO")) // foo
	fmt.Println(os.Getenv("BAR")) // bar
}
```
# simple go line notify

[![GoDoc](https://godoc.org/github.com/juunini/simple-go-line-notify?status.svg)](https://godoc.org/github.com/juunini/simple-go-line-notify)
[![Build Status](https://travis-ci.org/juunini/simple-go-line-notify.svg?branch=master)](https://travis-ci.org/juunini/simple-go-line-notify)
[![Go Report Card](https://goreportcard.com/badge/github.com/juunini/simple-go-line-notify)](https://goreportcard.com/report/github.com/juunini/simple-go-line-notify)

## Usage

```go
package main

import "github.com/juunini/simple-go-line-notify/notify"

func main() {
    bearer := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
    message := "hello, world!"

    if err := notify.Send(bearer, message); err != nil {
        panic(err)
    }
}
```

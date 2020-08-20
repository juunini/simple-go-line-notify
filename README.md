# simple go line notify

[![PkgGoDev](https://pkg.go.dev/badge/github.com/juunini/simple-go-line-notify/notify)](https://pkg.go.dev/github.com/juunini/simple-go-line-notify/notify)
[![Build Status](https://travis-ci.org/juunini/simple-go-line-notify.svg?branch=master)](https://travis-ci.org/juunini/simple-go-line-notify)
[![Go Report Card](https://goreportcard.com/badge/github.com/juunini/simple-go-line-notify)](https://goreportcard.com/report/github.com/juunini/simple-go-line-notify)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Usage

```go
package main

import "github.com/juunini/simple-go-line-notify/notify"

func main() {
    accessToken := "29jWoO****p70eK3AKA********ooHfusvDP6***ZmR"
    message := "hello, world!"

    if err := notify.SendText(accessToken, message); err != nil {
        panic(err)
    }
}
```

### Send Notify With Image

```go
accessToken := "your-access-token"
message := "your message. must not be empty"
imageURL := "image url. ex) https://..."

if err := notify.SendImage(accessToken, message, imageURL); err != nil {
    panic(err)
}
```
# simple golang line notify

[![PkgGoDev](https://pkg.go.dev/badge/github.com/juunini/simple-go-line-notify/notify)](https://pkg.go.dev/github.com/juunini/simple-go-line-notify/notify)
[![Build Status](https://travis-ci.org/juunini/simple-go-line-notify.svg?branch=master)](https://travis-ci.org/juunini/simple-go-line-notify)
[![Go Report Card](https://goreportcard.com/badge/github.com/juunini/simple-go-line-notify)](https://goreportcard.com/report/github.com/juunini/simple-go-line-notify)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Install

```bash
go get -u github.com/juunini/simple-go-line-notify
```

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

### Send Notify With Online Image

```go
accessToken := "your-access-token"
message := "your message. must not be empty"
imageURL := "image url. ex) https://..."

if err := notify.SendImage(accessToken, message, imageURL); err != nil {
    panic(err)
}
```

### Send Notify With Local Image (only jpg, png)
```go
accessToken := "your-access-token"
message := "your message. must not be empty"
imagePath := "/your/image/path.png"

if err := notify.SendLocalImage(accessToken, message, imagePath); err != nil {
    panic(err)
}
```

### Send Notify With Sticker

Sticker List is [See Here](https://devdocs.line.me/files/sticker_list.pdf)

```go
accessToken := "your-access-token"
message := "your message. must not be empty"
stickerPackageId := 1
stickerId := 113

if err := notify.SendSticker(accessToken, message, stickerPackageId, stickerId); err != nil {
    panic(err)
}
```

# Images (binding for Image-Magick)

[![Build Status](https://travis-ci.org/mantyr/images-imagick.svg?branch=master)](https://travis-ci.org/mantyr/images-imagick)
[![GoDoc](https://godoc.org/github.com/mantyr/images-imagick?status.png)](http://godoc.org/github.com/mantyr/images-imagick)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This not stable version

## Installation

    $ go get github.com/mantyr/golang-magick
    $ go get github.com/mantyr/images-imagick

## Example

```GO
package main

import (
    images "github.com/mantyr/images-imagick"
    "fmt"
)

func main() {
    address := "file.png"

    image, err := images.Open(address)
    if err != nil {
        fmt.Println(err)
        return
    }

    w := image.Width()
    h := image.Height()
    w = int(w/2)
    h = int(h/2)

    address = "file.gif"

    err = image.ResizeIn(w, h).SetGif().SetQuality(100).Save(address)
    if err != nil {
        fmt.Println(err)
        return
    }
    image.Dispose()  // It is very important, otherwise it will be a memory leak
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr

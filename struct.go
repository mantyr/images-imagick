package images

import (
    mag "github.com/mantyr/golang-magick"
)

type Image struct {
    Image       *mag.Image
    Address     string
    Format      string
    Quality     int
    width       int
    height      int
    Error       error
}
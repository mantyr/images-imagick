package images

import (
    mag "github.com/XuKaiqiang/golang-magick"
    "errors"
    "path"
    "os"
)

func Open(address string) (i *Image, err error) {
    i = new(Image)
    i.Address = address
    i.Quality = 100

    i.Image, i.Error = mag.DecodeFile(address)
    i.Format = i.Image.Format()
    return i, i.Error
}

func (i *Image) Resize(width int, height int) (image *Image) {
    image = new(Image)
    image.width   = width
    image.height  = height
    image.Quality = i.Quality

    image.Image, image.Error = i.Image.Resize(width, height, mag.FQuadratic)
    return
}

func (i *Image) Width() int {
    if i.Image == nil {
        return 0
    }
    return i.Image.Width()
}

func (i *Image) Height() int {
    if i.Image == nil {
        return 0
    }
    return i.Image.Height()
}

func (i *Image) ResizeIn(width int, height int) (image *Image) {
    var w_ratio float64 = float64(width)  / float64(i.Image.Width())
    var h_ratio float64 = float64(height) / float64(i.Image.Height())

    if w_ratio == h_ratio {
    } else if w_ratio < h_ratio {
        height = int(float64(i.Image.Height()) * w_ratio)
    } else {
        width  = int(float64(i.Image.Width()) * h_ratio)
    }

    image = i.Resize(width, height)

    image.width  = image.Image.Width()
    image.height = image.Image.Height()
    image.Quality = i.Quality
    return
}

func (i *Image) SetQuality(quality int) *Image {
    i.Quality = quality
    return i
}

func (i *Image) SetJpeg() *Image {
    i.Format = "JPEG"
    return i
}

func (i *Image) SetTiff() *Image {
    i.Format = "TIFF"
    return i
}

func (i *Image) SetPng() *Image {
    i.Format = "PNG"
    return i
}

func (i *Image) SetGif() *Image {
    i.Format = "GIF"
    return i
}

func (i *Image) Save(address string) (err error) {
    if i.Image == nil {
        if i.Error != nil {
            return i.Error
        }
        return errors.New("no image")
    }

    dir := path.Dir(address)
    if !IsDir(dir) {
        os.Mkdir(dir, 0777)
    }

    info := mag.NewInfo()
    info.SetFormat(i.Format)
    info.SetQuality(uint(i.Quality))

    file, err := os.Create(address)
    defer file.Close()
    if err != nil {
        return
    }
    err = i.Image.Encode(file, info)
    return
}
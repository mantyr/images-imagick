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

// todo
func (i *Image) ResizeIn(width int, height int) (image *Image) {
    max_width  := width
    max_haight := height

    image = i.Resize(max_width, max_haight)

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
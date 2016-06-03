package images

import (
    "os"
)

func IsDir(path string) (status bool) {
    info, err := os.Stat(path)
    if err != nil {
        return
    }
    return info.IsDir()
}
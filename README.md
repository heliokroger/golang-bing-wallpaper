# golang-bing-wallpaper
> Get the Bing wallpaper of the day using Golang
## Installing
To install, just run:
```bash
go get github.com/heliojuniorkroger/golang-bing-wallpaper
```
## Getting the wallpaper
Here a example code of how to get the wallpaper:
```go
package main

import (
    "fmt"
    "github.com/heliojuniorkroger/golang-bing-wallpaper"
)

func main() {
    wallpaper, err := bingWallpaper.Get(1, "pt-BR")
    if err != nil {
        panic(err)
    }
    fmt.Println(wallpaper)
}
```

# golang-bing-wallpaper
> Get the Bing wallpaper of the day using Golang
## Installing
To install, just run:
```bash
go get github.com/heliojuniorkroger/golang-bing-wallpaper
```
## Getting the wallpaper
You should use the **bingWallpaper.Get()** function to get the wallpaper, there are two parameters that you need to pass to the function to make it work. The first parameter is how many wallpapers do you want, and the second is the locale of the wallpaper. Here is some example of code: 
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

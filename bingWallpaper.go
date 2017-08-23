package bingWallpaper

import (
    "net/http"
    "encoding/json"
    "strconv"
)

type Wallpaper struct {
    Images []struct {
        Startdate string `json:"startdate"`
        Fullstartdate string `json:"fullstartdate"`
        Enddate string `json:"enddate"`
        Urlbase string `json:"urlbase"`
        Copyright string `json:"copyright"`
        Copyrightlink string `json:"copyrightlink"`
        Quiz string `json:"quiz"`
        Wp string `json:"wp"`
        Hsh string `json:"hsh"`
        Drk string `json:"drk"`
        Top string `json:"top"`
        Bot string `json:"bot"`
        Hs string `json:"hs"`
        Url string `json:"url"`
    } `json:"images"`
    Tooltips []struct {
        Loading string `json:"loading"`
        Previous string `json:"previous"`
        Next string `json:"next"`
        Walle string `json:"walle"`
        Walls string `json:"walls"`
    } `json:"tooltips"`
}

func Get(quantity int, locale string) (Wallpaper, error) {
    res, err := http.Get("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=" + strconv.Itoa(quantity) + "&mkt=" + locale)
    var wallpaper Wallpaper
    json.NewDecoder(res.Body).Decode(&wallpaper)
    return wallpaper, err
}

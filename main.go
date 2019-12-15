package main

import (
	"github.com/antonholmquist/jason"
	"github.com/gin-gonic/gin"
	"github.com/shinshin86/go-nconfig"
	"net/http"
)

func getWeatherJson() []*jason.Object {
	config := nconfig.New("default")
	apiKey := config.Get("apikey")

	cityId := "524901"
	url := "http://api.openweathermap.org/data/2.5/forecast?id=" + cityId + "&APPID=" + apiKey

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	root, err := jason.NewObjectFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	data, err := root.GetObjectArray("list")
	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(c *gin.Context) {
		data := getWeatherJson()
		c.HTML(200, "index.html", gin.H{"data": data})
	})
	router.Run()
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"net/url"
)

type PageValues struct {
	Title string
	Body  string
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.Static("/assets", "./assets")

	r.NoRoute(notFoundPage)

	r.GET("/", home)
	r.GET("/searchGopher", getGopher)

	r.Run(":9999")
}

func home(c *gin.Context) {

	pageValues := PageValues{
		Title: "Bowery Golang Demo",
		Body:  "May 4th be with you",
	}

	c.HTML(http.StatusOK, "home.html", pageValues)
}

func getGopher(c *gin.Context) {
	URL, err := url.Parse("https://pixabay.com/api")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

}

func notFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

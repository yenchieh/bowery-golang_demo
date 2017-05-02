package main

import (
	"net/http"

	"fmt"
	"net/url"

	"encoding/json"

	"io/ioutil"

	"os"

	"github.com/gin-gonic/gin"
)

type PixbayImageResponse struct {
	TotalHits   int64
	SearchQuery string
	Hits        []ImageDetail
}

type ImageDetail struct {
	PreviewHeight   int64
	Likes           int64
	Favorites       int64
	Tags            string
	WebformatHeight int64
	PreviewWidth    int64
	Comments        int64
	downloads       int64
	PageURL         string
	PreviewURL      string
	WebformateURL   string
	ImageWidth      int64
	UserID          int64
	User            string
	Type            string
	ID              int64
	UserImageURL    string
	ImageHeight     int64
}

var key string

func main() {

	fmt.Printf("%#v", os.Args)
	key = os.Args[1]

	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.Static("/assets", "./assets")

	r.NoRoute(notFoundPage)

	r.GET("/", home)
	r.GET("/search", searchImages)

	r.Run(":9999")
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func searchImages(c *gin.Context) {
	URL, err := url.Parse("https://pixabay.com/api")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

	q := c.Query("q")
	if q == "" {
		q = "gopher"
	}
	var req *http.Request
	form := &url.Values{
		"key":  {key},
		"q":    {q},
		"lang": {"en"},
	}

	urlWithQuery := fmt.Sprintf("%s?%s", URL.String(), form.Encode())

	req, err = http.NewRequest(http.MethodGet, urlWithQuery, nil)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var images PixbayImageResponse

	err = json.Unmarshal(body, &images)
	if err != nil {
		fmt.Println(err)
		return
	}

	images.SearchQuery = q

	c.HTML(http.StatusOK, "home.html", images)
}

func notFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}

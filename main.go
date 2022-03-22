package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

type Sector struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imgUrl"`
	Routes      []Route `json:"routes"`
}

type Route struct {
	Name        string `json:"name"`
	Length      string `json:"length"`
	Grade       string `json:"grade"`
	Description string `json:"description"`
}

func loadDataFromJSON(filePath string) ([]*Sector, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sectors []*Sector
	err = json.NewDecoder(file).Decode(&sectors)

	return sectors, err
}

func makeTemplates() multitemplate.Render {
	templates := multitemplate.New()

	templates.AddFromFiles("index",
		"web/base.html",
		"web/index.html",
		"web/header.html",
		"web/footer.html",
	)

	templates.AddFromFiles("about",
		"web/base.html",
		"web/about.html",
		"web/header.html",
		"web/footer.html",
	)

	return templates
}

func main() {

	sectors, err := loadDataFromJSON("data.json")
	if err != nil {
		panic(err)
	}

	for _, v := range sectors {
		fmt.Println(v.Name)
	}

	router := gin.Default()

	// If we are developing, reload templates in each request
	if gin.Mode() == "debug" {
		router.Use(func(c *gin.Context) {
			router.HTMLRender = makeTemplates()
		})
	}

	// Static Folder
	router.StaticFS("/static", http.Dir("static"))
	router.HTMLRender = makeTemplates()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"Title":   "Main website",
			"sectors": sectors,
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{
			"Title": "About",
		})
	})

	router.Run()
}

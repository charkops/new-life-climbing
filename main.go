package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

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
			"Title": "Main website",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about", gin.H{
			"Title": "About",
		})
	})

	router.Run()
}

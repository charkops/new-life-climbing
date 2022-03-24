package main

import (
	"climbing/multitemplate"
	"climbing/util"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Sector struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imgUrl"`
	Slug        string  `json:"slug"`
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

// Creates the templates from files
func makeTemplates() multitemplate.Multitemplate {
	templates := multitemplate.New()

	templates.AddFromFiles("index",
		"web/base.html",
		"web/index.html",
		"web/header.html",
		"web/footer.html",
	)

	templates.AddFromFiles("sector",
		"web/base.html",
		"web/sector.html",
		"web/header.html",
		"web/footer.html",
	)

	templates.AddFromFiles("about",
		"web/base.html",
		"web/about.html",
		"web/header.html",
		"web/footer.html",
	)

	templates.AddFromFiles("404",
		"web/base.html",
		"web/404.html",
		"web/header.html",
		"web/footer.html",
	)

	return templates
}

func executeTemplate(templates multitemplate.Multitemplate, templateName string, outName string, data interface{}) {
	tmpl, exists := templates[templateName]
	if !exists {
		fmt.Errorf("template %s doesn't exist", templateName)
	}

	if !strings.HasSuffix(outName, ".html") {
		outName = outName + ".html"
	}

	f, err := os.Create("dist/" + outName)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func createDistFolder() {
	// Create folder
	os.RemoveAll("dist")
	err := os.Mkdir("dist", 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("dist/sector", 0755)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("dist/static", 0755)
	if err != nil {
		panic(err)
	}
	util.CopyFolder("static", "dist/static")
}

func main() {

	sectors, err := loadDataFromJSON("data.json")
	if err != nil {
		panic(err)
	}

	createDistFolder()

	templates := makeTemplates()
	executeTemplate(templates, "index", "index", struct {
		Title   string
		Sectors []*Sector
	}{
		"Index",
		sectors,
	})

	executeTemplate(templates, "about", "about", struct {
		Title string
	}{
		"About",
	})

	for _, v := range sectors {
		executeTemplate(templates, "sector", "sector/"+v.Name, struct {
			Title  string
			Sector Sector
		}{
			v.Name,
			*v,
		})
	}

	executeTemplate(templates, "404", "404", nil)
}

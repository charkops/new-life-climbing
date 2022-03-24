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

	return templates
}

// Executes a specific template with context
func executeTemplate(templates multitemplate.Multitemplate, name string, data interface{}) {
	tmpl, exists := templates[name]
	if !exists {
		fmt.Errorf("template %s doesn't exist", name)
	}

	if !strings.HasSuffix(name, ".html") {
		name = name + ".html"
	}

	f, err := os.Create("dist/" + name)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func executeTemplateWithDifferentOutName(templates multitemplate.Multitemplate, name string, outName string, data interface{}) {
	tmpl, exists := templates[name]
	if !exists {
		fmt.Errorf("template %s doesn't exist", name)
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

func main() {

	sectors, err := loadDataFromJSON("data.json")
	if err != nil {
		panic(err)
	}

	// Create folder
	os.RemoveAll("dist")
	err = os.Mkdir("dist", 0755)
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

	templates := makeTemplates()
	executeTemplate(templates, "index", struct {
		Title   string
		Sectors []*Sector
	}{
		"Index",
		sectors,
	})

	executeTemplate(templates, "about", struct {
		Title string
	}{
		"About",
	})

	for _, v := range sectors {
		executeTemplateWithDifferentOutName(templates, "sector", "sector/"+v.Name, struct {
			Title  string
			Sector Sector
		}{
			v.Name,
			*v,
		})
	}
}

package main

import (
	"awesomeProject/routes"
	"awesomeProject/session"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func main() {
	//fmt.Println("kek")
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))
	m.Get("/", routes.IndexHandler)
	staticoptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticoptions))

	dbsession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	var memorySession *session.Session
	memorySession = session.NewSession()
	m.Map(memorySession)
	db := dbsession.DB("blog3")
	m.Map(db)
	m.Get("/write", routes.WriteHandler)
	m.Get("/login", routes.GetLoginHandler)
	m.Post("/login", routes.PostLoginHandler)
	m.Get("/edit", routes.EditHandler)
	m.Get("/delete", routes.DeleteHandler)
	m.Post("/SavePost", routes.SavePostHandler)

	m.Run()
}

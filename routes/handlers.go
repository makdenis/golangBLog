package routes

import (
	"awesomeProject/db/documents"
	"awesomeProject/models"
	"awesomeProject/session"
	"awesomeProject/utils"
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
)

func IndexHandler(rnd render.Render, r *http.Request, db *mgo.Database, memorySession *session.Session) {
	postCollection := db.C("posts")
	cookie, _ := r.Cookie(session.COOKIE)
	if cookie != nil {
		fmt.Println(memorySession.Get(cookie.Value))
	} else {

	}
	postDocuments := []documents.PostDocument{}
	postCollection.Find(nil).All(&postDocuments)
	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.Content}
		posts = append(posts, post)
	}
	//fmt.Println(posts)
	rnd.HTML(200, "index", posts)
}

func SavePostHandler(rnd render.Render, r *http.Request, db *mgo.Database) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	postCollection := db.C("posts")
	postDocument := documents.PostDocument{id, title, content}
	if id != "" {

		postCollection.UpdateId(id, postDocument)

	} else
	{
		id := utils.GenerateId()
		postDocument.Id = id
		postCollection.Insert(postDocument)
	}

	rnd.Redirect("/")
}

func DeleteHandler(rnd render.Render, r *http.Request, db *mgo.Database) {

	id := r.FormValue("id")
	postCollection := db.C("posts")
	if id == "" {
		rnd.Redirect("/")
	} else {
		//var tmp int
		//err := postCollection.Find(bson.M{"id": id}).Select(bson.M{"_id": nil}).One(&tmp)

		postCollection.RemoveId(id)
	}
	rnd.Redirect("/")

}

func EditHandler(rnd render.Render, r *http.Request, db *mgo.Database) {
	postCollection := db.C("posts")
	id := r.FormValue("id")
	//fmt.Println(id)
	postDocument := documents.PostDocument{}

	err := postCollection.FindId(id).One(&postDocument)

	//fmt.Println(postDocument.Id)
	if err != nil {
		rnd.Redirect("/")
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.Content}
	//fmt.Println(post)
	rnd.HTML(200, "write", post)
}

func WriteHandler(rnd render.Render) {
	//fmt.Fprint(w, "Hi")
	rnd.HTML(200, "write", nil)
}

package routes

import (
	"awesomeProject/session"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

func PostLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter, memorySession *session.Session) {
	//fmt.Fprint(w, "Hi")
	user:=r.FormValue("username")
	//pass:=r.FormValue("password")
	sessionId:=memorySession.Init(user)
	cookie:= &http.Cookie{
		Name:session.COOKIE,
		Value:sessionId,
		Expires:time.Now().Add(5*time.Minute),
	}
	http.SetCookie(w, cookie)
	rnd.Redirect("/")
}


func GetLoginHandler(rnd render.Render) {
	//fmt.Fprint(w, "Hi")
	rnd.HTML(200, "login", nil)

}

package utility

import (
	"os"
	"text/template"

	//"log"
	//"fmt"

	"net/http"

	"github.com/gorilla/sessions"
)

// Template Pool
var View *template.Template

// Session Store
var Store *sessions.FilesystemStore

type Flash struct {
	Type    string
	Message string
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, template string, data interface{}) {
	session, _ := Store.Get(r, os.Getenv("SESSION_NAME"))
	tmplData := make(map[string]interface{})
	tmplData["data"] = data
	tmplData["flash"] = viewFlash(w, r)
	tmplData["session"] = map[string]interface{}{"email": session.Values["email"], "user_id": session.Values["user_id"], "type": session.Values["type"], "profile_log": session.Values["profile_log"]}
	tmplData["config"] = map[string]interface{}{"APPURL": os.Getenv("APPURL")}
	View.ExecuteTemplate(w, template, tmplData)
}

func viewFlash(w http.ResponseWriter, r *http.Request) interface{} {
	session, err := Store.Get(r, os.Getenv("SESSION_NAME"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fm := session.Flashes("message")
	if fm == nil {
		return nil
	}
	session.Save(r, w)
	return fm
}

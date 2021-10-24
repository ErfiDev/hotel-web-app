package utils

import (
	"html/template"
	"net/http"
	"os"
)

func RenderTemplate(w http.ResponseWriter , tmpl string , data interface{}) {
	tmp , _ := template.ParseFiles("./views/" + tmpl , "./views/base.layout.gohtml")
	tmp.Execute(w , data)
	tmp.Execute(os.Stdout , nil)
}
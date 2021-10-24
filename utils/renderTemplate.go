package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter , tmpl string , data interface{}) {
	tmp , _ := template.ParseFiles("./views/" + tmpl , "./views/base.layout.gohtml")
	tmp.Execute(w , data)
}
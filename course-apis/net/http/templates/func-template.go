package main

import (
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"dateOnly": timeformat,
	"kitchen":  tf_kitchen,
}

func main() {

	tpl := template.Must(template.New("").Funcs(fm).ParseGlob("files/*"))

	tpl.ExecuteTemplate(os.Stdout, "huit.gohtml", time.Now())
}

func timeformat(t time.Time) string {
	return t.Format(time.DateOnly)
}

func tf_kitchen(t time.Time) string {
	return t.Format(time.Kitchen)
}

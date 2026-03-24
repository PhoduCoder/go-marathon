package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"dateOnly": timeformat,
	"kitchen":  tf_kitchen,
	"upper":    strings.ToUpper,
}

func main() {

	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("files/*"))

	tpl.ExecuteTemplate(os.Stdout, "neuf.gohtml", time.Now())
}

func timeformat(t time.Time) string {
	return t.Format(time.DateOnly)
}

func tf_kitchen(t time.Time) string {
	return t.Format(time.Kitchen)
}

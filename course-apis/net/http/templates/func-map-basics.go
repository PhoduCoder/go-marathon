package main

import (
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
}

func main() {

	// The below doesn't work

	// tpl := template.Must(template.New("mytemp").Parse("This is the sample text {{ upper . }}"))
	// tpl = tpl.Funcs(fm)

	tpl := template.Must(template.New("mytemp").Funcs(fm).Parse("This is the sample text in upper case {{ upper . }} and this is in lower case {{ lower .}} \n "))
	tpl.ExecuteTemplate(os.Stdout, "mytemp", "Gaurav")

}

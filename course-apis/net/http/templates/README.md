Template/text package 

Two main functions 
a) tpl, err := template.ParseFiles("name of files")
It is a variadic function, that accepts one or more files
It returns a pointer to template, which is a container that holds all the files 

b) Next run tpl.Execute(writer, data)
you can pass a file since it implements writer interface 
or os.Stdout 

c) tpl.ExecuteTemplate( writer, name, data)
This takes writer juust like above, but the name of the template to write to, if we use
variadic params in ParseFiles and data 

Remember data is of type interface{}, i.e. any type 

Also in parseFiles, we have to pass all the names of the files.
Instead one can use parseGlobs where we just pass a pattern 
It takes a folder and any files inside that 

Eg, tpl, err := template.ParseGlobs("templates/*.gohtml" )

d) Sometimes you will also see inside the init function uusing a template.Must function. 

The code will loook like this 

func init(){
    tpl := template.Must(template.ParseGlob("templates/*"))
}

What template.Must takes is exactly what template.ParseGlob returns 
which is pointer to template holding all the templates and error

The templates.Must does the conditional error check for you and just returns the error.







====

```
In GoLang, the execution order is as below

a) Package level variables are defined
b) Init() functions execute (in the order they are defined)
Keep in mind, init() functions run only once before main.
They are automatically called by Go Runtime

One can define multiple init functions, they execute in the order how they were defined before the main 
```

package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutPath  = "views/layouts/"
	TemplateDir = "views/"
	TemplateExt = ".gohtml"
)

func NewView(Layout string, files ...string) *View {

	addTemplatePath(files)
	addTemplateExt(files)

	files = append(files, LayoutFiles()...)

	t, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{Template: t, Layout: Layout}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func LayoutFiles() []string {

	files, err := filepath.Glob(layoutPath + "*" + TemplateExt)

	if err != nil {
		panic(err)
	}

	return files
}

// addTemplatePath takes in a slice of strings
// representing the filepaths for templates, and it prepends
// the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if the TemplateDir == ""views/
func addTemplatePath(files []string) {
	for i, value := range files {
		files[i] = TemplateDir + value
	}
}

// addTemplateExt takes in a slice of strings
// representing the filepath for templates, and it appends
// the TemplateExt extension to each string in the slice
//
// Eg the input "{"home"} would result in the output
// {"home.gohtml"} if the TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, value := range files {
		files[i] = value + TemplateExt
	}
}

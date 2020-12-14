package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutPath = "views/layouts/"
	layoutExt  = ".gohtml"
)

func NewView(Layout string, files ...string) *View {

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

	files, err := filepath.Glob(layoutPath + "*" + layoutExt)

	if err != nil {
		panic(err)
	}

	return files
}

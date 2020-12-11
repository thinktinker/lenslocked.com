package views

import (
	"html/template"
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

func LayoutFiles() []string {

	files, err := filepath.Glob(layoutPath + "*" + layoutExt)

	if err != nil {
		panic(err)
	}

	return files
}

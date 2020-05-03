package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) //unpacking slice to string
	t, err := template.ParseFiles(files...) // here ... converts slice to a string
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {

	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

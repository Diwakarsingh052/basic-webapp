package views

import "html/template"

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...) // here ... converts slice to a string
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}

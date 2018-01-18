package view

import "net/http"

//Page elements to display and render in a page
type Page struct {
	Title string
	Menu  string
	Data  interface{}
}

//NewPage create a new page
func NewPage(Title, Menu string, Data interface{}) *Page {
	return &Page{
		Title, Menu, Data,
	}
}

//RenderTemplate renders a jet template with given name and data
func RenderTemplate(w http.ResponseWriter, name string, page *Page) error {

	vw, err := views.GetTemplate(name)
	if err != nil {
		//utils.Log(500, err.Error())
		return err
	}

	err = vw.Execute(w, nil, page)
	if err != nil {
		//utils.Log(500, err.Error())
	}
	//utils.Log(200, "rendered "+name)
	return err
}

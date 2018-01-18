package view

import "github.com/CloudyKit/jet"

var views = jet.NewHTMLSet("./view")

func init() {
	views.SetDevelopmentMode(true) //TODO: read from config and remove in production
	/*views.AddGlobal("Version", config.Version)
	views.AddGlobal("IsProduction", config.IsProduction)
	views.AddGlobal("URL", config.FullBaseURL)*/
}

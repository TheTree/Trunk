package controllers

import "github.com/robfig/revel"
import "strings"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func init() {
	revel.TemplateFuncs["active"] = func(s string) string {
		if strings.Contains(App.Request.URL.String(), "active") {
			return "active"
		}
		return ""
	}
}

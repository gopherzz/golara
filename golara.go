package golara

import (
	"net/http"
	"text/template"

	"github.com/gopherzz/golara/page"
	"github.com/gorilla/mux"
)

type CMS struct {
	router *mux.Router
}

func New() *CMS {
	return &CMS{
		router: mux.NewRouter(),
	}
}

func (c *CMS) renderPageComponents(p page.Page) string {
	res := ""
	for _, c := range p.Render() {
		res += c.Render()
	}
	return res
}

func (c *CMS) renderPage(route string, p page.Page) *template.Template {
	page := c.renderPageComponents(p)
	return template.Must(template.New(route).Parse(page))
}

func (c *CMS) RegisterPage(route string, page page.Page) {
	c.router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		templ := c.renderPage(route, page)
		templ.ExecuteTemplate(w, route, nil)
	})
}

func (c *CMS) Router() http.Handler {
	return c.router
}

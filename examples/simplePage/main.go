package main

import (
	"net/http"

	"github.com/gopherzz/golara"
	"github.com/gopherzz/golara/component"
	"github.com/gopherzz/golara/element"
)

type HelloComponent struct {
	name string
}

func (c HelloComponent) Render() string {
	return element.H1(element.A{}, "Hello, "+c.name)
}

type HelloPage struct{}

func (p HelloPage) Render() []component.Component {
	return []component.Component{
		&HelloComponent{"Gopherz"},
	}
}

func main() {
	c := golara.New()
	c.RegisterPage("/hello", HelloPage{})
	http.ListenAndServe(":7000", c.Router())
}

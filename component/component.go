package component

import "github.com/gopherzz/golara/element"

type Component interface {
	Render() string
}

type DefaultComponent struct {
}

func (c *DefaultComponent) Render() string {
	return element.Div(
		element.A{"class": "default"},
		element.H1(element.A{}, "Default Component"),
	)
}

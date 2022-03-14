package page

import "github.com/gopherzz/golara/component"

type Page interface {
	Render() []component.Component
}

type DefaultPage struct{}

func (p DefaultPage) Render() []component.Component {
	return []component.Component{
		&component.DefaultComponent{},
	}
}

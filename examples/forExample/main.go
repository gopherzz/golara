package main

import (
	"fmt"
	"net/http"

	"github.com/gopherzz/golara"
	"github.com/gopherzz/golara/component"
	"github.com/gopherzz/golara/element"
)

type User struct {
	name string
	age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s - %d years old", u.name, u.age)
}

type UsersComponent struct {
	users []User
}

func (c UsersComponent) Render() string {
	user := func(i int) string {
		return element.Div(
			element.A{
				"class": "user",
				"id":    c.users[i].name,
			},
			element.H1(element.A{}, c.users[i].String()),
		)
	}
	return element.Div(element.A{"class": "users"}, element.For(len(c.users), user))
}

type UsersPage struct{}

func (p UsersPage) Render() []component.Component {
	return []component.Component{
		UsersComponent{[]User{{"gopherzz", 17}, {"guest", 21}}},
	}
}

func main() {
	c := golara.New()
	c.RegisterPage("/users", UsersPage{})
	http.ListenAndServe(":7000", c.Router())
}

package element

import "testing"

func TestElement(t *testing.T) {
	tablecase := []struct {
		want string
		got  string
	}{
		{"<div><h1>test</h1></div>", Div(A{}, H1(A{}, "test"))},
		{"<div><h1>test</h1><div><h1>test2</h1></div></div>", Div(A{}, H1(A{}, "test"), Div(A{}, H1(A{}, "test2")))},
	}

	for _, tc := range tablecase {
		if tc.want != tc.got {
			t.Errorf("want: %s, got: %s", tc.want, tc.got)
		}
	}
}

func TestElement_pairElem(t *testing.T) {
	tablecase := []struct {
		want       string
		attributes A
		name       string
		content    string
	}{
		{"<h1>text</h1>", A{}, "h1", "text"},
		{"<p>content</p>", A{}, "p", "content"},
		{`<div class="user">content</div>`, A{"class": "user"}, "div", "content"},
	}

	for _, tc := range tablecase {
		if got := pairElem(tc.name, tc.attributes, tc.content); got != tc.want {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}

// func TestElement_joinMap(t *testing.T) {
// 	tablecase := []struct {
// 		want       string
// 		attributes map[string]string
// 	}{
// 		{` class="name" id="age"`, map[string]string{"class": "name", "id": "age"}},
// 	}

// 	for _, tc := range tablecase {
// 		if got := joinMap(tc.attributes); got != tc.want {
// 			t.Errorf("want: %s, got: %s", tc.want, got)
// 		}
// 	}
// }

func TestElement_h(t *testing.T) {
	tablecase := []struct {
		want    string
		level   int
		content []string
	}{
		{"<h1>text</h1>", 1, []string{"text"}},
		{"<h2>content\nmultiline</h2>", 2, []string{"content", "multiline"}},
	}

	for _, tc := range tablecase {
		if got := h(A{}, tc.level, tc.content...); got != tc.want {
			t.Errorf("want: %s, got: %s", tc.want, got)
		}
	}
}

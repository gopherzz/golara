package element

import (
	"fmt"
	"strings"
)

// A Attributes structure
type A map[string]string

func joinMap(m map[string]string) string {
	res := ""
	for k, v := range m {
		res += fmt.Sprintf(" %s=%q", k, v)
	}
	return res
}

func attr(attributes A) string {
	if len(attributes) == 0 {
		return ""
	}
	return joinMap(attributes)
}

func pairElem(name string, a A, content string) string {
	return fmt.Sprintf("<%s%s>%s</%s>", name, attr(a), content, name)
}

func Div(a A, els ...string) string {
	return pairElem("div", a, strings.Join(els, ""))
}

func h(a A, level int, content ...string) string {
	hWithLevel := fmt.Sprintf("h%d", level)
	return pairElem(hWithLevel, a, strings.Join(content, "\n"))
}

func H1(a A, els ...string) string {
	return h(a, 1, els...)
}

func H2(a A, els ...string) string {
	return h(a, 2, els...)
}

func H3(a A, els ...string) string {
	return h(a, 3, els...)
}

func H4(a A, els ...string) string {
	return h(a, 4, els...)
}

func H5(a A, els ...string) string {
	return h(a, 5, els...)
}

func H6(a A, els ...string) string {
	return h(a, 6, els...)
}

func For(count int, f func(i int) string) string {
	res := ""
	for i := 0; i < count; i++ {
		res += f(i)
	}
	return res
}

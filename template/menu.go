package template

import (
	"fmt"
	"strings"
)

type MenuBuilder struct {
	Process string
	Orders  []string
}

func (mb *MenuBuilder) Build() string {
	snippet := `tell application "System Events"
  tell process "%s"
    set frontmost to true
    %s
  end tell
end tell`

	var menus []string
	menuSnippet := `menu item "%s" of menu "%s"`
	for i := len(mb.Orders) - 1; i > 0; i-- {
		menus = append(menus, fmt.Sprintf(menuSnippet, mb.Orders[i], mb.Orders[i-1]))
	}
	if len(mb.Orders) == 1 {
		menus = append(menus, fmt.Sprintf("menu \"%s\"", mb.Orders[0]))
	}
	menus = append(menus, "menu bar 1")
	return fmt.Sprintf(snippet, mb.Process, "click "+strings.Join(menus, " of "))
}

func NewMenuBuilder(process string, orders []string) *MenuBuilder {
	return &MenuBuilder{
		Process: process,
		Orders:  orders,
	}
}

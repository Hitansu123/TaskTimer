package main

import(
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{
	tasks []string
	cursor int
	selected string
}

func newModel() model{
	return model{
		tasks: make([]string, 0),
		cursor: 0,
		selected: "",
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil
}

func (m model) View() string {
    return "Hello Bubble Tea"
}

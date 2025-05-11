package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	text "github.com/charmbracelet/bubbles/textinput"

)

type model struct{
	tasks []string
	cursor int
	selected string
	textInput text.Model
}

func newModel() model{
	t1:=text.New()
	t1.Placeholder="Enter Task"
	t1.Focus()
	t1.CharLimit=100
	t1.Width=20
	return model{
		tasks: make([]string, 0),
		cursor: 0,
		selected: "",
		textInput: t1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			input := m.textInput.Value()
			if input != "" {
				m.tasks = append(m.tasks, input)
				m.textInput.Reset()
			}
		}
	}

	// Always update the text input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
} 
func (m model) View() string {
	s := "Your Tasks:\n"
	for i, task := range m.tasks {
		s += fmt.Sprintf("%d. %s\n", i+1, task)
	}
	s += "\n" + m.textInput.View()
	s += "\n(press Enter to add task, q to quit)"
	return s

}

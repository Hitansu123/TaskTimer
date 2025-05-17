package main

import (
	"fmt"
	"time"

	text "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	//lipgloss "github.com/charmbracelet/lipgloss"
)

type model struct{
	tasks []string
	cursor int
	selected string
	textInput text.Model
	screen int
}
type timer struct{
	start string
	duration string
}

func newModel() *model{
	t1:=text.New()
	t1.Placeholder="Enter Task"
	t1.Focus()
	t1.CharLimit=100
	t1.Width=20
	return &model{
		tasks: make([]string, 0),
		cursor: 0,
		selected: "",
		textInput: t1,
		screen: 0,
	}
}
func (t timer) startTimer() *timer{
	s:=time.Now()
	newTimer:=timer{
		start: time.Now().Format("15:04:05"),
		duration: time.Since(s).String(),
	}
	return &newTimer
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
		case "s":
			m.textInput.Blur()
			if m.cursor < len(m.tasks) {
				m.selected = m.tasks[m.cursor]
			}
		case "t":
			if m.selected==""{
				//fmt.Println("First select something")
			}else{
				//t:=timer{}
				//storeTime:=t.startTimer()
				m.screen=1
			}
		case "esc":
			m.textInput.Focus()
		case "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "k","up":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}

	// Always update the text input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
} 
func (m model) View() string {
	if m.screen==1{
		t:=timer{}
		storeTime:=t.startTimer()
		text:=fmt.Sprintf("Timer started %v \n Duration is %v",storeTime.start,storeTime.duration)
		return text
	}
	s := "Your Tasks:\n"
	for i, task := range m.tasks {
		cursor := "  "
		if m.cursor == i {
			cursor = "->"
		}

		selected := ""
		if m.selected == task {
			selected = " [selected]"
		}

		s += fmt.Sprintf("%s %d. %s%s\n", cursor, i+1, task, selected)

	}
	//right:="Timer"

	//input:=m.textInput.View()
	//s+=fmt.Sprintf("%-30s%s\n",s,right)
	s += "\n" + m.textInput.View()
	s += "\n(press Enter to add task, s to select, q to quit)"
	return s

}

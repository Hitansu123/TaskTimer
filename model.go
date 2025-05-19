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
	startTime time.Time
	duration time.Duration
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
		startTime: time.Now(),
		duration: 0,
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
		case "s":
			m.textInput.Blur()
			if m.cursor < len(m.tasks) {
				m.selected = m.tasks[m.cursor]
			}
		case "t":
			if m.selected!=""{
				//mt.Println
				m.screen = 1
				m.startTime= time.Now() // Add this field to your model to track time
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
					return t
			})

			}else{
				//t:=timer{}
				//storeTime:=t.startTimer()
				m.screen=1
			}

		case "esc":
			m.screen = 0
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
	case time.Time:
			if m.screen == 1 {
			m.duration = msg.Sub(m.startTime)
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return t
			})
		}
	}	
	

	// Always update the text input
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
} 
func (m model) View() string {
	if m.screen == 1 {
		return fmt.Sprintf(
			"⏱ Timer started at: %s\n⏳ Duration: %s \n\nPress 'esc' to return.",
			m.startTime.Format("15:04:05"),
			m.duration.Round(time.Second),
		)
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

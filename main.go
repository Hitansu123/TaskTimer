package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Println("Hello World")
	f,err:=tea.LogToFile("log.txt", "debug")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
  p := tea.NewProgram(newModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

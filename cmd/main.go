package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"perico6255/goffee/bubble/app"
)

func main() {
	app := app.New("prueba prueba")
	program := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

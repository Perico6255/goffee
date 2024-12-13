package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	prueba string
}

func New(prueba string) Model {
	return Model{
		prueba: prueba,
	}
}

func (m Model) View() string {
	return fmt.Sprintf("Hola app\n esto es una prueba: %s", m.prueba)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

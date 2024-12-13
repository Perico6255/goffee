package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	prueba string
	width  int
	height int
}

// Constructor
func New(prueba string) Model {
	return Model{
		prueba: prueba,
	}
}

// Utility Functions

// With the values of the with and heigh, set the atributes that depends on it
func (m *Model) setSizeAtributes(width int, height int) {
	m.width = width
	m.height = height
}

// BUBBLE TEA FUNTIONS

func (m Model) View() string {
	return fmt.Sprintf(
		"Hello app\n Mensage: %s \n width: %d\n height: %d",
		m.prueba,
		m.width,
		m.height,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.setSizeAtributes(msg.Width, msg.Height)

	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

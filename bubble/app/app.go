package app

import (
	tea "github.com/charmbracelet/bubbletea"

	mainbox "perico6255/goffee/bubble/components/mainBox"
)

type Model struct {
	prueba string
	width  int
	height int

	mainbox mainbox.Model
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
	mainbox := mainbox.New("/home/perico/Documents/", m.width-10, 12)
	m.mainbox = mainbox
}

// BUBBLE TEA FUNTIONS

func (m Model) View() string {
	// return fmt.Sprintf(
	// 	"Hello app\n Mensage: %s \n width: %d\n height: %d",
	// 	m.prueba,
	// 	m.width,
	// 	m.height,
	// )
	return m.mainbox.View()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.setSizeAtributes(msg.Width, msg.Height)

	}
	m.mainbox, cmd = m.mainbox.Update(msg)
	return m, cmd
}

func (m Model) Init() tea.Cmd {
	return nil
}

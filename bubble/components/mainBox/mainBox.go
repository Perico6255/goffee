package mainbox

import (
	"fmt"
	"io/fs"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cursor      int
	path        string
	files       []fs.FileInfo
	height      int
	width       int
	startIndex  int
	filesLenght int
	endIndex    int

	lineStyle           lipgloss.Style
	numberStyle         lipgloss.Style
	selectedNumberStyle lipgloss.Style
	nameStyle           lipgloss.Style
	permStyle           lipgloss.Style
}

func New(path string, width int, height int) Model {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// numberStyle := lipgloss.NewStyle().
	// 	Width(3).Align(lipgloss.Right)
	// nameStyle := lipgloss.NewStyle().
	// 	Width(30).Align(lipgloss.Left).Padding(0, 0, 0, 2)
	// permStyle := lipgloss.NewStyle().
	// 	Width(10).Align(lipgloss.Right)
	// selectedNumber := lipgloss.NewStyle().
	// 	Width(3).Align(lipgloss.Left).Background(lipgloss.Color(1))

	m := Model{
		cursor:    0,
		path:      path,
		files:     []fs.FileInfo{},
		lineStyle: lipgloss.Style{},
		height:    height,
		width:     width,
		endIndex:  height - 3,

		// selectedNumberStyle: selectedNumber,
		// numberStyle:         numberStyle,
		// nameStyle:           nameStyle,
		// permStyle:           permStyle,
	}
	for _, entry := range entries {
		file, err := entry.Info()
		if err != nil {
			panic(err)
		}
		m.files = append(m.files, file)
	}
	m.endIndex = len(m.files)
	if m.endIndex > m.height-2 {
		m.endIndex = m.height - 2
	}
	m.filesLenght = len(m.files)
	return m
}

func (m Model) String() string {
	var s string
	for _, file := range m.files {
		s += fmt.Sprintf(
			"Name: %s, IsDir: %v, Size: %v, More: %s \n",
			file.Name(),
			file.IsDir(),
			file.Size(),
			file.Mode().Perm(),
		)
	}
	return s
}

func (m *Model) up() {
	if m.cursor > 0 {
		m.cursor--
	}
	if m.startIndex+2 > m.cursor && m.startIndex > 0 {
		m.startIndex--
		m.endIndex--
	}
}

func (m *Model) down() {
	if m.cursor < len(m.files)-1 {
		m.cursor++
	}
	if m.endIndex-3 < m.cursor && m.endIndex < m.filesLenght {
		m.startIndex++
		m.endIndex++
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.up()

		case "down", "j":
			m.down()
		}
	}
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

// View implements tea.Model.
func (m Model) View() string {
	var s string
	var line int
	s += fmt.Sprintf("Index: %d, start: %d, end %d\n\n", m.cursor, m.startIndex, m.endIndex)
	for i := m.startIndex; i < m.endIndex; i++ {
		line = m.cursor - i
		if line < 0 {
			line = -line
		}
		s += fmt.Sprintf("%d %s\n", line, m.files[i].Name())
	}
	return s
}

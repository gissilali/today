package list

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gissilali/today/repositories"
)

func InitListTasksProgram() {
	program := tea.NewProgram(initialModel())
	if err := program.Start(); err != nil {
		fmt.Println("Error starting the program.")
	}
}

type model struct {
	tasks  []repositories.Task
	marked map[int]uint
	cursor int
	err    error
}

type errorMessage struct {
	error
}

func initialModel() tea.Model {
	var tasks []repositories.Task
	db := repositories.CurrentDB()
	db.Find(&tasks)
	return model{
		tasks:  tasks,
		marked: make(map[int]uint),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	// The header
	s := "\nTasks for today\n"

	// Iterate over our choices
	for i, choice := range m.tasks {
		renderedTask := choice.Task
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.marked[i]; ok || m.tasks[i].IsDone {
			checked = "X"
			renderedTask = selectedStyle.Render(choice.Task) // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, renderedTask)
	}

	// The footer
	s += "\nPress q to quit.\n"
	// Send the UI for rendering
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlS:
			return m, tea.Quit
		case tea.KeySpace:
			ok := m.tasks[m.cursor].IsDone
			taskId := m.tasks[m.cursor].ID
			db := repositories.CurrentDB()
			if ok {
				m.tasks[m.cursor].IsDone = !ok
				db.Exec("UPDATE tasks SET is_done = ? WHERE id = ?", !ok, taskId)
			} else {
				m.tasks[m.cursor].IsDone = true
				db.Exec("UPDATE tasks SET is_done = ? WHERE id = ?", true, taskId)
			}
		}

		switch msg.String() {
		case "q":
			////ids := make([]uint, 0, len(m.marked))
			//db := repositories.CurrentDB()
			//for _, value := range m.tasks {
			//	fmt.Println(value.IsDone)
			//	db.Exec("UPDATE tasks SET is_done = ? WHERE id = ?", value.IsDone, value.ID)
			//	//ids = append(ids, value)
			//}
			//
			//fmt.Println("Saved every thing")
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errorMessage:
		m.err = msg
		return m, nil
	}

	return m, cmd
}

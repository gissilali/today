package services

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"strconv"
)

func InitTasksProgram() {
	program := tea.NewProgram(initialModel())

	if err := program.Start(); err != nil {
		fmt.Println("Error starting the program.")
	}
}

type model struct {
	textInput  textinput.Model
	addedTasks []string
	err        error
}

func initialModel() tea.Model {
	taskInput := textinput.New()
	taskInput.Placeholder = ""
	taskInput.Focus()
	return model{
		textInput: taskInput,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
	tasksAddedCount := len(m.addedTasks)
	var helpText string
	if helpText = ""; tasksAddedCount > 0 {
		helpText = "(esc to quit) " + strconv.Itoa(tasksAddedCount) + " tasks added"
	} else {
		helpText = "(esc to quit)"
	}

	return fmt.Sprintf(
		"What do you got going on?\n\n%s\n\n%s",
		m.textInput.View(),
		helpText,
	) + "\n"
}

type errorMessage struct {
	error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.addedTasks = append(m.addedTasks, m.textInput.Value())
			m.textInput.SetValue("")
		case tea.KeyCtrlS:
			for _, task := range m.addedTasks {
				fmt.Println("✔", task)
			}

			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errorMessage:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

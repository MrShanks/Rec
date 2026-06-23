package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func NewApp() *Model {
	return &Model{}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m Model) View() string {
	return "Welcome to your recipe dashboard"
}

func main() {
	a := NewApp()
	p := tea.NewProgram(a)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Fuck something went seriously wrong: %v", err)
	}
}

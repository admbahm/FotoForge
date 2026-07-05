// Package tui contains reusable Bubble Tea models and Lip Gloss styles.
package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

// StatusModel is a minimal base for future interactive, cancellable operations.
type StatusModel struct {
	Title   string
	Spinner spinner.Model
}

func (m StatusModel) Init() tea.Cmd { return m.Spinner.Tick }
func (m StatusModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var command tea.Cmd
	m.Spinner, command = m.Spinner.Update(message)
	return m, command
}
func (m StatusModel) View() string { return titleStyle.Render(m.Title) + " " + m.Spinner.View() }

package main

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#00D4AA")).
		Padding(0, 3).
		Bold(true).
		MarginBottom(1)

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("#00D4AA")).
		Padding(1, 2).
		Width(45).
		MarginBottom(1)

	InputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	HintStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Italic(true)

	LabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00D4AA")).
		Bold(true)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF87")).
		Bold(true)

	QRBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00D4AA")).
		Padding(1).
		MarginTop(1).
		MarginBottom(1)

	StatusStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)
)
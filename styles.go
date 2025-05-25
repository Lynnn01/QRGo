package main

import "github.com/charmbracelet/lipgloss"

var (
	// Dark Theme เท่ๆ
	TitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#00FF87")).
		Padding(0, 3).
		Bold(true).
		MarginBottom(1)

	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#00FF87")).
		Padding(1, 2).
		Width(50).
		MarginBottom(1)

	InputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	HintStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555")).
		Italic(true)

	LabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF87")).
		Bold(true)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF87")).
		Bold(true)

	MenuStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#CCCCCC")).
		Bold(true)

	SelectedMenuStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#00FF87")).
		Bold(true).
		Padding(0, 1)

	StatusStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)

	LoadingStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		Bold(true)

	ProgressStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF87")).
		Bold(true)

	ErrorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6B6B")).
		Bold(true)
)
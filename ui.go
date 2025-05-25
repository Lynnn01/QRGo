package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderUI(input string) string {
	var s strings.Builder

	// Header แบบเท่
	s.WriteString(TitleStyle.Render(" ⚡ QR GENERATOR "))
	s.WriteString("\n\n")

	// Status line
	s.WriteString(StatusStyle.Render("⚙️  INPUT MODE"))
	s.WriteString("\n\n")

	// Input label
	s.WriteString(LabelStyle.Render("→ ENTER TEXT OR URL"))
	s.WriteString("\n\n")

	// Input box
	displayText := input
	if displayText == "" {
		displayText = HintStyle.Render("type here...")
	} else {
		displayText = InputStyle.Render(displayText + "▋")
	}

	s.WriteString(BoxStyle.Render(displayText))
	s.WriteString("\n")

	// Controls
	s.WriteString(HintStyle.Render("ENTER → generate • BACKSPACE → delete • CTRL+C → exit"))

	return s.String()
}

func RenderResult(message string) string {
	var s strings.Builder

	// Header
	s.WriteString(TitleStyle.Render(" ✨ QR GENERATED "))
	s.WriteString("\n\n")

	// Status
	s.WriteString(StatusStyle.Render("🎯 SUCCESS"))
	s.WriteString("\n\n")

	// Result message
	s.WriteString(SuccessStyle.Render(message))
	s.WriteString("\n\n")

	// Instructions
	s.WriteString(HintStyle.Render("ANY KEY → continue"))

	return s.String()
}

func RenderMenu(options []string, cursor int) string {
	var s strings.Builder

	// Header
	s.WriteString(TitleStyle.Render(" 🎯 CHOOSE ACTION "))
	s.WriteString("\n\n")

	// Status
	s.WriteString(StatusStyle.Render("📋 WHAT'S NEXT?"))
	s.WriteString("\n\n")

	// Menu options
	for i, option := range options {
		pointer := "  "
		if cursor == i {
			pointer = "▶ "
		}

		optionText := option
		if cursor == i {
			optionText = lipgloss.NewStyle().Foreground(lipgloss.Color("#00D4AA")).Bold(true).Render(option)
		} else {
			optionText = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Render(option)
		}

		s.WriteString(fmt.Sprintf("%s%s\n", pointer, optionText))
	}

	s.WriteString("\n")
	s.WriteString(HintStyle.Render("↑/↓ → select • ENTER → confirm • CTRL+C → exit"))

	return s.String()
}
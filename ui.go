package main

import (
	"fmt"
	"strings"
)

func RenderUI(input string) string {
	var s strings.Builder

	s.WriteString(TitleStyle.Render(" ⚡ QR GENERATOR "))
	s.WriteString("\n\n")

	s.WriteString(StatusStyle.Render("⚙️  INPUT MODE"))
	s.WriteString("\n\n")

	s.WriteString(LabelStyle.Render("▶ ENTER TEXT OR URL"))
	s.WriteString("\n\n")

	displayText := input
	if displayText == "" {
		displayText = HintStyle.Render("type here...")
	} else {
		displayText = InputStyle.Render(displayText + "▋")
	}

	s.WriteString(BoxStyle.Render(displayText))
	s.WriteString("\n")

	s.WriteString(HintStyle.Render("ENTER → generate • BACKSPACE → delete • CTRL+C → exit"))

	return s.String()
}

func RenderLoading(step int) string {
	var s strings.Builder

	s.WriteString(TitleStyle.Render(" ⚡ QR GENERATOR "))
	s.WriteString("\n\n")

	s.WriteString(StatusStyle.Render("🔥 GENERATING QR CODE"))
	s.WriteString("\n\n")

	// Animated spinner
	spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spinners[step%len(spinners)]
	
	s.WriteString(LoadingStyle.Render(fmt.Sprintf("%s PROCESSING...", spinner)))
	s.WriteString("\n\n")

	// Progress bar
	progress := (step % 30) * 100 / 30
	progressBar := ""
	for i := 0; i < 30; i++ {
		if i < (step%30) {
			progressBar += "█"
		} else {
			progressBar += "░"
		}
	}

	s.WriteString(ProgressStyle.Render(fmt.Sprintf("[%s] %d%%", progressBar, progress)))
	s.WriteString("\n\n")

	s.WriteString(HintStyle.Render("⏳ Please wait... Creating your QR code"))

	return s.String()
}

func RenderResultWithMenu(message string, options []string, cursor int) string {
	var s strings.Builder

	s.WriteString(TitleStyle.Render(" ✨ COMPLETED "))
	s.WriteString("\n\n")

	s.WriteString(StatusStyle.Render("🎯 SUCCESS"))
	s.WriteString("\n\n")

	// Result message
	s.WriteString(SuccessStyle.Render(message))
	s.WriteString("\n\n")

	// Separator
	s.WriteString(strings.Repeat("▔", 50))
	s.WriteString("\n\n")

	// Menu
	s.WriteString(LabelStyle.Render("▶ WHAT'S NEXT?"))
	s.WriteString("\n\n")

	for i, option := range options {
		pointer := "  "
		if cursor == i {
			pointer = "▶ "
		}

		optionText := option
		if cursor == i {
			optionText = SelectedMenuStyle.Render(" " + option + " ")
		} else {
			optionText = MenuStyle.Render(option)
		}

		s.WriteString(fmt.Sprintf("%s%s\n", pointer, optionText))
	}

	s.WriteString("\n")
	s.WriteString(HintStyle.Render("↑/↓ → select • ENTER → confirm • CTRL+C → exit"))

	return s.String()
}
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Input         string
	ShowResult    bool
	ResultMessage string
	ShowMenu      bool
	MenuCursor    int
	MenuOptions   []string
}

func NewModel() Model {
	return Model{
		Input:       "",
		ShowResult:  false,
		ShowMenu:    false,
		MenuCursor:  0,
		MenuOptions: []string{"🔄 เริ่มใหม่", "❌ ออก"},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// หน้าเมนูหลังเจน
		if m.ShowMenu {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.MenuCursor > 0 {
					m.MenuCursor--
				}
			case "down", "j":
				if m.MenuCursor < len(m.MenuOptions)-1 {
					m.MenuCursor++
				}
			case "enter", " ":
				switch m.MenuCursor {
				case 0: // เริ่มใหม่
					clearScreen()
					m.Input = ""
					m.ShowResult = false
					m.ShowMenu = false
					m.MenuCursor = 0
					m.ResultMessage = ""
					return m, nil
				case 1: // ออก
					return m, tea.Quit
				}
			}
			return m, nil
		}

		// หน้าผลลัพธ์
		if m.ShowResult {
			m.ShowResult = false
			m.ShowMenu = true
			return m, nil
		}

		// หน้า input
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if strings.TrimSpace(m.Input) != "" {
				err := CreateAndSaveQR(m.Input)
				if err != nil {
					m.ResultMessage = fmt.Sprintf("❌ %v", err)
				} else {
					fileName := GenerateFileName(m.Input)
					m.ResultMessage = fmt.Sprintf("🎉 สร้าง QR Code สำเร็จ!\n📁 ไฟล์: qrcode/%s.png\n📋 รูปถูก copy เข้า clipboard แล้ว", fileName)
				}
				m.ShowResult = true
				return m, nil
			}

		case "backspace":
			if len(m.Input) > 0 {
				m.Input = m.Input[:len(m.Input)-1]
			}

		default:
			if len(msg.String()) == 1 {
				m.Input += msg.String()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.ShowMenu {
		return RenderMenu(m.MenuOptions, m.MenuCursor)
	}
	if m.ShowResult {
		return RenderResult(m.ResultMessage)
	}
	return RenderUI(m.Input)
}

// Clear screen ตาม OS
func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
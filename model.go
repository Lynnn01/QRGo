package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Input         string
	Loading       bool
	LoadingStep   int
	ShowResult    bool
	ResultMessage string
	MenuCursor    int
	MenuOptions   []string
}

// Messages สำหรับ loading
type loadingMsg int
type doneMsg string

func NewModel() Model {
	return Model{
		Input:       "",
		Loading:     false,
		LoadingStep: 0,
		ShowResult:  false,
		MenuCursor:  0,
		MenuOptions: []string{"🔄 เริ่มใหม่", "❌ ออก"},
	}
}

// Loading command
func loadingTick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return loadingMsg(1)
	})
}

// Generate QR command
func generateQR(input string) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(3 * time.Second) // จำลอง processing 3 วิ
		err := CreateAndSaveQR(input)
		if err != nil {
			return doneMsg(fmt.Sprintf("❌ %v", err))
		}
		fileName := GenerateFileName(input)
		return doneMsg(fmt.Sprintf("🎉 QR CODE GENERATED\n📁 FILE: qrcode/%s.png\n📋 IMAGE COPIED TO CLIPBOARD", fileName))
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	// Loading animation
	case loadingMsg:
		if m.Loading {
			m.LoadingStep = (m.LoadingStep + 1) % 30 // 30 steps animation
			return m, loadingTick()
		}
		return m, nil
	
	// Generation complete
	case doneMsg:
		m.Loading = false
		m.LoadingStep = 0
		m.ResultMessage = string(msg)
		m.ShowResult = true
		return m, nil

	case tea.KeyMsg:
		// หน้าผลลัพธ์พร้อมเมนู
		if m.ShowResult {
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
					m = NewModel() // Reset ทุกอย่าง
					return m, nil
				case 1: // ออก
					return m, tea.Quit
				}
			}
			return m, nil
		}

		// หน้า input (ถ้าไม่ loading)
		if !m.Loading {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit

			case "enter":
				if strings.TrimSpace(m.Input) != "" {
					m.Loading = true
					m.LoadingStep = 0
					return m, tea.Batch(loadingTick(), generateQR(m.Input))
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
	}
	return m, nil
}

func (m Model) View() string {
	if m.Loading {
		return RenderLoading(m.LoadingStep)
	}
	if m.ShowResult {
		return RenderResultWithMenu(m.ResultMessage, m.MenuOptions, m.MenuCursor)
	}
	return RenderUI(m.Input)
}

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
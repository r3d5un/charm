package progressbar

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var ProgressBarCmd = &cobra.Command{
	Use:     "progressbar",
	Aliases: []string{"prog", "pb"},
	Short:   "Displays a progressbar, then exits",
	Run: func(cmd *cobra.Command, args []string) {
		m := newModel()
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("an error occurred", err)
			os.Exit(1)
		}
	},
}

type workDoneMsg struct {
	count *int64
}

type doneMsg struct{}

type model struct {
	progress progress.Model
	done     bool
	total    int64
	count    int64
}

func (m model) Init() tea.Cmd {
	return tea.Batch(doWork(&m), tickCmd())
}

// Update function for Bubble Tea
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.done {
			return m, nil
		}

		if m.count >= m.total {
			m.done = true
			return m, func() tea.Msg { return doneMsg{} }
		}

		val := float64(m.count) / float64(m.total)
		fmt.Println(val)

		cmd := m.progress.SetPercent(float64(m.count) / float64(m.total))
		return m, tea.Batch(tickCmd(), cmd)

	case workDoneMsg:
		m.done = true
		return m, nil

	case doneMsg:
		m.done = true
		return m, tea.Quit

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)

	if m.done {
		return "\n" +
			pad + "Done" + "\n"
	}

	return "\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func newModel() model {
	return model{
		progress: progress.New(progress.WithDefaultGradient()),
		total:    1_000,
		count:    0,
	}
}

func doWork(m *model) tea.Cmd {
	return func() tea.Msg {
		workers := 100
		var wg sync.WaitGroup
		wg.Add(workers)

		for i := 0; i < workers; i++ {
			go func() {
				defer wg.Done()
				for {
					current := atomic.AddInt64(&m.count, 1)
					// fmt.Println(current)
					if current > m.total {
						break
					}
					// Simulate work
					time.Sleep(time.Millisecond * 500)
				}
			}()
		}

		wg.Wait()
		return workDoneMsg{count: &m.count}
	}
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

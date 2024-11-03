package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
	"math"
	"n0rdy.foo/calcli/calc"
	"n0rdy.foo/calcli/calc/parser"
	"n0rdy.foo/calcli/utils"
)

type Model struct {
	ti             textinput.Model
	result         string
	history        list.Model
	historyPointer int
	help           help.Model
	c              calc.CalcProcessor
	mode           int
	err            error
}

const (
	calculatorMode = iota
	historyMode
	cmdMode
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

func InitialModel(c calc.CalcProcessor) *Model {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 80
	ti.ShowSuggestions = true
	ti.SetSuggestions(utils.KnownCalculatorTokens)

	history := list.New([]list.Item{}, list.NewDefaultDelegate(), 20, 14)
	history.Title = "History"
	history.SetShowHelp(false)
	history.DisableQuitKeybindings()

	h := help.New()
	h.ShowAll = true

	return &Model{
		ti:             ti,
		result:         "",
		history:        history,
		historyPointer: -1,
		help:           h,
		mode:           calculatorMode,
		c:              c,
		err:            nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// switches to the command mode
		if msg.String() == ":" {
			m.mode = cmdMode
			m.ti.SetValue(":")
			m.ti.SetSuggestions(cmdTokens)
			return m, nil
		}

		switch msg.Type {
		// quits the program
		case tea.KeyCtrlC, tea.KeyEscape:
			return m, tea.Quit
		case tea.KeyUp:
			// handles history navigation backwards for calculator mode
			if m.mode == calculatorMode {
				return m.nextHistoryRecord()
			}
		case tea.KeyDown:
			// handles history navigation forwards for calculator mode
			if m.mode == calculatorMode {
				return m.prevHistoryRecord()
			}
		// handles finalized text input
		case tea.KeyEnter:
			return m.processInput()
		}
	}

	var cmd tea.Cmd
	switch m.mode {
	case calculatorMode:
		m.ti, cmd = m.ti.Update(msg)
		m.historyPointer = len(m.history.Items()) - 1
	case cmdMode:
		m.ti, cmd = m.ti.Update(msg)
	case historyMode:
		m.history, cmd = m.history.Update(msg)
	}
	return m, cmd
}

func (m Model) View() string {
	switch m.mode {
	case calculatorMode:
		return m.calculatorModeView()
	case historyMode:
		return m.historyModeView()
	case cmdMode:
		return m.cmdModeView()
	default:
		return m.calculatorModeView()
	}
}

func (m Model) nextHistoryRecord() (Model, tea.Cmd) {
	if len(m.history.Items()) == 0 {
		return m, nil
	}
	if m.historyPointer < 0 {
		return m, nil
	}

	m.ti.SetValue(m.history.Items()[m.historyPointer].(historyRecord).input)
	m.historyPointer--
	return m, nil
}

func (m Model) prevHistoryRecord() (Model, tea.Cmd) {
	if len(m.history.Items()) == 0 {
		return m, nil
	}
	if m.historyPointer >= len(m.history.Items())-1 {
		return m, nil
	}

	m.historyPointer++
	m.ti.SetValue(m.history.Items()[m.historyPointer].(historyRecord).input)
	return m, nil
}

func (m Model) processInput() (Model, tea.Cmd) {
	switch m.mode {
	case calculatorMode:
		return m.processCalculatorModeInput(m.ti.Value())
	case cmdMode:
		return m.processCmdModeInput(m.ti.Value())
	case historyMode:
		hr := m.history.SelectedItem().(historyRecord)
		m.err = nil
		m.ti.SetValue(hr.input)
		m.result = ""
		m.mode = calculatorMode
		return m, nil
	default:
		return m, nil
	}
}

func (m Model) processCalculatorModeInput(input string) (Model, tea.Cmd) {
	res, err := m.c.Process(input)
	if err != nil {
		m.err = err
		return m, nil
	}

	m.err = nil
	// to show the previous calculation input
	m.ti.Placeholder = m.ti.Value()
	m.ti.SetValue("")
	// resetting suggestions in case if the new variable was added
	m.ti.SetSuggestions(utils.KnownCalculatorTokens)
	m.result = m.processCalcResult(*res)
	// adding the new calculation to the history
	m.history.InsertItem(0, historyRecord{input: input, result: m.result})
	m.historyPointer = len(m.history.Items()) - 1

	return m, nil
}

func (m Model) processCalcResult(res parser.CalcResult) string {
	if res.Type == parser.OutputType {
		return res.Output
	}

	val := res.Value
	if math.IsNaN(val) {
		return "NaN"
	}
	if math.IsInf(val, -1) {
		return "-∞"
	}
	if math.IsInf(val, 1) {
		return "+∞"
	}

	utils.SetPreviousResult(val)
	return humanize.Commaf(val)
}

func (m Model) processCmdModeInput(input string) (Model, tea.Cmd) {
	switch input {
	case historyCmdToken:
		m.mode = historyMode
		m.err = nil
		m.ti.SetValue("")
		return m, nil
	case calculatorCmdToken:
		m.mode = calculatorMode
		m.err = nil
		m.ti.SetValue("")
		return m, nil
	default:
		m.err = fmt.Errorf("unknown command: %s", input)
		return m, nil
	}
}

func (m Model) calculatorModeView() string {
	var view string
	view += "Type your expression and press Enter\n\n"
	view += m.ti.View() + "\n\n"
	if m.err != nil {
		view += m.err.Error() + "\n\n"
	} else if m.result != "" {
		view += m.result + "\n\n"
	}
	view += m.help.View(calculatorModeKeys)
	return view
}

func (m Model) historyModeView() string {
	var view string
	view += docStyle.Render(m.history.View()) + "\n\n"
	view += m.help.View(historyModeKeys)
	return view
}

func (m Model) cmdModeView() string {
	var view string
	view += "Type your command and press Enter\n\n"
	view += m.ti.View() + "\n\n"
	if m.err != nil {
		view += m.err.Error() + "\n\n"
	}
	view += m.help.View(cmdModeKeys)
	return view
}

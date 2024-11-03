package tui

import "github.com/charmbracelet/bubbles/key"

type keyMapCalculatorMode struct {
	Enter key.Binding
	Up    key.Binding
	Down  key.Binding
	Colon key.Binding
	Quit  key.Binding
}

type keyMapCmdMode struct {
	Enter      key.Binding
	Calculator key.Binding
	History    key.Binding
	Quit       key.Binding
}

type keyMapHistoryMode struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Filter key.Binding
	Quit   key.Binding
}

var (
	quitKeyBinding = key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c/esc", "Quit"),
	)

	calculatorModeKeys = keyMapCalculatorMode{
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Calculate the expression"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "Previous history entry"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "Next history entry"),
		),
		Colon: key.NewBinding(
			key.WithKeys(":"),
			key.WithHelp(":", "Enter command mode"),
		),
		Quit: quitKeyBinding,
	}

	cmdModeKeys = keyMapCmdMode{
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Execute the command"),
		),
		Calculator: key.NewBinding(
			key.WithKeys(calculatorCmdToken),
			key.WithHelp(calculatorCmdToken, "Switch to calculator mode"),
		),
		History: key.NewBinding(
			key.WithKeys(historyCmdToken),
			key.WithHelp(historyCmdToken, "Switch to history mode"),
		),
		Quit: quitKeyBinding,
	}

	historyModeKeys = keyMapHistoryMode{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "Navigate up the list"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "Navigate down the list"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Select and switch to calculator mode"),
		),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "Filter"),
		),
		Quit: quitKeyBinding,
	}
)

// ShortHelp returns keybindings to be shown in the mini help view
func (k keyMapCalculatorMode) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k keyMapCalculatorMode) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Colon, k.Up, k.Down}, // first column
		{k.Quit},                         // second column
	}
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k keyMapCmdMode) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k keyMapCmdMode) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter, k.Calculator, k.History}, // first column
		{k.Quit},                           // second column
	}
}

// ShortHelp returns keybindings to be shown in the mini help view
func (k keyMapHistoryMode) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

// FullHelp returns keybindings for the expanded help view
func (k keyMapHistoryMode) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Filter}, // first column
		{k.Enter, k.Quit},        // second column
	}
}

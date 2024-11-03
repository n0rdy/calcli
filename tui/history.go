package tui

type historyRecord struct {
	input  string
	result string
}

func (hr historyRecord) Title() string       { return hr.input }
func (hr historyRecord) Description() string { return hr.result }
func (hr historyRecord) FilterValue() string { return hr.input }

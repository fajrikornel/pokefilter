package filter

import "errors"

// Filter interface
// BuildQuery should return a query string that returns a single column: `pokemon_id`
// Chain should accept another Filter and incorporate that other Filter's query inside its own.
type Filter interface {
	BuildQuery() (string, error)
	Chain(other Filter) (Filter, error)
	Display() string
}

type TerminalFilter struct {
	query string
}

func NewTerminalFilter() *TerminalFilter {
	return &TerminalFilter{
		query: `SELECT id AS pokemon_id FROM pokemon`,
	}
}

func (f *TerminalFilter) BuildQuery() (string, error) {
	return f.query, nil
}

func (f *TerminalFilter) Chain(_ Filter) (Filter, error) {
	return nil, errors.New("terminal filter should be the last filter, it cannot be chained with other filters")
}

func (f *TerminalFilter) Display() string {
	return "Terminal Filter. Not supposed to Display()!"
}

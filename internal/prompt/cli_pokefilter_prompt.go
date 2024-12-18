package prompt

import (
	"fmt"
	"github.com/buger/goterm"
	pokefilter "github.com/fajrikornel/pokefilter/internal"
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/manifoldco/promptui"
)

type CLIPrompter struct {
	pf *pokefilter.PokeFilter

	statFilterPrompter     *StatFilterPrompter
	typeFilterPrompter     *TypeFilterPrompter
	abilityFilterPrompter  *AbilityFilterPrompter
	moveFilterPrompter     *MoveFilterPrompter
	moveTypeFilterPrompter *MoveTypeFilterPrompter

	statFilter   *filter.StatFilter
	typeFilter   *filter.TypeFilter
	otherFilters []filter.Filter
}

func NewCLIPokefilterPrompter(
	pf *pokefilter.PokeFilter,
	statFilterPrompter *StatFilterPrompter,
	typeFilterPrompter *TypeFilterPrompter,
	abilityFilterPrompter *AbilityFilterPrompter,
	moveFilterPrompter *MoveFilterPrompter,
	moveTypeFilterPrompter *MoveTypeFilterPrompter,
) (*CLIPrompter, error) {
	return &CLIPrompter{
		pf:                     pf,
		statFilterPrompter:     statFilterPrompter,
		typeFilterPrompter:     typeFilterPrompter,
		abilityFilterPrompter:  abilityFilterPrompter,
		moveFilterPrompter:     moveFilterPrompter,
		moveTypeFilterPrompter: moveTypeFilterPrompter,
		statFilter:             nil,
		typeFilter:             nil,
		otherFilters:           []filter.Filter{},
	}, nil
}

func (p *CLIPrompter) Run() ([]model.Pokemon, error) {

	doneSelecting := false
	options := []string{
		"EXECUTE FILTER",
		"Add/modify pokemon stats filter",
		"Add/modify pokemon type filter",
		"Add/clear pokemon abilities/move filter",
	}
	for !doneSelecting {
		p.clearScreen()
		p.displayCurrentFilterInformation()

		prompt := promptui.Select{
			Label: "Select action",
			Items: options,
		}

		i, _, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		switch i {
		case 0:
			doneSelecting = true
		case 1:
			statFilter, err := p.statFilterPrompter.Run()
			if err != nil {
				return nil, err
			}

			p.statFilter = statFilter
		case 2:
			typeFilter, err := p.typeFilterPrompter.Run()
			if err != nil {
				return nil, err
			}

			p.typeFilter = typeFilter
		case 3:
			otherOptions := []string{
				"Clear all abilities/moves filters",
				"Add ability filter",
				"Add move filter by move name",
				"Add move filter by move type and/or damage class",
			}

			otherPrompt := promptui.Select{
				Label: "Select action",
				Items: otherOptions,
			}

			i, _, err := otherPrompt.Run()
			if err != nil {
				return nil, err
			}

			switch i {
			case 0:
				p.otherFilters = []filter.Filter{}
			case 1:
				abilityFilter, err := p.abilityFilterPrompter.Run()
				if err != nil {
					return nil, err
				}

				p.otherFilters = append(p.otherFilters, abilityFilter)
			case 2:
				moveFilter, err := p.moveFilterPrompter.Run()
				if err != nil {
					return nil, err
				}

				p.otherFilters = append(p.otherFilters, moveFilter)
			case 3:
				moveTypeFilter, err := p.moveTypeFilterPrompter.Run()
				if err != nil {
					return nil, err
				}

				p.otherFilters = append(p.otherFilters, moveTypeFilter)
			}
		}
	}

	if p.statFilter != nil {
		p.pf.AddFilter(p.statFilter)
	}

	if p.typeFilter != nil {
		p.pf.AddFilter(p.typeFilter)
	}

	for _, f := range p.otherFilters {
		p.pf.AddFilter(f)
	}

	return p.pf.ExecuteFilter()
}

func (p *CLIPrompter) displayCurrentFilterInformation() {
	fmt.Println("Current filters:")
	if p.statFilter == nil && p.typeFilter == nil && (p.otherFilters == nil || len(p.otherFilters) == 0) {
		fmt.Println("No filter yet.")
		return
	}

	if p.statFilter != nil {
		fmt.Println(p.statFilter.Display())
	}

	if p.typeFilter != nil {
		fmt.Println(p.typeFilter.Display())
	}

	for _, f := range p.otherFilters {
		fmt.Println(f.Display())
	}
}

func (p *CLIPrompter) clearScreen() {
	goterm.Clear()
	goterm.Flush()
}

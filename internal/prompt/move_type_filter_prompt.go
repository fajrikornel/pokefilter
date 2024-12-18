package prompt

import (
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/manifoldco/promptui"
	"strings"
)

type MoveTypeFilterPrompter struct {
	typeStore        *model.TypeStore
	damageClassStore *model.DamageClassStore
}

func NewMoveTypeFilterPrompter(typeStore *model.TypeStore, damageClassStore *model.DamageClassStore) *MoveTypeFilterPrompter {
	return &MoveTypeFilterPrompter{
		typeStore:        typeStore,
		damageClassStore: damageClassStore,
	}
}

func (p *MoveTypeFilterPrompter) Run() (*filter.MoveTypeFilter, error) {
	moveTypeFilterOptions := []string{
		"Specify a move type",
		"Specify a damage class (physical/special/status)",
		"Specify a move type and damage class",
	}

	prompt := promptui.Select{
		Label: "What do you want to filter by?",
		Items: moveTypeFilterOptions,
	}

	opt, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	switch opt {
	case 0:
		return p.specifyMoveTypePrompt()
	case 1:
		return p.specifyDamageClassPrompt()
	case 2:
		return p.specifyMoveTypeAndDamageClassPrompt()
	default:
		panic("invalid option")
	}
}

func (p *MoveTypeFilterPrompter) specifyMoveTypePrompt() (*filter.MoveTypeFilter, error) {
	types, err := p.typeStore.GetAllTypes()
	if err != nil {
		return nil, err
	}

	typeNames := make([]string, 0)
	for _, t := range types {
		typeNames = append(typeNames, t.Name)
	}

	prompt := promptui.Select{
		Label: "Select move type",
		Items: typeNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedType, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewMoveTypeFilter(selectedType, "")
}

func (p *MoveTypeFilterPrompter) specifyDamageClassPrompt() (*filter.MoveTypeFilter, error) {
	damageClasses, err := p.damageClassStore.GetAllDamageClasses()
	if err != nil {
		return nil, err
	}

	damageClassNames := make([]string, 0)
	for _, dc := range damageClasses {
		damageClassNames = append(damageClassNames, dc.Name)
	}

	prompt := promptui.Select{
		Label: "Select damage class",
		Items: damageClassNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(damageClassNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedDamageClass, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewMoveTypeFilter("", selectedDamageClass)
}

func (p *MoveTypeFilterPrompter) specifyMoveTypeAndDamageClassPrompt() (*filter.MoveTypeFilter, error) {
	types, err := p.typeStore.GetAllTypes()
	if err != nil {
		return nil, err
	}

	typeNames := make([]string, 0)
	for _, t := range types {
		typeNames = append(typeNames, t.Name)
	}

	prompt := promptui.Select{
		Label: "Select move type",
		Items: typeNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedType, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	damageClasses, err := p.damageClassStore.GetAllDamageClasses()
	if err != nil {
		return nil, err
	}

	damageClassNames := make([]string, 0)
	for _, dc := range damageClasses {
		damageClassNames = append(damageClassNames, dc.Name)
	}

	prompt = promptui.Select{
		Label: "Select damage class",
		Items: damageClassNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(damageClassNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedDamageClass, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewMoveTypeFilter(selectedType, selectedDamageClass)
}

package prompt

import (
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/manifoldco/promptui"
	"strings"
)

type TypeFilterPrompter struct {
	typeStore *model.TypeStore
}

func NewTypeFilterPrompter(typeStore *model.TypeStore) *TypeFilterPrompter {
	return &TypeFilterPrompter{
		typeStore: typeStore,
	}
}

func (p *TypeFilterPrompter) Run() (*filter.TypeFilter, error) {
	types, err := p.typeStore.GetAllTypes()
	if err != nil {
		return nil, err
	}

	typeNames := make([]string, 0)
	for _, t := range types {
		typeNames = append(typeNames, t.Name)
	}

	typeFilterOptions := []string{
		"Specify a type that the pokemon has",
		"Specify both types that the pokemon has",
		"Specify a type that the pokemon has, and specify types that the pokemon must not have",
		"Only specify types that the pokemon must not have",
	}

	prompt := promptui.Select{
		Label: "What do you want to filter by?",
		Items: typeFilterOptions,
	}

	opt, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	switch opt {
	case 0:
		return p.specifyType1(typeNames)
	case 1:
		return p.specifyType1AndType2(typeNames)
	case 2:
		return p.specifyType1AndExcludedTypes(typeNames)
	case 3:
		return p.specifyExcludedTypes(typeNames)
	default:
		panic("invalid option")
	}
}

func (p *TypeFilterPrompter) specifyType1(typeNames []string) (*filter.TypeFilter, error) {
	prompt := promptui.Select{
		Label: "Select type",
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

	return filter.NewTypeFilter(selectedType, "", []string{})
}

func (p *TypeFilterPrompter) specifyType1AndType2(typeNames []string) (*filter.TypeFilter, error) {
	prompt := promptui.Select{
		Label: "Select one of the pokemon's type",
		Items: typeNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	i, selectedType1, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	typeNames = append(typeNames[:i], typeNames[i+1:]...)
	prompt = promptui.Select{
		Label: "Select one more of the pokemon's type",
		Items: typeNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedType2, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewTypeFilter(selectedType1, selectedType2, []string{})
}

func (p *TypeFilterPrompter) specifyType1AndExcludedTypes(typeNames []string) (*filter.TypeFilter, error) {
	prompt := promptui.Select{
		Label: "Select one of the pokemon's type",
		Items: typeNames,
		Searcher: func(input string, index int) bool {
			return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	i, selectedType, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	typeNames = append(typeNames[:i], typeNames[i+1:]...)
	excludedTypeSelections := append([]string{"DONE SELECTING"}, typeNames...)

	doneSelecting := false
	excludedTypes := []string{}
	for !doneSelecting {
		prompt := promptui.Select{
			Label: "Select a type to be excluded",
			Items: excludedTypeSelections,
			Searcher: func(input string, index int) bool {
				return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
			},
			StartInSearchMode: true,
		}

		i, selectedExcludedType, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		if i == 0 {
			doneSelecting = true
		} else {
			excludedTypes = append(excludedTypes, selectedExcludedType)
			excludedTypeSelections = append(excludedTypeSelections[:i], excludedTypeSelections[i+1:]...)
		}
	}

	return filter.NewTypeFilter(selectedType, "", excludedTypes)
}

func (p *TypeFilterPrompter) specifyExcludedTypes(typeNames []string) (*filter.TypeFilter, error) {
	excludedTypeSelections := append([]string{"DONE SELECTING"}, typeNames...)

	doneSelecting := false
	excludedTypes := []string{}
	for !doneSelecting {
		prompt := promptui.Select{
			Label: "Select a type to be excluded",
			Items: excludedTypeSelections,
			Searcher: func(input string, index int) bool {
				return input == "" || strings.Contains(strings.ToLower(typeNames[index]), strings.ToLower(input))
			},
			StartInSearchMode: true,
		}

		i, selectedExcludedType, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		if i == 0 {
			doneSelecting = true
		} else {
			excludedTypes = append(excludedTypes, selectedExcludedType)
			excludedTypeSelections = append(excludedTypeSelections[:i], excludedTypeSelections[i+1:]...)
		}
	}

	return filter.NewTypeFilter("", "", excludedTypes)
}

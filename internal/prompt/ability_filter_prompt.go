package prompt

import (
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/manifoldco/promptui"
	"regexp"
	"strings"
)

type AbilityFilterPrompter struct {
	abilityStore *model.AbilityStore
}

func NewAbilityFilterPrompter(abilityStore *model.AbilityStore) *AbilityFilterPrompter {
	return &AbilityFilterPrompter{
		abilityStore: abilityStore,
	}
}

func (p *AbilityFilterPrompter) Run() (*filter.AbilityFilter, error) {
	abilities, err := p.abilityStore.GetAllAbilities()
	if err != nil {
		return nil, err
	}

	abilityNames := make([]string, 0)
	for _, dc := range abilities {
		abilityNames = append(abilityNames, dc.Name)
	}

	spaceRegex := regexp.MustCompile(`\s+`)
	dashRegex := regexp.MustCompile(`-+`)

	prompt := promptui.Select{
		Label: "Select an ability",
		Items: abilityNames,
		Searcher: func(input string, index int) bool {
			input = spaceRegex.ReplaceAllString(input, "")
			input = dashRegex.ReplaceAllString(input, "")

			ability := abilityNames[index]
			ability = spaceRegex.ReplaceAllString(ability, "")
			ability = dashRegex.ReplaceAllString(ability, "")

			return input == "" || strings.Contains(strings.ToLower(ability), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedAbility, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewAbilityFilter(selectedAbility)
}

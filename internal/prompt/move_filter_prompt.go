package prompt

import (
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/manifoldco/promptui"
	"regexp"
	"strings"
)

type MoveFilterPrompter struct {
	moveStore *model.MoveStore
}

func NewMoveFilterPrompter(moveStore *model.MoveStore) *MoveFilterPrompter {
	return &MoveFilterPrompter{
		moveStore: moveStore,
	}
}

func (p *MoveFilterPrompter) Run() (*filter.MoveFilter, error) {
	moves, err := p.moveStore.GetAllMoves()
	if err != nil {
		return nil, err
	}

	moveNames := make([]string, 0)
	for _, dc := range moves {
		moveNames = append(moveNames, dc.Name)
	}

	spaceRegex := regexp.MustCompile(`\s+`)
	dashRegex := regexp.MustCompile(`-+`)

	prompt := promptui.Select{
		Label: "Select a move",
		Items: moveNames,
		Searcher: func(input string, index int) bool {
			input = spaceRegex.ReplaceAllString(input, "")
			input = dashRegex.ReplaceAllString(input, "")

			move := moveNames[index]
			move = spaceRegex.ReplaceAllString(move, "")
			move = dashRegex.ReplaceAllString(move, "")

			return input == "" || strings.Contains(strings.ToLower(move), strings.ToLower(input))
		},
		StartInSearchMode: true,
	}

	_, selectedMove, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return filter.NewMoveFilter(selectedMove)
}

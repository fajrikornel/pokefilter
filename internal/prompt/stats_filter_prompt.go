package prompt

import (
	"errors"
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/manifoldco/promptui"
	"math"
	"strconv"
)

type StatFilterPrompter struct {
}

func NewStatFilterPrompter() *StatFilterPrompter {
	return &StatFilterPrompter{}
}

func (p *StatFilterPrompter) Run() (*filter.StatFilter, error) {
	statOptions := []struct {
		Stat    string
		Minimum int
		Maximum int
	}{
		{Stat: "DONE MODIFYING"},
		{Stat: "HP", Minimum: -1, Maximum: -1},
		{Stat: "Attack", Minimum: -1, Maximum: -1},
		{Stat: "Defense", Minimum: -1, Maximum: -1},
		{Stat: "Special Attack", Minimum: -1, Maximum: -1},
		{Stat: "Special Defense", Minimum: -1, Maximum: -1},
		{Stat: "Speed", Minimum: -1, Maximum: -1},
	}

	doneModifying := false
	for !doneModifying {
		prompt := promptui.Select{
			Label: "Which one do you want to modify?",
			Items: statOptions,
			Size:  len(statOptions),
			Templates: &promptui.SelectTemplates{
				Active:   "{{ if eq .Stat \"DONE MODIFYING\"}}\u25B8 DONE MODIFYING{{ else }}▸ {{ .Stat }} / {{ if eq .Minimum -1 }}0{{ else }}{{ .Minimum }}{{ end }} - {{ if eq .Maximum -1 }}\u221E{{ else }}{{ .Maximum }}{{ end }}{{ end }}",
				Inactive: "{{ if eq .Stat \"DONE MODIFYING\"}}DONE MODIFYING{{ else }}{{ .Stat }} / {{ if eq .Minimum -1 }}0{{ else }}{{ .Minimum }}{{ end }} - {{ if eq .Maximum -1 }}\u221E{{ else }}{{ .Maximum }}{{ end }}{{ end }}",
				Selected: "Modifying: {{ .Stat }}",
			},
		}

		i, _, err := prompt.Run()
		if err != nil {
			return nil, err
		}

		if i == 0 {
			doneModifying = true
		} else {

			values := []struct {
				Label string
				Value int
			}{
				{Label: "DONE MODIFYING", Value: 0},
				{Label: "Minimum", Value: statOptions[i].Minimum},
				{Label: "Maximum", Value: statOptions[i].Maximum},
			}

			doneSetting := false
			for !doneSetting {
				prompt = promptui.Select{
					Label: "Which one do you want to modify?",
					Items: values,
					Templates: &promptui.SelectTemplates{
						Active:   "{{ if eq .Label \"DONE MODIFYING\"}}\u25B8 DONE MODIFYING{{ else }}▸ {{ .Label }} = {{ if eq .Value -1 }}{{ if eq .Label \"Minimum\" }}0{{ else }}\u221E{{ end }}{{ else }}{{ .Value }}{{ end }}{{ end }}",
						Inactive: "{{ if eq .Label \"DONE MODIFYING\"}}DONE MODIFYING{{ else }}{{ .Label }} = {{ if eq .Value -1 }}{{ if eq .Label \"Minimum\" }}0{{ else }}∞{{ end }}{{ else }}{{ .Value }}{{ end }}{{ end }}",
						Selected: " ",
					},
				}

				j, _, err := prompt.Run()
				if err != nil {
					return nil, err
				}

				switch j {
				case 0:
					doneSetting = true
				case 1:
					valuePrompt := promptui.Prompt{
						Label: "Set the value",
						Validate: func(s string) error {
							if _, err := strconv.Atoi(s); err != nil {
								return err
							} else {
								return nil
							}
						},
					}

					valueString, err := valuePrompt.Run()
					if err != nil {
						return nil, err
					}

					value, err := strconv.Atoi(valueString)
					if err != nil {
						return nil, err
					}

					values[j].Value = value
				case 2:
					valuePrompt := promptui.Prompt{
						Label: "Set the value",
						Validate: func(s string) error {
							if _, err := strconv.Atoi(s); err != nil {
								return err
							} else {
								return nil
							}
						},
					}

					valueString, err := valuePrompt.Run()
					if err != nil {
						return nil, err
					}

					value, err := strconv.Atoi(valueString)
					if err != nil {
						return nil, err
					}

					values[j].Value = value
				default:
					panic(errors.New("invalid value"))

				}
			}

			statOptions[i].Minimum = values[1].Value
			statOptions[i].Maximum = values[2].Value
		}
	}

	for i := range statOptions {
		if statOptions[i].Minimum == -1 {
			statOptions[i].Minimum = 0
		}

		if statOptions[i].Maximum == -1 {
			statOptions[i].Maximum = math.MaxInt
		}
	}

	return filter.NewStatFilter(
		statOptions[1].Minimum, statOptions[1].Maximum,
		statOptions[2].Minimum, statOptions[2].Maximum,
		statOptions[3].Minimum, statOptions[3].Maximum,
		statOptions[4].Minimum, statOptions[4].Maximum,
		statOptions[5].Minimum, statOptions[5].Maximum,
		statOptions[6].Minimum, statOptions[6].Maximum,
	)
}

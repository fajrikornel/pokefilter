package model

type AbilityStore struct {
	sqlStore *InMemorySqlStore
}

type Ability struct {
	Id   int
	Name string
}

func NewAbilityStore(sqlStore *InMemorySqlStore) *AbilityStore {
	return &AbilityStore{
		sqlStore: sqlStore,
	}
}

func (m *AbilityStore) GetAllAbilities() ([]Ability, error) {
	res, err := m.sqlStore.ExecuteQuery("SELECT id, name FROM abilities")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	abilities := make([]Ability, 0)
	for res.Next() {
		var r Ability
		err = res.Scan(&r.Id, &r.Name)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, r)
	}

	return abilities, nil
}

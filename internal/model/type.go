package model

type TypeStore struct {
	sqlStore *InMemorySqlStore
}

type Type struct {
	Id   int
	Name string
}

func NewTypeStore(sqlStore *InMemorySqlStore) *TypeStore {
	return &TypeStore{
		sqlStore: sqlStore,
	}
}

func (m *TypeStore) GetAllTypes() ([]Type, error) {
	res, err := m.sqlStore.ExecuteQuery("SELECT id, name FROM types")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	types := make([]Type, 0)
	for res.Next() {
		var r Type
		err = res.Scan(&r.Id, &r.Name)
		if err != nil {
			return nil, err
		}

		types = append(types, r)
	}

	return types, nil
}

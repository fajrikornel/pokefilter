package model

type DamageClassStore struct {
	sqlStore *InMemorySqlStore
}

type DamageClass struct {
	Id   int
	Name string
}

func NewDamageClassStore(sqlStore *InMemorySqlStore) *DamageClassStore {
	return &DamageClassStore{
		sqlStore: sqlStore,
	}
}

func (m *DamageClassStore) GetAllDamageClasses() ([]DamageClass, error) {
	res, err := m.sqlStore.ExecuteQuery("SELECT id, name FROM damage_classes")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	damageClasses := make([]DamageClass, 0)
	for res.Next() {
		var r DamageClass
		err = res.Scan(&r.Id, &r.Name)
		if err != nil {
			return nil, err
		}

		damageClasses = append(damageClasses, r)
	}

	return damageClasses, nil
}

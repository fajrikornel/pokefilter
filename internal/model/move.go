package model

type MoveStore struct {
	sqlStore *InMemorySqlStore
}

type Move struct {
	id            int
	name          string
	typeId        int
	damageClassId int
}

func NewMoveStore(sqlStore *InMemorySqlStore) *MoveStore {
	return &MoveStore{
		sqlStore: sqlStore,
	}
}

func (m *MoveStore) GetAllMoves() ([]Move, error) {
	res, err := m.sqlStore.ExecuteQuery("SELECT id, name, type_id, damage_class_id FROM moves")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	moves := make([]Move, 0)
	for res.Next() {
		var r Move
		err = res.Scan(&r.id, &r.name, &r.typeId, &r.damageClassId)
		if err != nil {
			return nil, err
		}

		moves = append(moves, r)
	}

	return moves, nil
}

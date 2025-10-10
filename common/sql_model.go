package common

type SQLModel struct {
	ID     int     `json:"-" gorm:"column:id;primaryKey;autoIncrement;"`
	Status int     `json:"status" gorm:"column:status"`
	FakeId *string `json:"id" gorm:"-"`
}

func (m *SQLModel) GenUID() {
	uid, _ := MaskID(int64(m.ID))
	m.FakeId = &uid
}

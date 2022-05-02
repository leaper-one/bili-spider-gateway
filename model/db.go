package model

type RankModel struct {
	ID        int `gorm:"primaryKey"`
	Uid       string
	RoomId    string
	Rank      int  `gorm:"default:0"`
	qn        int  `gorm:"default:0"`
	HasNote   bool `gorm:"default:False"`
	Timestamp int  `gorm:"default:0;not null:True'"`
}

package model

type RankModel struct {
	ID        int   `gorm:"primaryKey"`          // id
	Uid       int   `gorm:"default:0"`           // uid
	RoomId    int   `gorm:"default:0"`           // room_id
	Rank      int   `gorm:"default:0"`           // rank
	Qn        int   `gorm:"default:0"`           // qn
	HasNote   bool  `gorm:"default:False"`       // has_note
	Timestamp int64 `gorm:"default:0;not null'"` // timestamp
}

func (RankModel) TableName() string {
	return "model_rank"
}

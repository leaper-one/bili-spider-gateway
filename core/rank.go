package core

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type (
	BiliRank struct {
		gorm.Model
		Buid      int64 `gorm:"default:0"`
		RoomId    int64 `gorm:"default:0"`
		Rank      int64 `gorm:"default:0"`
		Weight    int64 `gorm:"default:0"`
		IsConcern bool  `gorm:"default:false"`
	}

	BiliRankStore interface {
		Save(ctx context.Context, rank *BiliRank) error
		GetRanks(ctx context.Context, roomId, buid int64, time_stamp time.Time) ([]*BiliRank, error)
	}
)

package rank

import (
	"context"
	"time"

	"github.com/leaper-one/2SOMEone/util"
	"github.com/leaper-one/bili-spider-gateway/core"
)

func New(db *util.DB) core.BiliRankStore {
	return &biliRankStore{db: db}
}

type biliRankStore struct {
	db *util.DB
}

func toUpdateParams(bili_rank *core.BiliRank) map[string]interface{} {
	return map[string]interface{}{
		"buid":       bili_rank.Buid,
		"room_id":    bili_rank.RoomId,
		"rank":       bili_rank.Rank,
		"gift_value": bili_rank.GiftValue,
		"is_concern": bili_rank.IsConcern,
	}
}

func update(db *util.DB, bili_rank *core.BiliRank) (int64, error) {
	updates := toUpdateParams(bili_rank)
	tx := db.Update().Model(bili_rank).Where("buid = ?", bili_rank.Buid).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (b *biliRankStore) Save(_ context.Context, bili_rank *core.BiliRank) error {
	return b.db.Tx(func(tx *util.DB) error {
		var rows int64
		var err error
		rows, err = update(tx, bili_rank)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(bili_rank).Error
		}

		return nil
	})
}

func (b *biliRankStore) GetRanks(_ context.Context, room_id, limit int64, time_stamp time.Time) ([]*core.BiliRank, error) {
	rankdb := b.db.View().Where("room_id = ? AND updated_at > ?", room_id, time_stamp).Model(&core.BiliRank{})

	// var count int64
	// rankdb.Count(&count)

	bili_ranks := []*core.BiliRank{}
	rankdb.Order("updated_at").Limit(int(limit)).Find(&bili_ranks)
	return bili_ranks, nil
}

package rank

import (
	"github.com/leaper-one/bili-spider-gateway/global"
	"github.com/leaper-one/bili-spider-gateway/model"
)

func Hasnote(hasnote []model.HasnoteDatas) {
	for _, h := range hasnote {
		var rank model.RankModel

		r := global.DB.Where("room_id = ?", h.Roomid)
		r.Where("uid = ?", h.Uid)
		r.Find(&rank).Model(&rank).Update("has_note", true)
	}
}

func Denote(hasnote []model.HasnoteDatas) {
	for _, h := range hasnote {
		var rank model.RankModel
		r := global.DB.Where("room_id = ?", h.Roomid)
		r.Where("uid = ?", h.Uid)
		r.Find(&rank).Model(&rank).Update("has_note", false)
	}
}

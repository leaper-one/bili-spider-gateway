package rank

import (
	"log"

	"github.com/leaper-one/bili-spider-gateway/errorinit"
	"github.com/leaper-one/bili-spider-gateway/global"
	"github.com/leaper-one/bili-spider-gateway/model"
)

func Getrank(info model.Getrank) (datas []model.GetrankDatas, err error) {
	var ranks []model.RankModel

	switch {
	case info.Roomid == 0 && info.Timestamp == 0:
		global.DB.Where("has_note = ?", true).Find(&ranks)
	case info.Roomid != 0 && info.Timestamp == 0:
		r := global.DB.Where("has_note = ?", true)
		r.Where("room_id = ?", info.Roomid)
		r.Find(&ranks)
	case info.Roomid == 0 && info.Timestamp != 0:
		r := global.DB.Where("has_note = ?", true)
		r.Where("timestamp > ?", info.Timestamp)
		r.Find(&ranks)
	case info.Roomid != 0 && info.Timestamp != 0:
		r := global.DB.Where("has_note = ?", true)
		r.Where("room_id = ?", info.Roomid)
		r.Where("timestamp > ?", info.Timestamp)
		r.Find(&ranks)
	default:
		err := errorinit.PostValueError
		log.Println(err)
		return nil, err
	}

	for _, r := range ranks {
		datas = append(datas, model.GetrankDatas{Uid: r.Uid, Roomid: r.RoomId, Rank: r.Rank})
	}

	return
}

package handle

import (
	"ginForBH/errorinit"
	"ginForBH/global"
	"ginForBH/model"
	"log"
	"time"
)

func PostGetrank(info model.Getrank) (back model.GetrankBack, err error) {

	var ranks []model.RankModel
	var datas []model.GetrankDatas
	var timestamp = time.Now().Unix()

	switch {
	case info.Roomid == 0 && info.Timestamp == 0:
		global.DB.Where("has_note = ?", true).Find(&ranks)
	case info.Roomid != 0 && info.Timestamp == 0:
		r := global.DB.Where("has_note = ?", true).Find(&ranks)
		r.Where("room_id = ?", info.Roomid).Find(&ranks)
	case info.Roomid == 0 && info.Timestamp != 0:
		r := global.DB.Where("has_note = ?", true).Find(&ranks)
		r.Where("timestamp > ?", info.Timestamp)
	case info.Roomid != 0 && info.Timestamp != 0:
		r := global.DB.Where("has_note = ?", true).Find(&ranks)
		r.Where("room_id = ?", info.Roomid).Find(&ranks)
		r.Where("timestamp > ?", info.Timestamp)
	default:
		err = errorinit.PostValueError
		log.Println(err)
		return
	}

	for _, r := range ranks {
		datas = append(datas, model.GetrankDatas{Uid: r.Uid, Roomid: r.RoomId, Rank: r.Rank})
	}

	back = model.GetrankBack{
		Msg:       "ok",
		Timestamp: timestamp,
		Data:      datas,
	}

	return
}

func PostHasnote(info model.Hasnote) (err error) {
	var hasnote = info.Data

	if info.Secret != global.Secret {
		err = errorinit.SecretError
		return
	}

	if hasnote != nil {
		for _, h := range hasnote {
			var rank model.RankModel
			r := global.DB.Where("has_note = ?", false).Find(&rank)
			r.Where("room_id = ?", h.Roomid).Find(&rank)
			r.Model(&rank).Where("uid = ?", h.Uid).Update("has_note", true)
		}
	}

	return
}

func PostDenote(info model.Hasnote) (err error) {
	var hasnote = info.Data

	if info.Secret != global.Secret {
		err = errorinit.SecretError
		return
	}

	if hasnote != nil {
		for _, h := range hasnote {
			var rank model.RankModel
			r := global.DB.Where("has_note = ?", true).Find(&rank)
			r.Where("room_id = ?", h.Roomid).Find(&rank)
			r.Where("uid = ?", h.Uid).Find(&rank)
			r.Model(&rank).Update("has_note", false)
		}
	}

	return
}

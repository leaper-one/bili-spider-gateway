package rank

import (
	"ginForBH/global"
	"ginForBH/model"
	"time"
)

/*
该文件用于对提交的Uid roomid qn进行排名
*/

func Rank(n model.RankeDatas) (ranke model.RankModel) {
	var ranks []model.RankModel
	search := global.DB.Where("uid = ?", n.Uid)
	search.Where("room_id = ?", n.Roomid).Find(&ranks)

	if search.RowsAffected == 0 {
		ranke = createFunc(n, false)
	} else {
		ranke = updateFunc(n)
	}

	return
}

func updateFunc(n model.RankeDatas) (ranke model.RankModel) {
	var rank model.RankModel
	search := global.DB.Where("uid = ?", n.Uid)
	search.Where("room_id = ?", n.Roomid).Find(&rank)
	qnDb := rank.Qn
	hasnoteDb := rank.HasNote
	n.Qn = qnDb + n.Qn
	search.Delete(&rank)
	ranke = createFunc(n, hasnoteDb)
	return
}

func createFunc(n model.RankeDatas, hasnote bool) (ranke model.RankModel) {
	var ranks []model.RankModel
	var r = 0

	search := global.DB.Where("room_id = ?", n.Roomid)
	search.Order("rank desc").Find(&ranks)
	if search.RowsAffected == 0 {
		r = (global.Rankmax - global.Rankmin) / 2
	} else {
		switch {
		case n.Qn > ranks[0].Qn:
			r = ranks[0].Rank + global.Rankadd
		case n.Qn < ranks[len(ranks)-1].Qn:
			r = ranks[len(ranks)-1].Rank - global.Rankadd
		case len(ranks) == 1:
			switch {
			case n.Qn >= ranks[0].Qn:
				r = ranks[0].Rank + global.Rankadd
			default:
				r = ranks[0].Rank - global.Rankadd
			}
		default:
			for i, _ := range ranks {
				if ranks[i].Qn >= n.Qn && n.Qn >= ranks[i+1].Qn {
					r = (ranks[i].Rank + ranks[i+1].Rank) / 2
					break
				}
			}
		}
	}

	ranke = model.RankModel{
		Uid:       n.Uid,
		RoomId:    n.Roomid,
		Rank:      r,
		Qn:        n.Qn,
		HasNote:   hasnote,
		Timestamp: time.Now().Unix(),
	}
	return
}

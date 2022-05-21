package handle

import (
	"ginForBH/errorinit"
	"ginForBH/global"
	"ginForBH/model"
	"ginForBH/rank"
	"time"
)

/*
该文件为处理函数
*/

func PostGetrank(info model.Getrank) (back model.GetrankBack, err error) {
	timestamp := time.Now().Unix()

	datas, err := rank.Getrank(info)
	if err != nil {
		return
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

	if len(hasnote) != 0 {
		switch {
		case info.Msg == "hasnote":
			rank.Hasnote(hasnote)
		case info.Msg == "denote":
			rank.Denote(hasnote)
		default:
			err = errorinit.PostValueError
			return err
		}
	}

	return
}

func PostRanke(info model.Ranke) (err error) {
	var nrank = info.Data
	var rankes []model.RankModel

	if info.Secret != global.Secret {
		err = errorinit.SecretError
		return
	}

	if nrank != nil {
		for _, r := range nrank {
			rankes = append(rankes, rank.Rank(r))
		}
	}

	global.DB.Create(&rankes)
	return
}

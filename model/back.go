package model

type GetrankBack struct {
	Msg       string         `json:"msg"`
	Timestamp int64          `json:"timestamp"`
	Data      []GetrankDatas `json:"data"`
}

type GetrankDatas struct {
	Uid    int `json:"uid"`
	Roomid int `json:"roomid"`
	Rank   int `json:"rank"`
}

type HasnoteBack struct {
}

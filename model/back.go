package model

/*
该文件的模型为返回值的JSON结构体
*/

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

package model

type Getrank struct {
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	Roomid    int    `json:"roomid"`
}

type Hasnote struct {
	Msg       string         `json:"msg"`
	Timestamp int64          `json:"timestamp"`
	Secret    string         `json:"secret"`
	Data      []HasnoteDatas `json:"data"`
}

type HasnoteDatas struct {
	Uid    int `json:"uid"`
	Roomid int `json:"roomid"`
}

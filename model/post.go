package model

/*
该文件下模型为Post传入Json模型
*/

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

type Ranke struct {
	Msg       string       `json:"msg"`
	Timestamp int64        `json:"timestamp"`
	Secret    string       `json:"secret"`
	Data      []RankeDatas `json:"data"`
}

type RankeDatas struct {
	Roomid int `json:"roomid"`
	Uid    int `json:"uid"`
	Qn     int `json:"qn"`
}

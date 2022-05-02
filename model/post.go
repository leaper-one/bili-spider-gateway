package model

type PostFormGet struct {
	Msg       string    `json:"msg"`
	Timestamp int       `json:"timestamp"`
	Datas     []DataGet `json:"data"`
}

type DataGet struct {
	Uid    string `json:"uid"`
	Roomid string `json:"roomid"`
}

type PostFormBack struct {
	Msg       string     `json:"msg"`
	Timestamp int        `json:"timestamp"`
	Datas     []DataBack `json:"data"`
}

type DataBack struct {
	Uid    string `json:"uid"`
	Roomid string `json:"roomid"`
	Rank   int    `json:"rank"`
}

package handle

import "ginForBH/model"

func Handle(get model.PostFormGet) (back model.PostFormBack, err error) {

	back.Msg = get.Msg
	back.Timestamp = get.Timestamp
	back.Datas = nil

	return
}

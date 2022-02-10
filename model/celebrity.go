package model

type Celebrity struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	Gender   string `json:"gender"`
	Sign     string `json:"sign"` // 星座
	Birth    string `json:"birth"`
	Hometown string `json:"hometown"`
	Job      string `json:"job"`
	IMDb     string `json:"IMDb"`
	Brief    string `json:"brief"`
}

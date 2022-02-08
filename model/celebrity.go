package model

type Celebrity struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	NameEn   string `json:"name_en,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Sign     string `json:"sign,omitempty"` // 星座
	Birth    string `json:"birth,omitempty"`
	Hometown string `json:"hometown,omitempty"`
	Job      string `json:"job,omitempty"`
	IMDb     string `json:"IMDb,omitempty"`
	Brief    string `json:"brief"`
}

package tables

type Currency struct {
	id       int    `json:"id"`
	Code     string `json:"code"`
	FullName string `json:"fullname"`
	Sign     string `json:"sign"`
}

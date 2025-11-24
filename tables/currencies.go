package tables

type Currency struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	FullName string  `json:"fullname"`
	Sign     string  `json:"sign"`
	ToDollar float32 `json:"toDollar"`
}

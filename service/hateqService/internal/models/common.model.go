package models

type ResponseData struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type CurrentTokenReq struct {
	Date    string
	AdminID string
}

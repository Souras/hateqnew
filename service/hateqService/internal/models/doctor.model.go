package models

type AdminUserData struct {
	ID             int64  `json:"ID"`
	AdminID        string `json:"adminID"`
	Name           string `json:"name"`
	Pwd            string `json:"pwd"`
	IsOnline       bool   `json:"isOnline"`
	AcceptingToken bool   `json:"isAcceptingToken"`
	AutoRefresh    bool   `json:"autoRefresh"`
	Coordinates    string `json:"coordinates"`
}

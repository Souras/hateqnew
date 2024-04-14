package models

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type QueueData struct {
	ID          int64  `json:"ID"`
	TokenNur    string `json:"tokenNur"`
	Name        string `json:"name"`
	IsActive    bool   `json:"isActive"`
	IsCancelled bool   `json:"isCancelled"`
	TimeSlot    string `json:"timeslot"`
	AdminID     string `json:"adminID"`
	MobileNo    string `json:"mobileNo"`
	InsertTime  string `json:"insertTime"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Operating   string `json:"operating"`
	OsVersion   string `json:"osVersion"`
	Duration    string `json:"duration,omitempty"`
}

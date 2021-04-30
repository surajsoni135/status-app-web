package models

type User struct {
	MobileNo      string `json:"mobileNo"`
	U_Id          string `json:"trillId"`
	LastLoginTime string `json:"lastLoginTime"`
}

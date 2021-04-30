package models

type Status struct {
	S_Id    string `json:"s_id"`
	S_CId   string `json:"s_cid"`
	S_Date  string `json:"s_date"`
	S_Hen   string `json:"s_hen"`
	S_Hindi string `json:"s_hindi"`
	S_Like  string `json:"s_like"`
	S_Share string `json:"s_share"`
}

type StatusList struct {
	Page_id       string    `json:"page_id"`
	Total_Records string    `json:"total_records"`
	Results       []*Status `json:"results"`
}

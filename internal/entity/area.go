package entity

type Area struct {
	CordsList  [][]int `json:"coords_list"`
	CreateBy   int     `json:"create_by"`
	CreateDate string  `json:"create_date"`
	Id         int     `json:"id"`
	IsActive   bool    `json:"is_active"`
	Name       string  `json:"name"`
	UpdateBy   int     `json:"update_by"`
	UpdateDate string  `json:"update_date"`
}

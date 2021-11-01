package global

type SortInfo struct {
	SortBy string `json:"sortBy" form:"sortBy"`
	Desc   bool   `json:"desc" form:"desc"`
}

type RequestPage struct {
	SortInfo
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

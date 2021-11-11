package global

type RequestSort struct {
	SortBy string `json:"sortBy" form:"sortBy"`
	Desc   bool   `json:"desc" form:"desc"`
}

type RequestPage struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type RequestPageAndSort struct {
	RequestPage
	RequestSort
}

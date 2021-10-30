package system

type RequestDelete struct {
	Id uint `json:"id"`
}

type RequestQueryById struct {
	Id uint `json:"id"`
}

type RequestQueryByValue struct {
	Value string `json:"value"`
}

type RequestQueryByParentId struct {
	ParentId uint `json:"parentId"`
}

type RequestPage struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type RequestPageWithParentId struct {
	Page     int  `json:"page" form:"page"`
	PageSize int  `json:"pageSize" form:"pageSize"`
	ParentId uint `json:"parentId" form:"parentId"`
}

type ResponsePage struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ResponsePageWithParentId struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	ParentId uint        `json:"parentId"`
}

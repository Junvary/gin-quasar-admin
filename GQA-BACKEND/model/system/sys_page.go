package system

type RequestQueryById struct {
	Id uint `json:"id"`
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

package utlis

type Pagination struct {
	PageSize   int         `json:"page_size"`
	Page       int         `json:"page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

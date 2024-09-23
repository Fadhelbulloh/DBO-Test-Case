package params

type Pagination struct {
	Page   int `json:"page" form:"page"`
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}

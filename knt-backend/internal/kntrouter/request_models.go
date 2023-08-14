package kntrouter

type ErrorModel struct {
	Err string `json:"Error"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type SpentFormat struct {
	Cost int `json:"moneySpent"`
}

type WebhookFormat struct {
	Balance   int    `json:"balance"`
	VunetId   string `json:"vunetid"`
	Reference string `json:"reference"`
}

type Pagination struct {
	Page    int `json:"page" validator:"required"`
	PerPage int `json:"perPage" validator:"required"`
}
